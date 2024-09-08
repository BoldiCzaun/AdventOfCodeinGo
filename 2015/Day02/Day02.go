package day02

import (
	utils "aoc2015/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type box struct {
	l int
	w int
	h int
}

func newBox(l, w, h int) box {
	return box{
		l: l,
		w: w,
		h: h,
	}
}

func Day02_main() {
	lines, err := utils.ReadInput("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1:\n%d sqrft of paper should be ordered\n", d02_p1(lines))
	utils.PrintDiv()
	fmt.Printf("Part 2:\n%d ft of ribbon is needed\n", d02_p2(lines))
	utils.PrintDiv()
}

func d02_p1(lines []string) (sum int) {
	boxes := parse(lines)
	sum = 0
	for _, b := range boxes {
		//surface
		sum += 2 * (b.l*b.w + b.w*b.h + b.h*b.l)
		// slack
		sum += min(b.l*b.w, b.w*b.h, b.h*b.l)
	}
	return
}

func d02_p2(lines []string) (sum int) {
	boxes := parse(lines)
	sum = 0
	for _, b := range boxes {
		ribbon := (b.l + b.h + b.w - max(b.l, b.h, b.w)) * 2
		bow := b.l * b.h * b.w
		sum += ribbon + bow
	}

	return
}

func parse(lines []string) []box {
	var boxes []box

	for _, line := range lines {
		params := strings.Split(line, "x")
		l, err1 := strconv.Atoi(params[0])
		w, err2 := strconv.Atoi(params[1])
		h, err3 := strconv.Atoi(params[2])
		if err1 != nil || err2 != nil || err3 != nil {
			panic("Invalid input")
		}
		boxes = append(boxes, newBox(l, w, h))
	}

	return boxes
}
