package main

import (
	"errors"
	"log"
	"os"

	lang "cantlang.org/cant/lang"
)

func main() {
	log.SetPrefix("CANT: ")

	if len(os.Args) < 2 {
		log.Fatal(errors.New("Missing input file argument!"))
	}

	arg1 := os.Args[1]
	log.Println("input:", arg1)

	lang.Cant(arg1)
}
