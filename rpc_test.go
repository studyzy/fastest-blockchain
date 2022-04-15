package main

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
)

type server struct {
}

var start time.Time
var complete = false

func (s server) SendTx(txServer RpcServer_SendTxServer) error {
	start = time.Now()
	for {
		_, err := txServer.Recv()
		if err != nil {
			complete = true
			return err
		}
		ServerReceiveTxCount++
	}
	return nil
}

func TestGrpc(t *testing.T) {
	//init server
	listen, err := net.Listen("tcp", NET_ADDRESS)
	if err != nil {
		fmt.Println("listen failed,err", err)
		return
	}
	//创建grpc服务
	srv := grpc.NewServer()
	//注册服务
	RegisterRpcServerServer(srv, &server{})
	go srv.Serve(listen)
	time.Sleep(1 * time.Second)
	go clientSendTx(t)
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		t.Logf("recevied:%d,cost:%v,TPS:%v",
			ServerReceiveTxCount, time.Since(start), int(float64(ServerReceiveTxCount)/time.Since(start).Seconds()))
		if complete {
			return
		}
	}
}
func clientSendTx(t *testing.T) {
	//init client
	conn, err := grpc.Dial(NET_ADDRESS, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	conn.Connect()
	c := NewRpcServerClient(conn)
	sendClient, err := c.SendTx(context.Background())
	sign := [73]byte{}
	for i := 0; i < TOTAL_TX*10; i++ {
		err := sendClient.Send(&Transaction{
			Payload:   Uint32ToBytes(uint32(i)),
			Sender:    []byte{1},
			Signature: sign[:],
			TxHash:    sign[0:32],
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	err = sendClient.CloseSend()
	if err != nil {
		t.Fatal(err)
	}
}
