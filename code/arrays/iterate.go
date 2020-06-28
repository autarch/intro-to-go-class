package main

import "fmt"

func main() {
	a := []string{"My", "fancy", "slice!"}
	// range loops over each index and value
	for i, v := range a {
		fmt.Printf("a[%d] = %s\n", i, v)
	}

	// Ignore the index
	for _, v := range a {
		fmt.Println(v)
	}
}
