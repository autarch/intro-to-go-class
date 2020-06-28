package main

import (
	"fmt"
	"time"
)

func main() {
	for _, i := range []int{2, 4, 6, 8} {
		go func(n int) { // HL
			fmt.Println(n) // HL
		}(i) // HL
	}
	time.Sleep(1000 * time.Millisecond)
}
