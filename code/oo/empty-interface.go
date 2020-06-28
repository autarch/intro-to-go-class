package main

import "fmt"

type user struct {
	username string
	email    string
}

func main() {
	printIt(42)
	printIt("a string")
	printIt([]float64{1.2, 2.3, 1.12345679})
	printIt(user{"ringo", "ringo@shiina.com"})
}

// The empty interface type accepts anything
func printIt(v interface{}) { // HL
	fmt.Printf("%v\n", v)
}
