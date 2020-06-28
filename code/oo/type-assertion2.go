package main

import "fmt"

func main() {
	printIt(42)
	printIt("a string")
	printIt([]float64{1.2, 2.3, 1.12345679})
}

func printIt(v interface{}) {
	// If the value in v cannot be treated as an int this will cause a panic
	i := v.(int) // HL
	fmt.Printf("Int = %d\n", i)
}
