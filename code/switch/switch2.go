package main

import "fmt"

func main() {
	ints := []int{0, 41, 42, 43, 400}
	for _, i := range ints {
		fmt.Printf("%d is %s\n", i, howClose(i))
	}
}

func howClose(i int) string {
	// No variable after the "switch"
	switch { // HL
	case i == 41 || i == 43: // HL
		return "close" // HL
	case i == 42: // HL
		return "the answer" // HL
	default: // HL
		return "not even close" // HL
	} // HL
}
