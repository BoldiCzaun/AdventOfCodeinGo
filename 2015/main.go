package main

import (
	"aoc2015/day01"
	"aoc2015/day02"
	"aoc2015/day03"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide exactly one day")
	}
	dayNo, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		log.Fatal("Please provide a number")
	}

	switch dayNo {
	case 1:
		day01.Day01_main()
	case 2:
		day02.Day02_main()
	case 3:
		day03.Day03_main()
	default:
		log.Fatal("The day you have been asked for has not been implemented yet")
	}

}
