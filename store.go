package main

import "errors"

type Store struct {
	headers         map[uint64]*BlockHeader
	txs             map[string]*Transaction
	lastBlockHeight uint64
}

func NewStore() *Store {
	return &Store{
		headers:         make(map[uint64]*BlockHeader),
		txs:             make(map[string]*Transaction),
		lastBlockHeight: 0,
	}
}

func (s *Store) SaveBlock(block *Block) error {
	if block.Header.BlockHeight != s.lastBlockHeight+1 {
		return errors.New("invalid block height")
	}
	s.headers[block.Header.BlockHeight] = block.Header
	for _, tx := range block.Txs {
		s.txs[string(tx.TxHash)] = tx
	}
	s.lastBlockHeight = block.Header.BlockHeight
	return nil
}
func (s *Store) GetLastBlockHeight() uint64 {
	return s.lastBlockHeight
}
