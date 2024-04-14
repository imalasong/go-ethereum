package main

import "fmt"

func main() {

	definedMap()
}

func definedMap() {

	var m1 map[string]int

	m2 := map[string]int{}
	m3 := map[string]int{"k1": 11}
	m4 := make(map[string]int)
	m5 := make(map[string]int, 2)

	//m1["aa"] = 11  //error  nil map assigment

	m2["22"] += 4

	fmt.Printf("%T,%v,%v,%v\n", m1, m1, len(m1), m1 == nil)
	fmt.Printf("%T,%v,%v,%v\n", m2, m2, len(m2), m2 == nil)
	fmt.Printf("%T,%v,%v,%v\n", m3, m3, len(m3), m3 == nil)
	fmt.Printf("%T,%v,%v\n", m4, m4, m4 == nil)
	fmt.Printf("%T,%v,%v\n", m5, m5, m5 == nil)

	if v2, ok := m2["yy"]; ok {
		fmt.Printf("存在", v2)
	} else {
		fmt.Printf("不存在")
	}
}
