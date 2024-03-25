package main

import (
	"fmt"
	"reflect"
)

func main() {

	var b Blanka = &User{}
	b.Fly()
	fmt.Println(reflect.TypeOf(b))

}

type Blanka interface {
	Fly()
}

func (u *User) Fly() {
	fmt.Println("ffffffffffffffff")
}

type User struct {
	Name string
	age  int
}
