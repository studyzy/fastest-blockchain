package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
)

type Client struct {
	client RpcServer_SendTxStreamClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(NET_ADDRESS, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.Connect()
	c := NewRpcServerClient(conn)

	sendClient, err := c.SendTxStream(context.Background())
	checkError(err)
	return &Client{client: sendClient}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (client *Client) SendTx(tx *Transaction) {
	ClientSendTxCount++
	//通过句柄进行调用服务端函数SayHello
	err := client.client.Send(tx)
	if err != nil {
		fmt.Println("calling SendTx() error", err)
		return
	}
	ClientSendSuccessTxCount++
	//fmt.Println(re1.Message)
}
func (client *Client) Close() {
	response, err := client.client.CloseAndRecv()
	checkError(err)
	fmt.Printf(response.Message)
}
