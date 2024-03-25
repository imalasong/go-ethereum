package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:9090")

	if err != nil {
		fmt.Println("服务启动失败", err)
	}

	defer listen.Close()
	fmt.Println("服务启动成功", listen.Addr().String())
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("连接异常:", err)
		}
		go handler(accept)
	}

}

func handler(accept net.Conn) {
	defer accept.Close()

	bytes := make([]byte, 1024)
	n, err := accept.Read(bytes)
	if err != nil {
		fmt.Println("读数据异常：", err)
	} else {
		msg := string(bytes[:n])
		fmt.Println("接受到数据：", msg)
		for i := 1; i <= 10; i++ {
			accept.Write([]byte("hello"))
			time.Sleep(time.Second)
		}
	}
}
