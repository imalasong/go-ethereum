package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
)

func main() {

	marshal, err := json.Marshal(Person{})

	if err != nil {
		log.Error("encode error:", err)
	}
	fmt.Printf("%s\n", marshal)

	p := Person{}
	personJson := `{"name":"aaa","Age":0,"Gender":0}`
	//err = json.Unmarshal(marshal, &p)
	err = json.Unmarshal([]byte(personJson), &p)
	if err != nil {
		log.Error("decode error:", err)
	}
	fmt.Printf("decocd result:%#v", p)

}

type Person struct {
	Name   string `json:"name"`
	Age    int
	Gender int
}
