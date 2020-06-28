package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := make([]int64, 4)
	a[0] = 42
	a[1] = 12345678
	fmt.Printf("%v (%s)\n", a, reflect.TypeOf(a).Kind())
}
