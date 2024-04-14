package main

import (
	"fmt"
	"math/big"
)

func main() {

	op := NewEmployee()
	plan := NewEmployee2()
	op.ID = "1"
	plan.ID = "1"
	fmt.Println(op)
	fmt.Println(plan)

	//w := Wheels{}
	//w := Wheels{Pointer{"1","1"},Circular{Pointer{"10","1"},1.1},1.1}
	w := Wheels{Pointer: Pointer{"1", "1"}, Circular: Circular{Pointer{"10", "1"}, 1.1}, Speed: 1.1, lowerCase: lowerCase{"11"}}
	w.Y = "11.111"
	w.Pointer.Y = "111111"
	fmt.Printf("%v\n", w)
	fmt.Printf("%#v\n", w)
}

func NewEmployee() *Employee {

	return new(Employee)
}

func NewEmployee2() Employee {

	return *new(Employee)
}

func NewEmployee3() *Employee {

	//return &Employee{"01","11","11",1,2,2,false,*big.NewInt(1),nil}
	return &Employee{ID: "111", Name: "imalasong"}
}

type Employee struct {
	ID string

	Name, Address string

	Salary int

	Position float64

	BirthDay int64

	delFlag bool

	Saving big.Int

	/**  reference self only pointer */
	e *Employee
}

type Pointer struct {
	X, Y string
}

type Circular struct {
	Pointer
	Radius float64
}

type Wheels struct {
	Pointer
	Circular
	Speed float64
	lowerCase
}

type lowerCase struct {
	Name string
}
