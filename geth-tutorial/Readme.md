clef newaccount --keystore geth-tutorial/keystore

0x69537c1BBCC87f774eA115f57f7dBD56c481DD85
UTC--2024-04-05T09-17-58.267474000Z--69537c1bbcc87f774ea115f57f7dbd56c481dd85
0x785CFdf96D2788975fa4853eB3CBB7550180C156
UTC--2024-04-05T09-18-27.859435000Z--785cfdf96d2788975fa4853eb3cbb7550180c156

clef --keystore geth-tutorial/keystore --configdir geth-tutorial/clef --chainid 1337
geth-tutorial/clef/clef.ipc

./prysm.sh beacon-chain generate-auth-secret
jwt.hex

geth --sepolia --datadir geth-tutorial --authrpc.addr localhost --authrpc.port 8551 --authrpc.vhosts localhost --authrpc.jwtsecret geth-tutorial/jwt.hex --http --http.api eth,net --signer=geth-tutorial/clef/clef.ipc --http

./prysm beacon-chain --execution-endpoint=\\.\pipe\clef.ipc --sepolia --checkpoint-sync-url=https://sepolia.beaconstate.info --genesis-beacon-api-url=https://sepolia.beaconstate.info
./prysm.bat beacon-chain --execution-endpoint=http://localhost:8551 --sepolia --jwt-secret=../jwtsecret --accept-terms-of-use  --checkpoint-sync-url=https://checkpoint-sync.sepolia.ethpandaops.io --genesis-beacon-api-url=https://checkpoint-sync.sepolia.ethpandaops.io

./prysm.sh beacon-chain --execution-endpoint=http://localhost:8551 --sepolia --jwt-secret=../jwt.hex --checkpoint-sync-url=https://sepolia.beaconstate.info --genesis-beacon-api-url=https://sepolia.beaconstate.info

./deposit.exe new-mnemonic --num_validators=1 --mnemonic_language=english --chain=sepolia

index wise assault begin promote cloud link whip angry attend rare car must room assist tilt kick frame fury reopen hair sample ritual time

inde wise assa begi prom clou link whip angr atte rare car must room assi tilt kick fram fury reop hair samp ritu time


./prysm.bat validator accounts import --keys-dir=./validator_keys --sepolia

aaaba03c4842723cce1d0dc45ec867363eef6b9432fe81d235b376430656af24a54819b87d6a36e43eea592154ae9376

./prysm.bat validator --wallet-dir=./t --sepolia --suggested-fee-recipient=0xc0CB870C165C560d268B42560DE27c04FFBf97b0


eth.getBalance('0xda73f782bdFD88166503dDc18AF13B4207c3d529')


geth attach http://127.0.0.1:8545




## dev mode

geth --dev --datadir geth-tutorial --http --http.api eth,web3,net --http.corsdomain "https://remix.ethereum.org"

geth attach \\.\pipe\geth.ipc

eth.accounts
eth.getBalance(eth.accounts[0])/1e18

web3.fromWei(eth.getBalance(eth.accounts[0]))

clef newaccount --keystore geth-tutorial/keystore

eth.sendTransaction({from:eth.accounts[0], to: eth.accounts[1], value: web3.toWei(50, "ether")})

web3.fromWei(eth.getBalance(eth.accounts[1]))

eth.getTransaction("0x6b72b6ed867f78da4d684e958e81ee3de28d19a00df9fa2c75b8e5067071fff4")

web3.eth.getStorageAt("0x2e64cec1", 0)

eth.getCode("0x705280c319637e0C215B2A4b888Bf3169fb44e5A")


var contractABI = [{"constant":false,"inputs":[],"name":"doSomething","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}];
var contractBytecode = "0x608060405234801561001057600080fd5b5060d08061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633fa4f245146100465780636d4ce63c14610067578063e42cb22b14610090575b600080fd5b61004e61009d565b60405161005b91906100d2565b60405180910390f35b6100736100bb565b60405161008091906100d2565b60405180910390f35b6100986100e0565b6040516100a591906100d2565b60405180910390f35b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610145576100d2565b81600080810190506101496000805460ff19169055604080519081016040528093929190818152602001838360200280828437820191505050505050602060405180830381858888f1935050505015801561019c573d6000803e3d6000fd5b50565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610209576100d2565b81600080810190506102049000541561025e565b816000808101905061024f900054610273565b6000808081019050919050565b6000816000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102d8576100d2565b81600080810190506102c290005415610301565b8160008081019050610302900054600160a060020a03191633908117909155600380548216905560048401829055600280548216905560078054821690555600925054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638d70e88c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156103cb57600080fd5b505af11580156103df573d6000803e3d6000fd5b505050506040513d60208110156103f557600080fd5b810190808051906020019092919050505090506103f6565b005b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610495576100d2565b816000808101905061047e90005461049e565b6000808081019050919050565b6000816000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156104ff576100d2565b81600080810190506104e890005415610520565b8160008081019050610511900054600160a060020a031916339081179091556003805482169055600484018290556002805482169055600780548216905556006000816000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638d70e88c6040518163ffffffff167c010000000000000000000000000"


var contract = eth.contract(contractABI);
var deployedContract = contract.new({data: contractBytecode, from: eth.accounts[0], gas: 1000000});

