package main

import "fmt"

func main() {
	fmt.Print("a string\n")
	fmt.Print(42, 0, 4.3)
	fmt.Println("adds a newline")
	fmt.Println(42, 0, 4.3)
	fmt.Printf("%.2f %v\n", 42.123, []int64{1, 2, 42, -1234})
}
