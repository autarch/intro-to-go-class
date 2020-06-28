package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// Makes a pointer to the struct
	davePtr := &person{ // HL
		name: "Dave Rolsky",
		age:  41,
	}

	fmt.Println(davePtr.name)
}
