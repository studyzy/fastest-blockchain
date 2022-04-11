package main

type TxPool struct {
	txs []*Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{txs: make([]*Transaction, 0)}
}
func (p *TxPool) AddTx(tx *Transaction) {
	p.txs = append(p.txs, tx)
}
func (p *TxPool) FetchTxs() []*Transaction {
	txs := p.txs
	p.txs = make([]*Transaction, 0)
	return txs
}
