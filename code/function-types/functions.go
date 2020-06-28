package main

import "fmt"

func main() {
	adder := func(a, b int) int { return a + b }      // HL
	multiplier := func(a, b int) int { return a * b } // HL

	// Passing two different functions to doMath
	fmt.Printf("2 + 3 = %d\n", doMath(adder, 2, 3))
	fmt.Printf("2 * 3 = %d\n", doMath(multiplier, 2, 3))
	// Creating and passing an anonymous function
	fmt.Printf("2 - 3 = %d\n", doMath(
		func(a, b int) int { return a - b }, 2, 3)) // HL
}

// doMath accepts a function as its first argument
func doMath(math func(a, b int) int, a, b int) int { // HL
	// Calling the passed function looks just like calling any other function
	return math(a, b) // HL
}
