package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

func main() {
	inPath := "D:\\programs\\workspaces\\goproject\\src\\go-ethereum\\geth-tutorial\\keystore\\UTC--2024-04-07T08-55-04.006087700Z--c297a4f26cf416f628e3833161de8d41e66a964b"
	outPath := "key.hex"
	password := ""
	keyjson, e := ioutil.ReadFile(inPath)
	if e != nil {
		panic(e)
	}
	key, e := keystore.DecryptKey(keyjson, password)
	if e != nil {
		panic(e)
	}
	e = crypto.SaveECDSA(outPath, key.PrivateKey)
	if e != nil {
		panic(e)
	}
	fmt.Println("Key saved to:", outPath)
}
