package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func Uint32ToBytes(i uint32) []byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[0:4], i)
	return b[:]
}
func Hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
func CalcTxRoot(txs []*Transaction) []byte {
	var data []byte
	for i, tx := range txs {
		if tx == nil {
			panic(fmt.Sprintf("Tx[%d] is null in txs count %d", i, len(txs)))
		}
		data = append(data, tx.TxHash...)
	}
	return Hash(data)
}
