package main

import (
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
	lines, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}
	D02_p1(lines)

}

func D02_p1(lines []string) {
	fmt.Println("Part 1")

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

	sum := 0
	for _, b := range boxes {
		//surface
		sum += 2 * (b.l*b.w + b.w*b.h + b.h*b.l)
		// slack
		sum += min(b.l*b.w, b.w*b.h, b.h*b.l)
	}

	fmt.Printf("%d sqrft of paper should be ordered\n", sum)
}
