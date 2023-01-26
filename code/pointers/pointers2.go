package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{
		name: "Thomas Senlin",
		age:  25,
	}
	// We cannot pass "tom", we need to take a pointer
	takesPointer(&tom) // HL
}

func takesPointer(p *person) { // HL
	fmt.Println(p.name)
}
