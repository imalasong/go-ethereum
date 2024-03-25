package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/study/base/p1"
)

func main() {

	s := p1.Student{}
	fmt.Println(s)
	s.A()
	s.B()

	p1.C()

	book := p1.CreateBook()

	book.Name = "11"

}
