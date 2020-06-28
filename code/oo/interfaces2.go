package main

import (
	"errors"
	"fmt"
)

type FileError struct {
	filename string
	message  string
}

func main() {
	PrintError(errors.New("Some random error"))
	PrintError(FileError{filename: "path/to/file.txt", message: "File not found"})
}

func (h FileError) Error() string { return h.message }

// The first argument must implement the error interface
func PrintError(e error) {
	fmt.Println(e.Error())
}
