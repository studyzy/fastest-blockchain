package main

import (
	"fmt"
	"time"
)

type Core struct {
	txpool *TxPool
	store  *Store
}

func NewCore(txPool *TxPool, store *Store) *Core {
	return &Core{
		txpool: txPool,
		store:  store,
	}
}
func (core *Core) GenerateBlock() {
	var preHash []byte
	var preHeight uint64
	t := time.NewTicker(time.Millisecond * time.Duration(MyChainConfig.BlockInterval))
	start := time.Now()
	firstBlockStart := time.Now()
	complete := false
	go func() {
		for {
			<-t.C

			//从交易池获取未打包交易
			txs := core.txpool.FetchTxs()
			if len(txs) == 0 {
				if preHeight == 0 {
					continue
				}
				fmt.Printf("no tx in pool  cost: %v, TPS: %v\n", time.Since(firstBlockStart),
					float64(TOTAL_TX)/time.Since(firstBlockStart).Seconds())
				complete = true
				return
			}

			//产生新区块
			newBlock := GenerateBlock(preHeight+1, preHash, txs)
			//保存新区块
			core.store.SaveBlock(newBlock)
			//更新变量
			preHeight++
			preHash = newBlock.Header.BlockHash
			newBlockBytes, _ := newBlock.Marshal()
			fmt.Printf("Generate new block[%d] size=%d tx count= %d, cost: %v, TPS: %v\n", newBlock.Header.BlockHeight,
				len(newBlockBytes), len(txs), time.Since(start), float64(len(txs))/time.Since(start).Seconds())
			InBlockTxCount += uint32(len(txs))
			start = time.Now()
		}
	}()
	for {

		time.Sleep(time.Second * 1)
		PrintMonitorMessage()
		if complete {
			fmt.Println("complete")
			return
		}
	}
}

//GenerateBlock 根据一堆Txs和当前高度，上一区块Hash，生成新的区块
func GenerateBlock(height uint64, preBlockHash []byte, txs []*Transaction) *Block {
	for _, tx := range txs {
		RunVM(tx)
	}
	header := &BlockHeader{
		BlockHeight:    height,
		BlockHash:      nil,
		PreBlockHash:   preBlockHash,
		TxRoot:         CalcTxRoot(txs),
		BlockTimestamp: time.Now().Unix(),
		Proposer:       []byte{1},
		Signature:      nil,
	}
	headerBytes, _ := header.Marshal()
	header.Signature, _ = consPrivateKey.SignData(headerBytes)
	headerBytes, _ = header.Marshal()
	header.BlockHash = Hash(headerBytes)
	return &Block{
		Header: header,
		Txs:    txs,
	}
}
