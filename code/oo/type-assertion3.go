package main

import "fmt"

func main() {
	var i int64 = 42
	var f float64 = 42.0
	// Converts between two concrete types
	printIt(float64(i)) // HL
	printIt(f)
}

// v remains an interface{} throughout this func
func printIt(v interface{}) {
	// Asserts that an interface value is a specific type
	if f, ok := v.(float64); ok { // HL
		fmt.Printf("Float = %f\n", f)
		// This will never be true
	} else if i, ok := v.(int64); ok { // HL
		fmt.Printf("Int = %d\n", i)
	}
}
