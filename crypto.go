package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
)

func SignData(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
}
func VerifySignature(data, signature []byte) bool {
	hash := sha256.Sum256(data)
	return ecdsa.VerifyASN1(publicKey, hash[:], signature)
}
