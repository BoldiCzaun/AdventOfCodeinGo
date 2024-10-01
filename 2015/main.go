package main

import (
	"aoc2015/day01"
	"aoc2015/day02"
	"aoc2015/day03"
	"aoc2015/day04"
	"aoc2015/day05"
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
		day01.Main()
	case 2:
		day02.Main()
	case 3:
		day03.Main()
	case 4:
		day04.Main()
	case 5:
		day05.Main()
	case 6:
		day06.Main()
	default:
		log.Fatal("The day you have been asked for has not been implemented yet")
	}

}
