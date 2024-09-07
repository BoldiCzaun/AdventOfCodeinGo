package day03

import (
	utils "aoc2015/utils"
	"fmt"
	"log"
)

func Day03_main() {
	lines, err := utils.ReadInput()
	if err != nil {
		log.Panic(err)
	}
	input := lines[0]
	fmt.Print(input)
}
