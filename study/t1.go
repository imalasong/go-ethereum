package main

import "fmt"

type Usb interface {
	Say()
}
type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say()")
}
func main() {
	// 错误！ 会报 Stu类型没有实现Usb接口 ,
	// 如果希望通过编译,  var u Usb = &stu
	u := Stu{}
	u.Say()
	fmt.Println("here", u)
}
