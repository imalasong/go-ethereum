package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"7967cd207adb9abad5e03acb335e414649fb82e7","crypto":{"cipher":"aes-128-ctr","ciphertext":"4206a9bfab28599cdfb16bfc32bfc1df963600d3fda23879fd0973f2fa9368ec","cipherparams":{"iv":"41777683116dda04a1ddd0960eb9d293"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"97c6acbe56c67d93bd3181cec4c681b69feca4bea64f0a283f1996548604dcdd"},"mac":"fd2219236d44547740e5fffbf00c5e3e98df7033173e7c9a5611c86337ac2f52"},"id":"b256b8a4-0ecb-4de2-b3b8-386223de67c6","version":3}`

func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("/Users/xiaochangbai/workspaces/goprojects/src/go-ethereum/geth-tutorial/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, instance, err := DeployStorage(auth, conn)
	//, new(big.Int), "Storage contract in Go!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	// function call on `instance`. Retrieves pending name
	name, err := instance.Retrieve(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}
