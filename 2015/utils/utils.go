package utils

import (
	"bufio"
	"fmt"
	"os"
)

type TestCase struct {
	Input    string
	Expected any
}

func NewTestCase(input string, expected any) TestCase {
	return TestCase{
		Input:    input,
		Expected: expected,
	}
}

// func ReadInput() ([]string, error) {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return nil, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var lines []string
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return nil, err
// 	}
// 	return lines, nil
// }

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
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
