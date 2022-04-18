package main

import (
	"encoding/binary"
	"fmt"
)

func Uint32ToBytes(i uint32) []byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[0:4], i)
	return b[:]
}

func CalcTxRoot(txs []*Transaction) []byte {
	data := make([]byte, len(txs)*32)
	for i, tx := range txs {
		if tx == nil {
			panic(fmt.Sprintf("Tx[%d] is null in txs count %d", i, len(txs)))
		}
		copy(data[i*32:], tx.TxHash)
	}
	return Hash(data)
}
