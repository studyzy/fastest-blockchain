package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
)

var TOTAL_TX = uint32(1000000)

func main() {
	GenerateMemKey()

	txPool := NewTxPool()

	store := NewStore()
	core := NewCore(txPool, store)
	//不断产生新交易
	net := NewNetwork(func(msg []byte) {
		//网络收到消息后反序列化出Tx，验证签名通过，并放入TxPool
		tx := &Transaction{}
		tx.Unmarshal(msg)
		if VerifyTx(tx) == nil {
			txPool.AddTx(tx)
		}
	})
	//客户端不断产生新交易并放入网络模块
	go func() {
		for i := uint32(0); i < TOTAL_TX; i++ {
			tx := GenerateTx(i)
			txMsg, _ := tx.Marshal()
			net.SendMessage(txMsg)
		}
	}()
	//产块节点核心引擎不断产生新区块

	core.GenerateBlock()

}
func GenerateMemKey() error {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKey = priv
	publicKey = &priv.PublicKey
	return nil
}
func GenerateKeyFile() error {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKey = priv
	publicKey = &priv.PublicKey
	privBytes, _ := x509.MarshalPKCS8PrivateKey(priv)
	fmt.Printf("generate a new key:%x", privBytes)
	return ioutil.WriteFile("key.key", privBytes, fs.ModePerm)
}

func GenerateTx(i uint32) *Transaction {

	tx := &Transaction{
		Payload:   Uint32ToBytes(i),
		Sender:    []byte{1},
		Signature: nil,
		TxHash:    nil,
	}
	txBytes, _ := tx.Marshal()
	tx.Signature, _ = SignData(txBytes)
	txBytes, _ = tx.Marshal()
	tx.TxHash = Hash(txBytes)
	return tx
}
func GenerateTxs(count uint32) []*Transaction {
	result := make([]*Transaction, count)
	for i := uint32(0); i < count; i++ {
		tx := &Transaction{
			Payload:   Uint32ToBytes(i),
			Sender:    []byte{1},
			Signature: nil,
			TxHash:    nil,
		}
		txBytes, _ := tx.Marshal()
		tx.Signature, _ = SignData(txBytes)
		txBytes, _ = tx.Marshal()
		tx.TxHash = Hash(txBytes)
		result[i] = tx
	}
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
	if !VerifySignature(txBytes, tx.Signature) {
		return errors.New("verify fail")
	}

	return nil
}

func VerifyTxs(txs []*Transaction) error {
	for _, tx := range txs {
		tx2 := &Transaction{
			Payload:   tx.Payload,
			Sender:    tx.Sender,
			Signature: nil,
			TxHash:    nil,
		}
		txBytes, _ := tx2.Marshal()
		if !VerifySignature(txBytes, tx.Signature) {
			return errors.New("verify fail")
		}
	}
	return nil
}
