package main

import (
	"fmt"
	"time"
)

func main() {
	defer log("test")()
	defer func() {
		fmt.Println("22222222")
	}()
	defer func() {
		fmt.Println("3333333333333")
	}()
	defer func() {
		fmt.Println("4444444444444444444")
	}()

	fmt.Println("main run")
}

func log(modulo string) func() {
	now := time.Now()
	fmt.Printf("[%v]start run\n", modulo)

	return func() {
		fmt.Printf("[%v]end run,time:%d/mill\n", modulo, time.Since(now).Milliseconds())
	}
}
