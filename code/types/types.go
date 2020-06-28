package main

import (
	"fmt"
	"reflect"
)

type UserID uint64

func main() {
	var i uint64 = 42

	t := reflect.TypeOf(i)
	fmt.Printf("i = %d (type = %s, kind = %s)\n", i, t.Name(), t.Kind())

	u := UserID(i) // HL
	t = reflect.TypeOf(u)
	fmt.Printf("u = %d (type = %s, kind = %s)\n", u, t.Name(), t.Kind())
}
