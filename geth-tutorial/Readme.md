0xc0CB870C165C560d268B42560DE27c04FFBf97b0
UTC--2024-04-03T02-37-14.545606000Z--c0cb870c165c560d268b42560de27c04ffbf97b0

0x9804c5b2B1a0F2cc9F5b58A039c7d329dBF2e744
UTC--2024-04-03T02-37-45.573206700Z--9804c5b2b1a0f2cc9f5b58a039c7d329dbf2e744

clef --keystore geth-tutorial/keystore --configdir geth-tutorial/clef --chainid 11155111

\\.\pipe\clef.ipc

geth --sepolia --datadir geth-tutorial --authrpc.addr localhost --authrpc.port 8551 --authrpc.vhosts localhost --authrpc.jwtsecret geth-tutorial/jwtsecret --http --http.api eth,net --signer=\\.\pipe\clef.ipc --http

./prysm.bat beacon-chain --execution-endpoint=\\.\pipe\clef.ipc --sepolia --checkpoint-sync-url=https://sepolia.beaconstate.info --genesis-beacon-api-url=https://sepolia.beaconstate.info
./prysm.bat beacon-chain --execution-endpoint=http://localhost:8551 --sepolia --jwt-secret=../jwtsecret --accept-terms-of-use  --checkpoint-sync-url=https://checkpoint-sync.sepolia.ethpandaops.io --genesis-beacon-api-url=https://checkpoint-sync.sepolia.ethpandaops.io

./deposit.exe new-mnemonic --num_validators=1 --mnemonic_language=english --chain=sepolia

index wise assault begin promote cloud link whip angry attend rare car must room assist tilt kick frame fury reopen hair sample ritual time

inde wise assa begi prom clou link whip angr atte rare car must room assi tilt kick fram fury reop hair samp ritu time


./prysm.bat validator accounts import --keys-dir=./validator_keys --sepolia

aaaba03c4842723cce1d0dc45ec867363eef6b9432fe81d235b376430656af24a54819b87d6a36e43eea592154ae9376

./prysm.bat validator --wallet-dir=./t --sepolia --suggested-fee-recipient=0xc0CB870C165C560d268B42560DE27c04FFBf97b0