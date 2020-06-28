package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	go sendSomeInts(c)
	for i := range c {
		fmt.Println(i)
	}
}

func sendSomeInts(c chan<- int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := r.Intn(10) + 1
	for i := 1; i <= count; i++ {
		c <- r.Intn(500) + 1
	}
	close(c) // HL
}

// end-of-code - OMIT
