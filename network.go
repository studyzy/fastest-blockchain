package main

import (
	"fmt"
	"net"
	"os"
)

type Network struct {
	onRec func([]byte)
	conn  *net.UDPConn
}

func NewNetwork(onRec func([]byte)) *Network {
	return &Network{onRec: onRec}
}
func (n *Network) SendMessage(msg []byte) {
	n.onRec(msg)
}
func (n *Network) Start() {
	addr, err := net.ResolveUDPAddr("udp", NET_ADDRESS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		// Here must use make and give the length of buffer
		data := make([]byte, NET_MESSAGE_SIZE)
		_, rAddr, err := n.conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(rAddr.String(), err)
			continue
		}
		ServerReceiveTxCount++
		n.onRec(data)
		//_, err = n.conn.WriteToUDP([]byte("OK"), rAddr)
		//if err != nil {
		//	fmt.Println(err)
		//	continue
		//}
	}

}
func (n *Network) Stop() {
	n.conn.Close()
}
