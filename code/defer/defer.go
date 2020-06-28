package main

import "fmt"

func main() {
	func1()
	func2()
}

func func1() {
	defer fmt.Println("func1 deferred")
	fmt.Println("func1")
}

func func2() {
	defer func3()
	defer func() { fmt.Println("func2 deferred") }()
	fmt.Println("func2")
}

func func3() {
	fmt.Println("func3")
}
