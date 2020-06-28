package main

import "fmt"

type mystring string // HL

func main() {
	var m mystring
	m = "foo"
	m.meth()
}

func (m mystring) meth() { // HL
	fmt.Println(m)
}
