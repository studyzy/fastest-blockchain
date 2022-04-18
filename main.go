package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var TOTAL_TX = 100000 * runtime.NumCPU()

var accountMgr = NewAccountMgr()

func main() {
	fmt.Println("generate keys...")
	GenerateMemKey(accountMgr)

	txPool := NewTxPool()

	store := NewStore()
	core := NewCore(txPool, store)
	//不断产生新交易
	net := NewNetwork(func(tx *Transaction) {
		//网络收到消息后反序列化出Tx，验证签名通过，并放入TxPool
		if err := VerifyTx(tx); err != nil {
			fmt.Println("verify tx fail:" + err.Error())
			return
		}
		atomic.AddUint32(&VerifiedTx, 1)
		//判断Tx是否与账本中的重复
		if store.TxExist(tx.TxHash) {
			fmt.Printf("tx[%x] already exist in store", tx.TxHash)
			return
		}
		txPool.AddTx(tx)

	})
	net.Start()
	defer net.Stop()
	//客户端产生新交易并放入网络模块
	fmt.Println("Prepare tx...")
	start := time.Now()
	txs := GenerateTxs(TOTAL_TX)
	fmt.Printf("Generated %d tx, spend:%v\n", len(txs), time.Since(start))

	client := NewClient()
	go func(txs []*Transaction) {
		for i := 0; i < len(txs); i++ {
			tx := txs[i]
			client.SendTx(tx)
		}
	}(txs)
	//产块节点核心引擎不断产生新区块
	core.GenerateBlock()
}

func GenerateMemKey(a *AccountMgr) error {
	consPrivateKey, consPublicKey = GenerateNewKey()
	for i := 0; i < MyClientConfig.AccountCount; i++ {
		priv, pub := GenerateNewKey()
		a.AddNewAccount(IntToBytes(i), priv, pub)
	}
	return nil
}

//func GenerateKeyFile() error {
//	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
//	privateKey = priv
//	publicKey = &priv.PublicKey
//	privBytes, _ := x509.MarshalPKCS8PrivateKey(priv)
//	fmt.Printf("generate a new key:%x", privBytes)
//	return ioutil.WriteFile("key.key", privBytes, fs.ModePerm)
//}

func GenerateTx(i int) *Transaction {
	accountId := IntToBytes(i % MyClientConfig.AccountCount)
	tx := &Transaction{
		Payload:   GenPayload(uint32(i), MyClientConfig.PayloadSize),
		Sender:    accountId,
		Signature: nil,
		TxHash:    nil,
	}
	privKey, _ := accountMgr.GetAccountPrivKey(accountId)
	txBytes, _ := tx.Marshal()
	tx.Signature, _ = privKey.SignData(txBytes)
	txBytes, _ = tx.Marshal()
	tx.TxHash = Hash(txBytes)
	return tx
}
func GenerateTxs(count int) []*Transaction {

	result := make([]*Transaction, 0)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())
	countPerCPU := count / runtime.NumCPU()
	for cpu := 0; cpu < runtime.NumCPU(); cpu++ {
		go func(c int) {
			defer wg.Done()
			for i := 0; i < countPerCPU; i++ {
				tx := GenerateTx(c*countPerCPU + i)
				lock.Lock()
				result = append(result, tx)
				lock.Unlock()
			}

		}(cpu)
	}
	wg.Wait()

	return result
}

func VerifyTx(tx *Transaction) error {
	tx2 := &Transaction{
		Payload:   tx.Payload,
		Sender:    tx.Sender,
		Signature: nil,
		TxHash:    nil,
	}
	txBytes, _ := tx2.Marshal()
	pub, err := accountMgr.GetAccountPubKey(tx.Sender)
	if err != nil {
		return err
	}
	if !pub.VerifySignature(txBytes, tx.Signature) {
		return errors.New("verify fail")
	}
	return nil
}

func VerifyTxs(txs []*Transaction) error {
	for _, tx := range txs {
		if err := VerifyTx(tx); err != nil {
			return err
		}
	}
	return nil
}
