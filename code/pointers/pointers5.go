package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{name: "Thomas Senlin", age: 25}
	takesStruct(tom)
	fmt.Println(tom.name) // Name is unchanged
	takesPointer(&tom)
	fmt.Println(tom.name) // Name is now "Ringo Shiina"
}

func takesStruct(p person) { // HL
	p.name = "Ringo Shiina"
}

func takesPointer(p *person) { // HL
	p.name = "Ringo Shiina"
}
