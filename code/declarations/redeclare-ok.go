package main

import (
	"log"
	"os"
)

func main() {
	wd1, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(wd1)

	wd2, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(wd2)
}
