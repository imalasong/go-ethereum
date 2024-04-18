package main

import (
	"bytes"
	"strconv"
	"sync"
)

type ISet interface {
	Add(elem uint64)
	Remove(elem uint64)
	Has(elem uint64) bool
	Clear()
	String() string
}

type SetInt struct {
	data []uint64
	sync.Mutex
}

func (s *SetInt) Add(elem uint64) {
	//s.Lock()
	//defer s.Unlock()
	if s.data == nil {
		s.data = make([]uint64, 1)
	}
	index, bit := int(elem/64), elem%64
	if index >= len(s.data) {
		ss := make([]uint64, index-len(s.data)+1)
		s.data = append(s.data, ss...)
	}
	s.data[index] |= 1 << bit
}

func (s *SetInt) Remove(elem uint64) {
	//s.Lock()
	//defer s.Unlock()
	if s.data == nil {
		return
	}
	index, bit := int(elem/64), elem%64
	if s.data[index] == 0 {
		return
	}
	//位清空&^：
	//对于 a &^ b -->
	//对于b的每个数值：如果是0，取a对应位的数；如果是1，结果位是0
	s.data[index] &^= 1 << bit
}

func (s *SetInt) Has(elem uint64) bool {
	//s.Lock()
	//defer s.Unlock()
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

func (s *SetInt) String() string {
	//s.Lock()
	//defer s.Unlock()
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
