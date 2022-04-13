package main

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
}

func NewClient() *Client {
	conn, err := net.Dial("udp", NET_ADDRESS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &Client{conn: conn}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (client *Client) SendTx(tx *Transaction) {
	txMsg, _ := tx.Marshal()
	_, err := client.conn.Write(txMsg)

	checkError(err)
	ClientSendTxCount++
	//msg := make([]byte, config.SERVER_RECV_LEN)
	//n, err = conn.Read(msg)
	//checkError(err)
	//
	//fmt.Println("Response:", string(msg))

}
