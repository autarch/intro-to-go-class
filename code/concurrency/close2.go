package main

import "fmt"

func main() {
	c1 := sendSomeInts()

	for {
		i, ok := <-c1
		if ok {
			fmt.Printf("Channel says %d\n", i)
		} else {
			break
		}
	}
}

func sendSomeInts() <-chan int {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	return c
}

// end-of-code - OMIT
