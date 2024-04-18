package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestM1(t *testing.T) {
	var a1 m1 = 1
	var a2 m2 = 1
	fmt.Printf("%v,%T,%v\n", a1, a1, a2.toString())
}

type m1 int
type m2 int

func (m m1) P1() {

}

//func (m *m1) P1() {
//
//}

func (m m2) P1() {

}

func (m m2) toString() string {
	return "number:" + strconv.Itoa(int(m))
}
