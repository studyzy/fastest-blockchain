package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Network struct {
	onRec func(transaction *Transaction)
	srv   *grpc.Server
}

func (n *Network) SendTx(ctx context.Context, transaction *Transaction) (*SendTxResponse, error) {
	ServerReceiveTxCount++
	go n.onRec(transaction)
	return nil, nil
}

func (n *Network) SendTxStream(server RpcServer_SendTxStreamServer) error {
	for {
		tx, err := server.Recv()
		checkError(err)
		ServerReceiveTxCount++
		go n.onRec(tx)

	}
	return nil
}

//func (n *Network) SendTx(ctx context.Context, tx *Transaction) (*SendTxResponse, error) {
//	ServerReceiveTxCount++
//	n.onRec(tx)
//	return &SendTxResponse{Message: "OK"}, nil
//}

func NewNetwork(onRec func(transaction *Transaction)) *Network {
	return &Network{onRec: onRec}
}

func (n *Network) Start() {
	listen, err := net.Listen("tcp", MyServerConfig.RPCServerAddress)
	if err != nil {
		fmt.Println("listen failed,err", err)
		return
	}
	//创建grpc服务
	srv := grpc.NewServer()
	n.srv = srv
	//注册服务
	RegisterRpcServerServer(srv, n)
	go func() {
		err = srv.Serve(listen)
		if err != nil {
			fmt.Println("Serve error", err)
		}
	}()

}

func (n *Network) Stop() {
	n.srv.Stop()
}
