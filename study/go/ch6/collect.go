package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	set := new(SetInt)
	set.Add(1)
	set.Add(10)
	set.Add(10)
	set.Add(2)
	set.Add(0)
	set.Add(0)
	set.Remove(0)
	set.Remove(2)
	set.Remove(10)
	set.Remove(1)

	fmt.Println(set.Has(1))
	fmt.Println(set.Has(2))
	fmt.Println(set.Has(3))
	fmt.Println(set.Has(4))
	fmt.Println(set.Has(1))
	//fmt.Println(*set)
	//fmt.Println(set)
	set.Clear()
	fmt.Println(set.ToString())
}

type SetInt struct {
	data []uint64
}

func (s *SetInt) Add(elem uint64) {
	if s.data == nil {
		s.data = make([]uint64, 1)
	}
	index, bit := int(elem/64), elem%64
	if index > len(s.data) {
		s.data = append(s.data, make([]uint64, index-len(s.data)+1)...)
	}
	s.data[index] |= 1 << bit
}

func (s *SetInt) Remove(elem uint64) bool {
	if s.data == nil {
		return false
	}
	index, bit := int(elem/64), elem%64
	if s.data[index] == 0 {
		return false
	}
	//位清空&^：
	//对于 a &^ b -->
	//对于b的每个数值：如果是0，取a对应位的数；如果是1，结果位是0
	s.data[index] &^= 1 << bit
	return true
}

func (s *SetInt) Has(elem uint64) bool {
	if s.data == nil {
		return false
	}
	index := int(elem / 64)
	indexData := s.data[index]
	bit := elem % 64
	return indexData&(1<<bit) != 0
}

func (s *SetInt) Clear() {
	s.data = nil
}

func (s *SetInt) ToString() string {
	if s.data == nil {
		return ""
	}
	buffer := new(bytes.Buffer)
	buffer.WriteByte('{')
	for i, elem := range s.data {
		if elem == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if elem&(1<<j) != 0 {
				d := i*64 + j
				buffer.WriteString(strconv.Itoa(d) + " ")
			}
		}
	}
	buffer.WriteByte('}')
	return buffer.String()
}
