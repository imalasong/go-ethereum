package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4}

	//for _, v := range sl {
	//	fmt.Println(v)
	//	if v < 10 {
	//		sl = append(sl, v*100)
	//		fmt.Println("aaaaaaaa")
	//	}
	//	fmt.Printf("length:%d\n", len(sl))
	//}

	for i := 0; i < len(sl); i++ {
		v := sl[i]
		fmt.Println(v)
		if v < 10 {
			sl = append(sl, v*100)
		}
	}
}
