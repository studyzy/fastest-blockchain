package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
)

//GenerateNewKey 生成新的公私钥对
func GenerateNewKey() (*PrivKey, *PubKey) {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	return &PrivKey{privKey: &priv}, &PubKey{pubKey: &pub}
}

type PrivKey struct {
	privKey *ed25519.PrivateKey
}

type PubKey struct {
	pubKey *ed25519.PublicKey
}

func (p *PrivKey) SignData(data []byte) ([]byte, error) {
	return ed25519.Sign(*p.privKey, data), nil
}
func (p *PubKey) VerifySignature(data, signature []byte) bool {
	return ed25519.Verify(*p.pubKey, data, signature)
}
func Hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
