package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := sendWord(0)
	c2 := sendWord(1)
	fmt.Println(<-c1, <-c2)
}

var words []string // HL

func sendWord(x int) <-chan string {
	var once sync.Once
	c := make(chan string)
	go func() {
		once.Do(func() { // HL
			words = expensiveOperation() // HL
		}) // HL
		c <- words[x]
	}()
	return c
}

// end-of-code - OMIT

func expensiveOperation() []string {
	return []string{
		"dog",
		"cat",
		"pig",
		"goat",
	}
}
