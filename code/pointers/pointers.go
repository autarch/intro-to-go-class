package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func main() {
	tom := person{
		name: "Thomas Senlin",
		age:  25,
	}

	tomPtr := &tom // HL
	fmt.Printf("tomPtr is a %s to a %s\n",
		reflect.TypeOf(tomPtr).Kind(),
		reflect.ValueOf(tomPtr).Elem().Type().Name())
}
