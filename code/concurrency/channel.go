package main

import "fmt"

func main() {
	c := make(chan bool)
	for _, i := range []int{2, 4, 6, 8} {
		go printInt(i, c) // HL
	}

	for i := 0; i < 4; i++ {
		<-c
	}
}

func printInt(i int, c chan<- bool) { // HL
	fmt.Println(i)
	c <- true // HL
}
