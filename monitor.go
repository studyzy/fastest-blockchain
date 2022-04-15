package main

import "fmt"

var ClientSendTxCount int
var ClientSendSuccessTxCount int
var ServerReceiveTxCount int
var VerifiedTx int
var InBlockTxCount int

func PrintMonitorMessage() {
	fmt.Printf("client send tx: %d/%d,server receive tx: %d,verify:%d, tx in block:%d \n",
		ClientSendSuccessTxCount, ClientSendTxCount, ServerReceiveTxCount, VerifiedTx, InBlockTxCount)
}
