package day06

import (
	"aoc2015/utils"
	"errors"
	"fmt"
	"log"
	"strings"
)

func Main() {
	lines, err := utils.ReadInput("Day06/input.txt")
	if err != nil {
		log.Panic(err)
	}

	count, err := d06_p1(lines)
	if err != nil {
		panic(err)
	}
	count2, err := d06_p2(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1:\nAfter the sequence %d lights are on\n", count)
	utils.PrintDiv()
	fmt.Printf("Part 2:\nAfter the sequence %d is the sum brightness\n", count2)

}

func d06_p1(input []string) (int, error) {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, line := range input {
		var command string
		var from_x, from_y, to_x, to_y int

		// Adjust the format string to handle different command formats
		if strings.HasPrefix(line, "turn on") {
			_, err := fmt.Sscanf(line, "turn on %d,%d through %d,%d", &from_x, &from_y, &to_x, &to_y)
			if err != nil {
				return 0, fmt.Errorf("failed to parse 'turn on' command: %w", err)
			}
			command = "turn on"
		} else if strings.HasPrefix(line, "turn off") {
			_, err := fmt.Sscanf(line, "turn off %d,%d through %d,%d", &from_x, &from_y, &to_x, &to_y)
			if err != nil {
				return 0, fmt.Errorf("failed to parse 'turn off' command: %w", err)
			}
			command = "turn off"
		} else if strings.HasPrefix(line, "toggle") {
			_, err := fmt.Sscanf(line, "toggle %d,%d through %d,%d", &from_x, &from_y, &to_x, &to_y)
			if err != nil {
				return 0, fmt.Errorf("failed to parse 'toggle' command: %w", err)
			}
			command = "toggle"
		} else {
			return 0, fmt.Errorf("unknown command: %s", line)
		}

		switch command {
		case "turn on":
			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid[x][y] = true
				}
			}
		case "turn off":
			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid[x][y] = false
				}
			}
		case "toggle":
			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid[x][y] = !grid[x][y]
				}
			}
		default:
			return 0, errors.New("invalid input")
		}
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				count++
			}
		}
	}
	return count, nil
}

func d06_p2(input []string) (int, error) {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range input {
		var command string
		var from_x, from_y, to_x, to_y int

		// Adjust the format string to handle different command formats
		if strings.HasPrefix(line, "turn on") {
			_, err := fmt.Sscanf(line, "turn on %d,%d through %d,%d", &from_x, &from_y, &to_x, &to_y)
			if err != nil {
				return 0, fmt.Errorf("failed to parse 'turn on' command: %w", err)
			}
			command = "turn on"
		} else if strings.HasPrefix(line, "turn off") {
			_, err := fmt.Sscanf(line, "turn off %d,%d through %d,%d", &from_x, &from_y, &to_x, &to_y)
			if err != nil {
				return 0, fmt.Errorf("failed to parse 'turn off' command: %w", err)
			}
			command = "turn off"
		} else if strings.HasPrefix(line, "toggle") {
			_, err := fmt.Sscanf(line, "toggle %d,%d through %d,%d", &from_x, &from_y, &to_x, &to_y)
			if err != nil {
				return 0, fmt.Errorf("failed to parse 'toggle' command: %w", err)
			}
			command = "toggle"
		} else {
			return 0, fmt.Errorf("unknown command: %s", line)
		}

		switch command {
		case "turn on":
			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid[x][y]++
				}
			}
		case "turn off":
			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid[x][y]--
					grid[x][y] = max(grid[x][y], 0)
				}
			}
		case "toggle":
			for x := from_x; x <= to_x; x++ {
				for y := from_y; y <= to_y; y++ {
					grid[x][y] += 2
				}
			}
		default:
			return 0, errors.New("invalid input")
		}
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			count += grid[i][j]
		}
	}
	return count, nil
}
