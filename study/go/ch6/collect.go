package main

import (
	"bytes"
	"fmt"
)

func main() {
	set := new(SetInt)
	set.Add(1)
	fmt.Println(set.Has(1))
	fmt.Println(set.Has(2))
	fmt.Println(set.Has(3))
	fmt.Println(set.Has(4))
	fmt.Println(set.Has(1))
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

func (s *SetInt) Has(elem uint64) bool {
	if s.data == nil {
		return false
	}
	index := int(elem / 64)
	indexData := s.data[index]
	bit := elem % 64
	return indexData|1<<bit == 1
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
		for j := 0; i < 64; j++ {

		}
	}

	buffer.WriteByte('}')
	return buffer.String()
}
