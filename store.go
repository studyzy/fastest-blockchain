package main

import (
	"encoding/hex"
	"errors"
	"sync"
)

type Store struct {
	headers         map[uint64]*BlockHeader
	txs             map[string]*Transaction
	lastBlockHeight uint64
	lock            sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		headers:         make(map[uint64]*BlockHeader),
		txs:             make(map[string]*Transaction),
		lastBlockHeight: 0,
	}
}

//SaveBlock 保存一个区块到账本中
func (s *Store) SaveBlock(block *Block) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if block.Header.BlockHeight != s.lastBlockHeight+1 {
		return errors.New("invalid block height")
	}
	s.headers[block.Header.BlockHeight] = block.Header
	for _, tx := range block.Txs {
		s.txs[hex.EncodeToString(tx.TxHash)] = tx
	}
	s.lastBlockHeight = block.Header.BlockHeight
	return nil
}

//GetLastBlockHeight 获得存储的最新区块高度
func (s *Store) GetLastBlockHeight() uint64 {
	return s.lastBlockHeight
}

//TxExist 判断一个Tx是否已经存在于账本中
func (s *Store) TxExist(txHash []byte) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.txs[hex.EncodeToString(txHash)]
	return ok
}
