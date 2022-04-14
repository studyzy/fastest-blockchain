package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Network struct {
	onRec func(transaction *Transaction)
	conn  net.Listener
}

func (n *Network) SendTx(ctx context.Context, tx *Transaction) (*SendTxResponse, error) {
	ServerReceiveTxCount++
	n.onRec(tx)
	return &SendTxResponse{Message: "OK"}, nil
}

func NewNetwork(onRec func(transaction *Transaction)) *Network {
	return &Network{onRec: onRec}
}

func (n *Network) Start() {
	listen, err := net.Listen("tcp", NET_ADDRESS)
	if err != nil {
		fmt.Println("listen failed,err", err)
		return
	}
	//创建grpc服务
	srv := grpc.NewServer()
	//注册服务
	RegisterRpcServerServer(srv, n)
	err = srv.Serve(listen)
	if err != nil {
		fmt.Println("Serve error", err)
	}
	n.conn = listen
}

func (n *Network) Stop() {
	n.conn.Close()
}
