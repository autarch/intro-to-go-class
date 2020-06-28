package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	dave := person{name: "Dave Rolsky", age: 41}
	takesStruct(dave)
	fmt.Println(dave.name) // Name is unchanged
	takesPointer(&dave)
	fmt.Println(dave.name) // Name is now "Ringo Shiina"
}

func takesStruct(p person) { // HL
	p.name = "Ringo Shiina"
}

func takesPointer(p *person) { // HL
	p.name = "Ringo Shiina"
}
