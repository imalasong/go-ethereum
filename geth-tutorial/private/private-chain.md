
1、创建node1、node2目录

2、给node1、node2创建账户
geth --datadir node1 account new

> 0x3502C7d1285bbf42249655e29EC9Ed3Ce8589a0a

geth --datadir node2 account new
> 0xd67515c370A8006E977Ce1C6a5AC00095f6Fb54f


3、初始化gensis节点
geth init --datadir node1 genesis.json
geth init --datadir node2 genesis.json


4、配置bootnode
bootnode -genkey boot.key

bootnode -nodekey boot.key -addr :30305


5、启动节点
geth --datadir node1 --port 30306 --bootnodes enode://d725d5e0a6a562650ce3ce4a482ea064d5b1557302672a7668705a148d9560ebc8adb21f464d4bcc5604e0572704d59175edbf6469365b8165ca2eef5e9e2a97@127.0.0.1:0?discport=30305  --networkid 123454321 --unlock 0x3502C7d1285bbf42249655e29EC9Ed3Ce8589a0a --password node1/password.txt --authrpc.port 8551 --mine --miner.etherbase 0x3502C7d1285bbf42249655e29EC9Ed3Ce8589a0a --ipcdisable

geth --datadir node2 --port 30307 --bootnodes enode://d725d5e0a6a562650ce3ce4a482ea064d5b1557302672a7668705a148d9560ebc8adb21f464d4bcc5604e0572704d59175edbf6469365b8165ca2eef5e9e2a97@127.0.0.1:0?discport=30305  --networkid 123454321 --unlock 0xd67515c370A8006E977Ce1C6a5AC00095f6Fb54f --password node2/password.txt --authrpc.port 8552 --ipcdisable



geth attach datadir/geth.ipc

eth.getBalance(eth.accounts[0])

eth.sendTransaction({
to: '0x3502C7d1285bbf42249655e29EC9Ed3Ce8589a0a',
from: eth.accounts[0],
value: 25000
});

> 0x2d69c36012e7b69c89ec1e63872c75c29b3ffb43e97c2fa0b1d177b21b23baa5

eth.getBalance('0x3502C7d1285bbf42249655e29EC9Ed3Ce8589a0a');

