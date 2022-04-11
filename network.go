package main

type Network struct {
	onRec func([]byte)
}

func NewNetwork(onRec func([]byte)) *Network {
	return &Network{onRec: onRec}
}
func (net *Network) SendMessage(msg []byte) {
	net.onRec(msg)
}
