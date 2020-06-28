package main

import (
	"os"

	"github.com/pborman/uuid" // HL
)

func main() {
	id := uuid.NewUUID() // HL
	os.Stdout.WriteString(id.String() + "\n")
}
