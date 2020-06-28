package main

import "fmt"

func main() {
	printIt(42)
	printIt("a string")
	printIt([]float64{1.2, 2.3, 1.12345679})
}

func printIt(v interface{}) {
	// We must refer to t in each case, not v!
	switch t := v.(type) { // HL
	case int: // HL
		fmt.Printf("Int = %d\n", t)
	case string: // HL
		fmt.Printf("String = %s\n", t)
	default:
		fmt.Printf("Something = %v\n", t)
	}
}
