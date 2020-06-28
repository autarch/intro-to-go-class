package main

import "fmt"

func main() {
	a := []string{"a", "slice", "of", "strings"}
	fmt.Printf("a = %v\n", a)

	// Go does not modify a in place!
	b := append(a, "is", "longer", "now")
	fmt.Println("b := append(a, ...)")
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)

	// But you can re-assign to a variable
	a = append(a, "is", "also", "longer")
	fmt.Println("a = append(a, ...)")
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
}
