package main

import (
	"fmt"
	"github.com/holiman/uint256"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	//NewInt(1)
	var i int = 1
	var i2 int64 = 1

	fmt.Printf("%T,%v,%d\n", i, i, unsafe.Sizeof(i))
	fmt.Printf("%T,%v,%d\n", i2, i2, unsafe.Sizeof(i2))
}

func Test2(t *testing.T) {

	i1 := math.MaxInt64
	i2 := i1 + 1
	i3 := i1 + 1

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)

	b1 := big.NewInt(math.MaxInt64)
	b1.Add(b1, big.NewInt(math.MaxInt64))
	fmt.Println(math.MaxInt64)
	fmt.Println(b1.String())

	u1 := uint256.NewInt(math.MaxInt64)
	u1.Add(u1, uint256.NewInt(math.MaxInt64))
	fmt.Println(u1.String())

	fmt.Println("-----------------------------------")

	var f1 float64 = 143.66
	var f2 float64 = 14.55
	f3 := f1 - f2
	fmt.Println(f3)

	f4, ok := decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).Float64()
	if ok {
		fmt.Println(f4)
	}
	fmt.Println("00000000000")
	fmt.Println(decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).BigFloat())

	ff1 := big.NewFloat(f1)
	ff1 = ff1.Sub(ff1, big.NewFloat(f2))
	fmt.Println(ff1)
}
