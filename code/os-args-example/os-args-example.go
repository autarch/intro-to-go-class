package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for i, val := range os.Args {
			fmt.Printf("arg #%d = %s\n", i, val)
		}
	} else {
		fmt.Println("No args provided")
	}
}


