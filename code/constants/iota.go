package main

import "fmt"

const (
	Sunday = iota // incremented by 1 for each constant
	Monday
	Tuesday
	// ...
)

const (
	x = iota // resets to 0 in new const block
	y
)

func main() {
	fmt.Printf("Sunday = %d\n", Sunday)
	fmt.Printf("Tuesday = %d\n", Tuesday)
	fmt.Printf("y = %d\n", y)
}
