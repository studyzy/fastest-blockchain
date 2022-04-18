package main

import (
	"testing"
	"time"
)

func TestBuildMerkleTreeStore(t *testing.T) {
	TxCount := 100000
	tx10w := make([]*Transaction, TxCount)
	for i := 0; i < TxCount; i++ {
		hash := Hash(Uint32ToBytes(uint32(i)))
		tx10w[i] = &Transaction{TxHash: hash}
	}
	start := time.Now()
	for i := 0; i < 100; i++ {
		BuildMerkleTreeStore(tx10w)
		//t.Logf("tx root:%x", result[len(result)-1])
	}
	t.Logf("BuildMerkleTreeStore spend time:%v ms", time.Since(start).Milliseconds()/100)
}
func TestCalcTxRoot(t *testing.T) {
	TxCount := 100000
	tx10w := make([]*Transaction, TxCount)
	for i := 0; i < TxCount; i++ {
		hash := Hash(Uint32ToBytes(uint32(i)))
		tx10w[i] = &Transaction{TxHash: hash}
	}
	start := time.Now()
	for i := 0; i < 100; i++ {
		CalcTxRoot(tx10w)
		//t.Logf("tx root:%x", result)
	}
	t.Logf("CalcTxRoot spend time:%v ms", time.Since(start).Milliseconds()/100)
}
