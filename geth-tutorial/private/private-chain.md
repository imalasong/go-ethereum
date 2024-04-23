
1、创建node1、node2目录

2、给node1、node2创建账户
geth --datadir node1 account new

> 0x2C10ce25060D847B0e60E23C9a9eDD24371137F5

geth --datadir node2 account new
> 0x5E61992C69FE4e211a8ADA8Dc59805AC39F3c284 

geth --datadir node3 account new
> 0x87c796598Bad1c8216b6D6c98bE4DE169Fe7E93a


3、初始化gensis节点
geth init --datadir node1 genesis.json
geth init --datadir node2 genesis.json
geth init --datadir node3 genesis.json


4、配置bootnode
bootnode -genkey boot.key

bootnode -nodekey boot.key -addr :30305


5、启动节点
geth --datadir node1 --port 30306 --bootnodes "enode://e8bcc238f41b3da29d70d9d71f944bb454aac306850da608915dcd38add918b7746a9ad2a7e5b9dfe4917713f85286b0bb9737408a2289af7182d9521e55a395@127.0.0.1:0?discport=30305" --networkid 123454321 --unlock 0x2C10ce25060D847B0e60E23C9a9eDD24371137F5 --password node1/password.txt --authrpc.port 8551 --mine --miner.etherbase 0x2C10ce25060D847B0e60E23C9a9eDD24371137F5

geth --datadir node2 --port 30307 --bootnodes "enode://e8bcc238f41b3da29d70d9d71f944bb454aac306850da608915dcd38add918b7746a9ad2a7e5b9dfe4917713f85286b0bb9737408a2289af7182d9521e55a395@127.0.0.1:0?discport=30305" --networkid 123454321 --unlock 0x5E61992C69FE4e211a8ADA8Dc59805AC39F3c284 --password node2/password.txt --authrpc.port 8552
geth --datadir node3 --port 30308 --bootnodes "enode://e8bcc238f41b3da29d70d9d71f944bb454aac306850da608915dcd38add918b7746a9ad2a7e5b9dfe4917713f85286b0bb9737408a2289af7182d9521e55a395@127.0.0.1:0?discport=30305" --networkid 123454321 --unlock 0x87c796598Bad1c8216b6D6c98bE4DE169Fe7E93a --password node3/password.txt --authrpc.port 8553



geth attach node1/geth.ipc
geth attach node2/geth.ipc
geth attach node3/geth.ipc

eth.getBalance(eth.accounts[0])
eth.getBalance("0x2C10ce25060D847B0e60E23C9a9eDD24371137F5")
eth.getBalance("0x5E61992C69FE4e211a8ADA8Dc59805AC39F3c284")

eth.sendTransaction({
to: '0x5E61992C69FE4e211a8ADA8Dc59805AC39F3c284',
from: eth.accounts[0],
value: 100000
});

> 0xb737b2fac01de97b1d1b208f4341ee60c1b3df74611a4a53691707e590b662a2

eth.getBalance('0x3502C7d1285bbf42249655e29EC9Ed3Ce8589a0a');

