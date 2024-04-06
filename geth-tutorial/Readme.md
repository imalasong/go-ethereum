clef newaccount --keystore geth-tutorial/keystore

0x69537c1BBCC87f774eA115f57f7dBD56c481DD85
UTC--2024-04-05T09-17-58.267474000Z--69537c1bbcc87f774ea115f57f7dbd56c481dd85
0x785CFdf96D2788975fa4853eB3CBB7550180C156
UTC--2024-04-05T09-18-27.859435000Z--785cfdf96d2788975fa4853eb3cbb7550180c156

clef --keystore geth-tutorial/keystore --configdir geth-tutorial/clef --chainid 11155111
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