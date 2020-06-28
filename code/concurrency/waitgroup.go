package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // HL
	urls := []string{"http://blog.urth.org/", "http://www.google.com"}
	for _, url := range urls {
		// Adds 1 to the count of goroutines we're waiting for
		wg.Add(1) // HL
		go func(url string) {
			// Done() decrements the counter
			defer wg.Done() // HL
			fetchURL(url)
		}(url)
	}
	// Blocks until the internal counter of goroutines is 0
	wg.Wait() // HL
}

// end-of-code - OMIT

func fetchURL(url string) {
	fmt.Println("Fetched", url)
}
