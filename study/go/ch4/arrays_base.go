package main

import "fmt"

func main() {

	//defined();

	a1 := [...]int{1, 2}
	a2 := [...]int{1, 2}

	fmt.Println(a1 == a2)
	a1[0] = 2
	fmt.Println(a1 == a2)

	a3 := [3]int{1, 2, 3}

	a4 := a3[0:2]
	fmt.Printf("%T,%v,%v\n", a4, a4, a3)
}

func defined() {
	var a1 [3]int
	a2 := [3]int{}
	a3 := [...]int{}
	a4 := new([3]int)

	fmt.Printf("%T,%v\n", a1, a1)
	fmt.Printf("%T,%v\n", a2, a2)
	fmt.Printf("%T,%v\n", a3, a3)
	fmt.Printf("%T,%v\n", a4, a4)
}
