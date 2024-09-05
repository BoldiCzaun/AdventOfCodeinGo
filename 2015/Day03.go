package main

import (
	"fmt"
	"log"
)

func Day03_main() {
	lines, err := ReadInput()
	if err != nil {
		log.Panic(err)
	}
	input := lines[0]
	fmt.Print(input)
}
