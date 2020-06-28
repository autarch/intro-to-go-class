package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [4]string{"an", "array", "of", "strings"}
	fmt.Printf("a = %v (%s)\n", a, reflect.TypeOf(a).Kind())

	b := a[1:]
	fmt.Printf("b = %v (%s)\n", b, reflect.TypeOf(b).Kind())

	c := b[:1]
	fmt.Printf("c = %v (%s)\n", c, reflect.TypeOf(c).Kind())

	d := []string{"a", "slice", "of", "strings"}
	fmt.Printf("d = %v (%s)\n", d, reflect.TypeOf(d).Kind())
}
