package main

import (
	"bytes"
	"strconv"
)

type SetInt2 struct {
	data map[uint64]uint64
}

func (s *SetInt2) Add(elem uint64) {
	if s.data == nil {
		s.data = make(map[uint64]uint64, 10)
	}
	index, bit := elem/64, elem%64

	s.data[index] |= 1 << bit
}
func (s *SetInt2) Remove(elem uint64) {
	index, bit := elem/64, elem%64
	if s.data[index] == 0 {
		return
	}
	s.data[index] &^= 1 << bit
}
func (s *SetInt2) Has(elem uint64) bool {
	index, bit := elem/64, elem%64
	if s.data[index] == 0 {
		return false
	}
	return s.data[index]&(1<<bit) != 0
}
func (s *SetInt2) Clear() {
	s.data = nil
}
func (s *SetInt2) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('{')

	for index, v := range s.data {
		if v == 0 {
			continue
		}
		var i uint64 = 0
		for ; i < 64; i++ {
			if v&(1<<i) != 0 {
				d := index*64 + i
				buf.WriteString(strconv.Itoa(int(d)) + " ")
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}
