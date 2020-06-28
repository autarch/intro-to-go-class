package main

import "fmt"

type Printer interface { // HL
	Println() // Anything that provides a Println method implements the Printer interface
}

type MyString string

func (m MyString) Println() { fmt.Println(m) } // MyString implements the Printer interface

type MyInt int64

func (m MyInt) Println() { fmt.Println(m) } // and so does MyInt

func main() {
	var s MyString = "foo"
	var i MyInt = 42
	p(s)
	p(i)
}

func p(x Printer) { x.Println() } // Takes an arg which implements Printer
