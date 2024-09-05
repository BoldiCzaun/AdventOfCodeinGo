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

	D01_p1(line)
	D01_p2(line)
}

func D01_p1(line string) {
	fmt.Println("Part 1:")
	floor := 0

	for _, c := range line {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	fmt.Printf("The final floor is: %d\n", floor)

	PrintDiv()
}

func D01_p2(line string) {
	fmt.Println("Part 2:")
	floor := 0

	for i, c := range line {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}

		if floor == -1 {
			fmt.Printf("The first position to get to level -1 is: %d\n", (i + 1))
			break
		}
	}

	PrintDiv()
}
