package main

import "fmt"

func main() {
	ints := []int{0, 41, 42, 43, 400}
	for _, i := range ints {
		fmt.Printf("%d is %s\n", i, howClose(i))
	}
}

func howClose(i int) string {
	switch i { // HL
	case 41, 43: // HL
		return "close" // HL
	case 42: // HL
		return "the answer" // HL
	default: // HL
		return "not even close" // HL
	} // HL
}
