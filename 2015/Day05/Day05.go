package day05

import (
	"aoc2015/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Main() {
	lines, err := utils.ReadInput("day05/input.txt")
	if err != nil {
		log.Panic(err)
	}

	nice_strings := 0
	for _, line := range lines {
		// fmt.Print(line)
		if d05_p1(line) {
			nice_strings++
		}
	}
	fmt.Printf("Part 1:\nInput has %d nice strings\n", nice_strings)
	utils.PrintDiv()
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
			ts := string([]rune{d, c})
			if c == d && strings.Count(input, st) >= 3 {
				return true
			}
			if c != d && strings.Count(input, st) >= 2 && !strings.Contains(input, ts) {
				return true
			}
		}
	}
	return false
}

func has_repeat(input string) bool {
	pattern := `([a-z]).([a-z])`
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}
