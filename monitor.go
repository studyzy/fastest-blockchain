package main

import "fmt"

var ClientSendTxCount int
var ServerReceiveTxCount int
var VerifiedTx int
var InBlockTxCount int

func PrintMonitorMessage() {
	fmt.Printf("client send tx: %d,server receive tx: %d,tx in block:%d \n",
		ClientSendTxCount, ServerReceiveTxCount, InBlockTxCount)
}
