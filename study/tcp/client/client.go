package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	dial, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		fmt.Println("连接失败：", err)
	}

	go func() {
		for {
			bytes := make([]byte, 1024)
			n, err2 := dial.Read(bytes)
			if err2 == io.EOF {
				time.Sleep(time.Second)
				break
			}
			if n > 0 {
				fmt.Println("收到数据：", string(bytes[:n]))
			}
		}
	}()
	dial.Write([]byte("hello a"))

	time.Sleep(time.Hour)
}
