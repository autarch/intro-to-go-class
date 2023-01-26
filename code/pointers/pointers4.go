package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// Makes a pointer to the struct
	tomPtr := &person{ // HL
		name: "Thomas Senlin",
		age:  25,
	}

	fmt.Println(tomPtr.name)
}
