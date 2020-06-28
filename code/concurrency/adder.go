package main

import "fmt"

func main() {
	pairs := [][2]int{{1, 3}, {2, 5}, {-1, 2}}
	c := make(chan int)
	for _, p := range pairs {
		go addPair(p, c) // Pass each pair to a new goroutine
	}

	total := 0
	// We are done when we've received one result for each pair
	for i := 0; i < len(pairs); i++ {
		total += <-c
	}
	fmt.Println(total)
}

func addPair(p [2]int, out chan<- int) {
	out <- p[0] + p[1] // Write the sum to the channel
}
