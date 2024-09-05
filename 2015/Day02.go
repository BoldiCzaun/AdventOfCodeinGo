package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Now you can use the 'lines' array which contains the multiline input
	// fmt.Println(lines)

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

func part_1(