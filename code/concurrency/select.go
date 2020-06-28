// This example was mostly borrowed from
// http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := schedule(50)
	c2 := schedule(100)
	for i := 1; i <= 6; i++ {
		select {
		case msg := <-c1:
			fmt.Printf("Channel 1 says %s\n", msg)
		case msg := <-c2:
			fmt.Printf("Channel 2 says %s\n", msg)
		}
	}
}

func schedule(delay int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Millisecond * time.Duration(delay))
			c <- fmt.Sprintf("message %d (after %d milliseconds)", i, delay*i)
		}
	}()
	return c
}

// end-of-code - OMIT
