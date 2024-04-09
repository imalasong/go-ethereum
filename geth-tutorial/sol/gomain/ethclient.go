package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"c297a4f26cf416f628e3833161de8d41e66a964b","crypto":{"cipher":"aes-128-ctr","ciphertext":"c1f5e9e9d9153f0ad89a54d2cc3cb46c7955ae14ba2e909080f0d53a5bcfa40f","cipherparams":{"iv":"ffb22658bf903ad4343d6f54e817ad50"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"0d35e88da939e561adfb8b579be06b004fe003516291e4cc606989db028c5403"},"mac":"46aa50d83d44f02b018d0719360ad5e929672c798414a7ed459be5a4f6765780"},"id":"223559e7-5190-4260-8dd3-637e3968e4fe","version":3}`

func main() {

	//delay()
	//Contract pending deploy: 0x731580bd5ceafba6f4519f6495f88ab16433d3ed
	//Contract instance deploy: 0x&{{c000316008} {c000316008} {c000316008}}
	//Transaction waiting to be mined: 0x0cd7d7880fbfb7f4cb06e8a155c01da46e62e963d2264b280e1e697bea02f616
	//crud()

	simulator()
}

func simulator() {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	// Since we are using a simulated backend, we will get the chain id
	// from the same place that the simulated backend gets it.
	chainID := params.AllDevChainProtocolChanges.ChainID

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("Failed to make transactor: %v", err)
	}

	sim := simulated.NewBackend(map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(9e18)},
	})

	_, tx, store, err := DeployStorage(auth, sim.Client())
	if err != nil {
		log.Fatalf("Failed to deploy smart contract: %v", err)
	}

	fmt.Printf("Deploy pending: 0x%x\n", tx.Hash())

	sim.Commit()

	tx, err = store.Store(auth, big.NewInt(420))
	if err != nil {
		log.Fatalf("Failed to call store method: %v", err)
	}
	fmt.Printf("State update pending: 0x%x\n", tx.Hash())

	sim.Commit()

	val, err := store.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to call retrieve method: %v", err)
	}
	fmt.Printf("Value: %v\n", val)
}

func crud() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	pk, err := crypto.HexToECDSA("a872098b4084bb07eb72a399dd3d212c750ddb619b021c729a79113b37a7b88f")
	if err != nil {
		log.Fatalf("load private key error: %s", err)
	}
	log.Printf("account load success, address: %s", crypto.PubkeyToAddress(pk.PublicKey))

	//auth, err := bind.NewTransactor(strings.NewReader(key), "")
	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(1337))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy the contract passing the newly created `auth` and `conn` vars
	var address common.Address = [20]byte{}
	address.SetBytes([]byte("0cd7d7880fbfb7f4cb06e8a155c01da46e62e963d2264b280e1e697bea02f616"))
	instance, err := NewStorage(address, conn)
	//new(big.Int), "Storage contract in Go!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract instance deploy: 0x%x\n", instance)

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	//存
	tx, err := instance.Store(auth, big.NewInt(10))
	if err != nil {
		log.Fatalf("Failed to Store name: %v", err)
	}
	log.Println("tx:%v", tx.Hash())

	time.Sleep(2000 * time.Millisecond) // Allow it to be processed by the local node :P
	// function call on `instance`. Retrieves pending name
	valu, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", valu)

	// Wrap the Storage contract instance into a session
	session := &StorageSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: uint64(3141592),
		},
	}
	// Call the previous methods without the option parameters
	session.Store(big.NewInt(69))
}

func delay() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	pk, err := crypto.HexToECDSA("a872098b4084bb07eb72a399dd3d212c750ddb619b021c729a79113b37a7b88f")
	if err != nil {
		log.Fatalf("load private key error: %s", err)
	}
	log.Printf("account load success, address: %s", crypto.PubkeyToAddress(pk.PublicKey))

	//auth, err := bind.NewTransactor(strings.NewReader(key), "")
	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(1337))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, instance, err := DeployStorage(auth, conn)
	//new(big.Int), "Storage contract in Go!", 0, "Go!")
	if err != nil {
		log.Fatalf("Failed to deploy new storage contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Contract instance deploy: 0x%x\n", instance)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	//存
	tx, err = instance.Store(auth, big.NewInt(10))
	if err != nil {
		log.Fatalf("Failed to Store name: %v", err)
	}
	log.Println("tx:%v", tx.Hash())

	time.Sleep(2000 * time.Millisecond) // Allow it to be processed by the local node :P
	// function call on `instance`. Retrieves pending name
	valu, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", valu)
}
