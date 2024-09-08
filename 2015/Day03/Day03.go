package day03

import (
	utils "aoc2015/utils"
	"fmt"
	"log"
)

func Day03_main() {
	lines, err := utils.ReadInput("day03/input.txt")
	if err != nil {
		log.Panic(err)
	}
	input := lines[0]
	// fmt.Print(input)
	fmt.Printf("Part 1:\nNumber of houses visited: %d\n", d03_p1(input))
	utils.PrintDiv()
}

type coordinate struct {
	x int
	y int
}

func d03_p1(input string) int {
	visited := map[coordinate]bool{}
	curr := coordinate{0, 0}
	visited[curr] = true
	for _, d := range input {
		switch d {
		case '<':
			curr.x--
		case '>':
			curr.x++
		case '^':
			curr.y++
		case 'v':
			curr.y--
		}
		if _, exists := visited[curr]; !exists {
			visited[curr] = true
		}
	}

	return len(visited)
}
