# fastest-blockchain
Fastest block chain demo
一个极简的超级快的区块链系统。
## 各模块的选择
### 密码学库
* 签名与验签算法选择Ed25519，实测比Ecdsa快不少。
* 哈希算法选SHA256，因为MD5和SHA1虽然更快，但不安全
### 客户端与节点网络通讯
选GRPC，而且不启用TLS，理论上启用TLS会增加CPU消耗。
Client使用SendTxStream进行流式发送，实测比SendTx快10多倍。
## 性能指标
### 硬件
MacBook Pro M1芯片8核，16G内存，SSD硬盘，本地网络
Server 64核，128G内存，SSD硬盘，本地网络
### 存证性能
在存证模式下，200字节的存证数据，1s产块间隔
M1 TPS: 66660
Server: 27W