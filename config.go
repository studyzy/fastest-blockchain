package main

type ChainConfig struct {
	//产块间隔
	BlockInterval int
}
type ServerConfig struct {
	//GRPC的服务地址和端口
	RPCServerAddress string
}
type ClientConfig struct {
	//GRPC的服务地址和端口
	RPCServerAddress string
	//Bytes
	PayloadSize int
	//初始化多少个公私钥账号用于签名验签
	AccountCount int
}

var MyChainConfig = &ChainConfig{
	BlockInterval: 100,
}
var MyServerConfig = &ServerConfig{
	RPCServerAddress: ":9527",
}
var MyClientConfig = &ClientConfig{
	RPCServerAddress: "127.0.0.1:9527",
	PayloadSize:      200,
	AccountCount:     100,
}
