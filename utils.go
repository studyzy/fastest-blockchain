package main

import (
	"crypto/sha256"
	"encoding/binary"
)

func Uint32ToBytes(i uint32) []byte {
	var b [4]byte
	binary.BigEndian.PutUint32(b[0:4], i)
	return b[:]
}
func Hash(data []byte) []byte {
	h := sha256.Sum256([]byte(data))
	return h[:]
}
func CalcTxRoot(txs []*Transaction) []byte {
	data := []byte{}
	for _, tx := range txs {
		data = append(data, tx.TxHash...)
	}
	return Hash(data)
}
