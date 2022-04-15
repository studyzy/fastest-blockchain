package main

import "fmt"

var ClientSendTxCount uint32
var ClientSendSuccessTxCount uint32
var ServerReceiveTxCount uint32
var VerifiedTx uint32
var InBlockTxCount uint32

func PrintMonitorMessage() {
	fmt.Printf("client send tx: %d/%d,server receive tx: %d,verify:%d, tx in block:%d \n",
		ClientSendSuccessTxCount, ClientSendTxCount, ServerReceiveTxCount, VerifiedTx, InBlockTxCount)
}
