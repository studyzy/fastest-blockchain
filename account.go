package main

import (
	"encoding/hex"
	"errors"
)

type AccountMgr struct {
	//AccountID->PrivateKey的映射
	accountPrivKey map[string]*PrivKey
	//AccountID->PubKey的映射
	accountPubKey map[string]*PubKey
}

func NewAccountMgr() *AccountMgr {
	return &AccountMgr{
		accountPrivKey: make(map[string]*PrivKey),
		accountPubKey:  make(map[string]*PubKey),
	}
}
func (a *AccountMgr) AddNewAccount(id []byte, priv *PrivKey, pub *PubKey) {
	a.accountPrivKey[hex.EncodeToString(id)] = priv
	a.accountPubKey[hex.EncodeToString(id)] = pub
}
func (a *AccountMgr) GetAccountPubKey(id []byte) (*PubKey, error) {
	pubKey, ok := a.accountPubKey[hex.EncodeToString(id)]
	if ok {
		return pubKey, nil
	}
	return nil, errors.New("account not found")
}
func (a *AccountMgr) GetAccountPrivKey(id []byte) (*PrivKey, error) {
	privKey, ok := a.accountPrivKey[hex.EncodeToString(id)]
	if ok {
		return privKey, nil
	}
	return nil, errors.New("account not found")
}
