package day05

import (
	"aoc2015/utils"
	"fmt"
	"log"
	"strings"
)

func Main() {
	lines, err := utils.ReadInput("Day05/input.txt")
	if err != nil {
		log.Panic(err)
	}

	nice_strings := 0
	nice_strings2 := 0
	for _, line := range lines {
		// fmt.Print(line)
		if d05_p1(line) {
			nice_strings++
		}
		if d05_p2(line) {
			nice_strings2++
		}
	}
	fmt.Printf("Part 1:\nInput has %d nice strings\n", nice_strings)
	utils.PrintDiv()
	fmt.Printf("Part 2:\nInput has %d nice strings\n", nice_strings2)
}

func d05_p1(input string) (is_nice bool) {
	input = strings.ToLower(input)
	is_nice = has_three_vowels(input)
	is_nice = is_nice && has_double_character(input)
	is_nice = is_nice && !has_forbiden_string(input)
	return
}

func d05_p2(input string) (is_nice bool) {
	input = strings.ToLower(input)
	is_nice = true
	is_nice = has_pair_of_pairs(input)
	is_nice = is_nice && has_repeat(input)
	return
}

func has_three_vowels(input string) bool {
	return (strings.Count(input, "a") + strings.Count(input, "e") + strings.Count(input, "i") + strings.Count(input, "o") + strings.Count(input, "u")) >= 3
}

func has_double_character(input string) bool {
	for c := 'a'; c <= 'z'; c++ {
		st := string([]rune{c, c})
		if strings.Contains(input, st) {
			return true
		}
	}
	return false
}

func has_forbiden_string(input string) bool {
	forbiden_strings := []string{"ab", "cd", "pq", "xy"}
	for _, st := range forbiden_strings {
		if strings.Contains(input, st) {
			return true
		}
	}
	return false
}

func has_pair_of_pairs(input string) bool {
	for c := 'a'; c <= 'z'; c++ {
		for d := 'a'; d <= 'z'; d++ {
			st := string([]rune{c, d})
			if strings.Count(input, st) >= 2 {
				return true
			}
		}
	}
	return false
}

func has_repeat(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false
}
