package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{0, 42, -1, 10002, -124858124, 9}
	sort.Ints(a)
	fmt.Println(a)

	b := []string{"word", "salad", "你好","breakfast", "time"}
	sort.Strings(b)
	fmt.Println(b)
}
