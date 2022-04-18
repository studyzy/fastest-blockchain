package main

import "errors"

func RunVM(tx *Transaction) error {
	if tx.TxType == TxType_Evidence {
		//存证，什么都不用做
		return nil
	}
	return errors.New("unsupported tx type")
}
