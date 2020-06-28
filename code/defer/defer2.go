package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeText() // OMIT
	file, err := os.Open("some-text")
	if err != nil {
		panic(err)
	}
	defer file.Close() // HL
	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func writeText() {
	t := `Line 1
Line 2
And now we're done
`
	if err := ioutil.WriteFile("some-text", []byte(t), 0644); err != nil {
		panic(err)
	}
}
