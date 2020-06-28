package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := sendSomeInts()
	c2 := sendSomeInts()
	timeout := time.After(100 * time.Millisecond)

Outer:
	for {
		select {
		case i := <-c1:
			fmt.Printf("Channel 1 says %d\n", i)
		case i := <-c2:
			fmt.Printf("Channel 2 says %d\n", i)
		case <-timeout:
			fmt.Println("Timed out")
			break Outer
		}
	}
}

// end-of-code - OMIT

func sendSomeInts() <-chan int {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	return c
}
