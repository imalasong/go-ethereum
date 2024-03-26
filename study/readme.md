### 包目录分析
``` lua
accounts      --实现了一个高等级的以太坊账户管理
beacon
build         --主要是编译和构建的一些脚本和配置
cmd           --命令行工具，又分了很多的命令行工具，下面一个一个介绍
├── abidump   --
├── abigen    --源代码生成器，用于将以太坊合约定义转换为易于使用、编译时类型安全的Go包
├── blsync    
├── bootnode  --启动一个仅仅实现网络发现的节点
├── clef
├── devp2p
├── era
├── ethkey
├── evm       --以太坊虚拟机的开发工具， 用来提供一个可配置的，受隔离的代码调试环境-
├── geth      --以太坊命令行客户端，最重要的一个工具
├── p2psim    --提供了一个工具来模拟http的API
├── rlpdump   --提供了一个RLP数据的格式化输出
└── utils     --提供了一些公共的工具
common        --提供了一些公共的工具类
consensus     --提供了以太坊的一些共识算法，比如ethhash, clique(proof-of-authority)
console       --console类
core          --以太坊的核心数据结构和算法(虚拟机，状态，区块链，布隆过滤器)
crypto        --加密和hash算法
docs          --官方文档
eth           --实现了以太坊的协议
ethclient     --提供了以太坊的RPC客户端
ethdb         --eth的数据库(包括实际使用的leveldb和供测试使用的内存数据库)
ethstats      --提供网络状态的报告
event         --处理实时的事件
graphql
internal
log           --提供对人机都友好的日志信息
metrics       --提供磁盘计数器
miner         --提供以太坊的区块创建和挖矿
node          --以太坊的多种类型的节点
p2p           --以太坊p2p网络协议
params
rlp           --以太坊序列化处理
rpc           --远程方法调用
signer
swarm         --swarm网络处理
tests         --测试相关
trie          --以太坊重要的数据结构Package trie implements Merkle Patricia Tries.
triedb
```


推荐博客：
https://knarfeh.com/2018/03/10/go-ethereum%20%E6%BA%90%E7%A0%81%E7%AC%94%E8%AE%B0%EF%BC%88%E6%A6%82%E8%A7%88%EF%BC%89/