package main

import (
	"fmt"
	"time"
)

func main() {
	offlineLogged := time.Now()
	fmt.Println(time.Since(offlineLogged).Seconds())

	testChan()

	strings := make(chan string, 2)
	go func() {
		strings <- "1"
		fmt.Println("写1")
		strings <- "1"
		fmt.Println("写2")
		fmt.Println("业务开始执行")
		time.Sleep(time.Second * 5)
		fmt.Println("业务执行完了")
		//strings <- "1"
	}()

	fmt.Println("执行到这了")
	<-strings
	fmt.Println("读1")
	<-strings
	fmt.Println("读2")
	<-strings
	fmt.Println("读3")
	time.Sleep(time.Second * 40)
}

func testChan() {

}
