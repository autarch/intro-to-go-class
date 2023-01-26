package main

import (
	"os"

	"github.com/google/uuid" // HL
)

func main() {
	id := uuid.New() // HL
	os.Stdout.WriteString(id.String() + "\n")
}
