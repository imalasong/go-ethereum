package p1

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	a    book
}

func init() {
	fmt.Println("sssssssssssssssssssss1 init")
}

func (s *Student) A() {
	fmt.Printf("current %v,value=%v\n", reflect.TypeOf(s), "A method")
}

func (s Student) B() {
	fmt.Printf("current %v,value=%v\n", reflect.TypeOf(s), "B method")
}

func C() {
	fmt.Printf("%v\n", "A method")
}

func (s Student) String() string {
	return fmt.Sprintf("name = %v", s.Name)
}
