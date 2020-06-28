package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	// Creating a struct
	person := Person{ // HL
		firstName: "Ringo",
		lastName:  "Shiina",
	}
	// Accessing its fields
	fmt.Printf("Name = %s %s\n", person.firstName, person.lastName) // HL
}
