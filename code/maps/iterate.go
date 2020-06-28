package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["foo"] = 42
	a["bar"] = -1

	for k, v := range a {
		fmt.Printf("%s = %d\n", k, v)
	}

	// If you provide one variable you just get the keys.
	for k := range a {
		fmt.Printf("%s\n", k)
	}
}
