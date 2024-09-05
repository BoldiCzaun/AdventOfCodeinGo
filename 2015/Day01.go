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
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The input is: %s\n", line)

	part_1(line)
	part_2(line)
}

func part_1(line string) {
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
	for i := 0; i < 80; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func part_2(line string) {
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

	for i := 0; i < 80; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
