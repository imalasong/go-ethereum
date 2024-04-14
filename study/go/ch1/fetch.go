package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Error("Fetch error:", err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Error("read body error:", err)
		}
		fmt.Printf("%s", body)
	}

}
