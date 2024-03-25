package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	assert.True(t, 1 == 1)
}

func TestGoruting(t *testing.T) {
	var c chan int = make(chan int, 2)
	go func() {
		c <- 0
		time.Sleep(time.Second * 2)
		c <- 1
	}()
	_, ok := <-c
	if ok {
		fmt.Println("111")
	}
	_, ok = <-c
	if ok {
		fmt.Println("2222")
	}
	//c <- 1
	//c <- 2
	//fmt.Println(<-c)
	//c <- 3
	//
	//fmt.Println(<-c)
}
