package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := returnsPointer()
	// We can treat the pointer just like we treat the struct
	fmt.Println(tom.name) // HL
}

func returnsPointer() *person {
	tom := person{
		name: "Thomas Senlin",
		age:  25,
	}

	return &tom // HL
}
