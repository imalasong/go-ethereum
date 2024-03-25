package p1

import "fmt"

type book struct {
	Name string
}

func init() {
	fmt.Println("sssssssssssssssssssss2 init")
}

func CreateBook() *book {
	return &book{}
}
