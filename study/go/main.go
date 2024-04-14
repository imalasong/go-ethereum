package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	//var s1 []string = []string{"a", "b", "c", "d"}
	//var s1 []string = make([]string, 1, 5)
	//s1 = append(s1, "a")
	//fmt.Printf("befor real value:%s,-----%d,%d\n", s1, len(s1), cap(s1))
	////addItem(s1, 1)
	//modifyItem(s1, "t")
	//fmt.Printf("after real value:%s,-----%d,%d\n", s1, len(s1), cap(s1))
	//

	unsaferOpe()
}

func unsaferOpe() {
	s1 := []int{}
	s1 = append(s1, 1, 2, 3)
	//var s1 []string = []string{"a", "b", "c"}
	fmt.Printf("befor real value:%s,-----%d,%d\n", s1, len(s1), cap(s1))
	addItemInt(s1, 1)
	fmt.Printf("after real value:%s,-----%d,%d\n", s1, len(s1), cap(s1))

	p := unsafe.Pointer(&s1[2])
	q := uintptr(p) + 8
	t := (*int)(unsafe.Pointer(q))
	fmt.Println(*t)
}

func addItem(s1 []string, count int) {
	for i := 0; i < count; i++ {
		s1 = append(s1, strconv.Itoa(i))
	}
	fmt.Printf("inner value:%s,-----%d,%d\n", s1, len(s1), cap(s1))
}

func addItemInt(s1 []int, count int) {
	for i := 0; i < count; i++ {
		s1 = append(s1, count)
	}
	fmt.Printf("inner value:%s,-----%d,%d\n", s1, len(s1), cap(s1))
}
func modifyItem(s1 []string, item string) {
	s1[0] = item
	fmt.Printf("inner value:%s,-----%d,%d\n", s1, len(s1), cap(s1))
}
