package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

func Test(t *testing.T) {
	var set ISet = new(SetInt)
	set.Add(0)
	set.Add(10000)
	set.Add(10)
	set.Add(2)
	set.Add(0)
	set.Add(0)
	set.Remove(0)
	set.Remove(2)
	//set.Remove(10)
	set.Remove(1)

	set.Add(100)
	set.Add(500)
	set.Add(100)

	fmt.Println(set.Has(1))
	fmt.Println(set.Has(2))
	fmt.Println(set.Has(3))
	fmt.Println(set.Has(4))
	fmt.Println(set.Has(10))
	fmt.Println(set)
	fmt.Println(set.String())
	set.Clear()

}

func Test22(t *testing.T) {
	// 有序数据：1000000000
	//SetInt ：
	//comparable complete:2.082528/s
	//tostring complete:91.340366/s
	//SetInt2:
	//comparable complete:23.871638/s
	//tostring complete:92.677847/s

	//无序数据集：1000000000
	//SetInt ：
	//comparable complete:38.694504/s
	//tostring complete:137.799969/s
	//SetInt2:
	//comparable complete:188.250404/s

	var set ISet = new(SetInt)
	var i uint64 = 0
	start := time.Now()
	for ; i < 20000000; i++ {
		d := uint64(rand.Uint32())
		set.Add(d)
		if !set.Has(d) {
			fmt.Printf("error digital:%d\n", d)
		}
	}
	fmt.Printf("comparable complete:%f/s\n", time.Since(start).Seconds())
	start = time.Now()
	//fmt.Println("result:" + set.String())
	set.String()
	fmt.Printf("tostring complete:%f/s,size=%d\n", time.Since(start).Seconds(), unsafe.Sizeof(set))
}

func TestMake(t *testing.T) {
	var index uint64 = 2353586950434264312
	ss := make([]uint64, index-1, index-1)
	fmt.Println(ss)
}
