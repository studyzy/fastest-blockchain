package main

import "sync"

type TxPool struct {
	txs []*Transaction
	l   sync.Mutex
}

func NewTxPool() *TxPool {
	return &TxPool{txs: make([]*Transaction, 0)}
}
func (p *TxPool) AddTx(tx *Transaction) {
	if tx == nil {
		panic("TxPool add null tx")
	}
	p.l.Lock()
	p.txs = append(p.txs, tx)
	p.l.Unlock()
}
func (p *TxPool) FetchTxs() []*Transaction {
	p.l.Lock()
	defer p.l.Unlock()
	txs := p.txs
	p.txs = make([]*Transaction, 0)
	return txs
}
