package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
)

type Client struct {
	client RpcServerClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(NET_ADDRESS, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := NewRpcServerClient(conn)
	return &Client{client: c}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (client *Client) SendTx(tx *Transaction) {

	//通过句柄进行调用服务端函数SayHello
	_, err := client.client.SendTx(context.Background(), tx)
	if err != nil {
		fmt.Println("calling SendTx() error", err)
		return
	}
	ClientSendTxCount++
	//fmt.Println(re1.Message)
}
