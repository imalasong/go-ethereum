package main

import "fmt"

func main() {

	var a, b, c = true, true, false

	var s string
	switch {
	case a:
		s = `22`
	case b:
		s = `333`
	case c:
		s = `444`
	default:
		s = `666`
	}

	fmt.Println(s)

}
