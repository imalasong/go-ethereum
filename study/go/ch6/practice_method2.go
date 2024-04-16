package main

import "fmt"

func main() {
	p := Person{
		Name: "imalasong",
		S:    []int{1, 2, 3},
		A:    [2]int{1, 2},
		M:    map[int]int{1: 1, 2: 2},
	}
	fmt.Printf("%#v\n", p)
	//p.PlanM()
	p.PlanP()
	//(&p).PlanP()
	fmt.Printf("%#v\n", p)
}

type Person struct {
	Name string
	S    []int
	A    [2]int
	M    map[int]int
}

func (p Person) PlanM() {
	p.Name = "55"
	p.S[0] = 0
	p.S = append(p.S, 1)
	p.S[1] = 0
	p.A[0] = 9
	p.M[0] = 9
	fmt.Printf("p=%T,Name=%T,S=%T,A=%T,M=%T\n", p, p.Name, p.S, p.A, p.M)
}
func (p *Person) PlanP() {
	p.Name = "55"
	p.S[0] = 0
	p.S = append(p.S, 1)
	p.S[1] = 0
	p.A[0] = 9
	p.M[0] = 9
	fmt.Printf("p=%T,Name=%T,S=%T,A=%T,M=%T\n", p, p.Name, p.S, p.A, p.M)
}
