package main

import "fmt"

func main() {
	a := map[string]int{
		"foo": 42,
		"bar": -1,
	}
	fmt.Printf("a has %d elements\n", len(a))
	fmt.Printf("%v\n", a)

	a["baz"] = 99
	fmt.Printf("%v\n", a)

	delete(a, "foo")
	fmt.Printf("%v\n", a)

	if v, found := a["baz"]; found {
		fmt.Printf("baz = %d\n", v)
	} else {
		fmt.Println("no baz!")
	}
}
