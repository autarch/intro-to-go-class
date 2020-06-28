package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	dave := person{
		name: "Dave Rolsky",
		age:  41,
	}
	// We cannot pass "dave", we need to take a pointer
	takesPointer(&dave) // HL
}

func takesPointer(p *person) { // HL
	fmt.Println(p.name)
}
