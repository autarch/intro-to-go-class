package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	dave := returnsPointer()
	// We can treat the pointer just like we treat the struct
	fmt.Println(dave.name) // HL
}

func returnsPointer() *person {
	dave := person{
		name: "Dave Rolsky",
		age:  41,
	}

	return &dave // HL
}
