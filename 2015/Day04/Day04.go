package day04

import (
	"aoc2015/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func Main() {
	lines, err := utils.ReadInput("day04/input.txt")
	if err != nil {
		log.Panic(err)
	}
	input := lines[0]

	p1, err := d04_p1(input)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Part 1:\nThe secret key is: %d\n", p1)
	utils.PrintDiv()

	p2, err := d04_p2(input)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Part 2:\nThe secret key is: %d\n", p2)
	utils.PrintDiv()
}

func d04_p1(input string) (int, error) {
	for i := 0; i < 10_000_000; i++ {
		str := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(str))
		hashString := hex.EncodeToString(hash[:])
		if hashString[0:5] == "00000" {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no solution below 10 000 000")
}

func d04_p2(input string) (int, error) {
	for i := 0; i < 10_000_000; i++ {
		str := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(str))
		hashString := hex.EncodeToString(hash[:])
		if hashString[0:6] == "000000" {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no solution below 10 000 000")
}
