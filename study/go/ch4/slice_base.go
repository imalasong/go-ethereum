package main

import "fmt"

func main() {

	definedSlice()
}

func definedSlice() {

	var s []int
	s2 := []int{}
	s3 := []int{1}
	s4 := make([]int, 0, 2)
	a := [4]int{}
	s5 := a[0:2]
	s6 := new([]int)

	// s[0] = 1 compiler error
	s = append(s, 1)

	// s2[0] = 1 compiler error
	s2 = append(s2, 1)
	s3[0] = 1
	// s4[0] = 1 compiler error
	s4 = append(s4, 1)
	s5[0] = 1

	// (*s6)[0] = 1 compiler error
	*s6 = append(*s6, 1)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)

}
