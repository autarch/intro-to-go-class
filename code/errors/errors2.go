package main

import (
	"errors"
	"fmt"
)

func main() {
	// Common Go idiom
	if n, err := numCheck(41); err != nil { // HL
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}

func numCheck(n int) (int, error) {
	if n == 42 {
		// We still need to return *something* for the error - we cannot just
		// "return n"
		return n, nil
	}
	return 0, errors.New("This is not 42!")
}
