1、初始化账户
geth -datadir ./n1/data/ account new
生成public key：0xe76c26A6F2B10A852755b3ee98CD629d2126929D
geth -datadir ./n2/data/ account new
生成public key：0x1909192AC738EC0C0026A9Fb67C5Bf90b381627C
geth -datadir ./n3/data/ account new
生成public key：0xbfEcb4671e0949537d2CFa56f29d6c2B5Fb43fB2

2、将该地址复制到genesis的alloc 参数中

3、创建联盟链节点
geth -datadir ./n1/data/ init ./n1/genesis.json
geth -datadir ./n2/data/ init ./n2/genesis.json
geth -datadir ./n3/data/ init ./n3/genesis.json


4\搭建联盟链网络
geth -datadir ./n1/data --networkid 72 console