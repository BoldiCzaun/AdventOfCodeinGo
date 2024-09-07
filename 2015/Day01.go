package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day01_main() {
	fmt.Println("The puzzle input:")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lines, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	line := lines[0]

	fmt.Printf("The input is: %s\n", line)

	fmt.Printf("Part 1:\nThe final floor is: %d\n", D01_p1(line))
	PrintDiv()
	fmt.Printf("Part 2:\nThe first position to get to level -1 is: %d\n", D01_p2(line))
}

func D01_p1(line string) (floor int) {
	floor = 0

	for _, c := range line {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	return
}

func D01_p2(line string) int {
	// fmt.Println("Part 2:")
	floor := 0

	for i, c := range line {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}
	panic("Didn't reach floor -1")
}
