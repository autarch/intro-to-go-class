package main

import "fmt"

func main() {
	a := 2
	// Closes over a
	adderClosure := func(b int) int { return a + b } // HL

	fmt.Printf("a + 2 = %d\n", adderClosure(2))

	// Changes to a are seen by the closure
	a++ // HL

	fmt.Printf("a + 2 = %d\n", adderClosure(2))
}
