package main

import (
	"crypto/ed25519"
	"crypto/sha256"
)

func SignData(data []byte) ([]byte, error) {
	return ed25519.Sign(*privateKey, data), nil
	//return ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
}
func VerifySignature(data, signature []byte) bool {
	//hash := sha256.Sum256(data)
	//return ecdsa.VerifyASN1(publicKey, hash[:], signature)
	return ed25519.Verify(*publicKey, data, signature)
}
func Hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
