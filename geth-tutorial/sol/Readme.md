
# 生成 abi
# 需要安装 solc https://github.com/ethereum/solidity/releases
solc --abi Storage.sol -o build

# binging go app
abigen --abi ./build/Storage.abi --pkg main --type Storage --out Storage.go

# 编译字节码
solc --bin Storage.sol -o Storage.bin

abigen --abi ./build/Storage.abi --pkg main --type Storage --out Storage.go --bin Storage.bin/Storage.bin