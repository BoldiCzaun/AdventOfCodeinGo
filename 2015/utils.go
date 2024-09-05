package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide exactly one day")
	}
	dayNo, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		log.Fatal("Please provide a number")
	}

	switch dayNo {
	case 1:
		Day01_main()
	case 2:
		Day02_main()
	default:
		log.Fatal("The day you have been asked for has not been implemented yet")
	}

}

func ReadInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return lines, nil
}

func PrintDiv() {
	for i := 0; i < 80; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
