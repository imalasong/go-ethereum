package main

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {

	var i int = 1

	fmt.Printf("%T,%v\n", i, i)
	fmt.Printf("%T,%v\n", &i, &i)

	ip1 := &i
	fmt.Println("----------------------------")
	fmt.Printf("%T,%v\n", ip1, ip1)
	fmt.Printf("%T,%v\n", &ip1, &ip1)

	ip2 := &ip1
	fmt.Println("----------------------------")
	fmt.Printf("%T,%v\n", ip2, ip2)
	fmt.Printf("%T,%v\n", &ip2, &ip2)
	fmt.Printf("%T,%v\n", *ip2, *ip2)
	fmt.Printf("%T,%v\n", **ip2, **ip2)

}
