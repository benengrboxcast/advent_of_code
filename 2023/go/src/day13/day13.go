package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
)

func checkHorizontalReflection(lines []string, index int, errorCount int) bool {
	// We assume that index and index - 1 already match
	top := index - 2
	bottom := index + 1
	for top >= 0 && bottom < len(lines) {
		for col := 0; col < len(lines[0]); col++ {
			if lines[top][col] != lines[bottom][col] {
				if errorCount == 0 {
					return false
				} else {
					errorCount--
				}
			}
		}
		top--
		bottom++
	}
	return errorCount == 0
}

func findHorizontalReflection(lines []string, errorCount int) (int, error) {
	for row := 1; row < len(lines); row++ {
		valid := true
		thisError := errorCount
		for col := 0; col < len(lines[0]); col++ {
			if lines[row-1][col] != lines[row][col] {
				if thisError == 0 {
					valid = false
					break
				} else {
					thisError--
				}
			}
		}
		if valid {
			if checkHorizontalReflection(lines, row, thisError) {
				return row, nil
			}
		}
	}

	return -1, errors.New("no horizontal reflection found")
}

func checkVerticalReflection(lines []string, column int, errorCount int) bool {
	// We assume that col - 1 is equal to col
	left := column - 2
	right := column + 1

	for left >= 0 && right < len(lines[0]) {
		for line := 0; line < len(lines); line++ {
			if lines[line][left] != lines[line][right] {
				if errorCount == 0 {
					return false
				} else {
					errorCount--
				}
			}
		}
		left--
		right++
	}

	return errorCount == 0
}

func findVertialReflection(lines []string, errorCount int) (int, error) {
	for col := 1; col < len(lines[0]); col++ {
		valid := true
		thisError := errorCount
		for row := 0; row < len(lines); row++ {
			if lines[row][col] != lines[row][col-1] {
				if thisError == 0 {
					valid = false
					break
				} else {
					thisError--
				}
			}
		}
		if valid {
			if checkVerticalReflection(lines, col, thisError) {
				return col, nil // The column INDEX that's to the left of the reflection
				// line is col - 1 and the COLUMN NUMBER is col
			}
		}
	}

	return -1, errors.New("no vertical reflection found")
}

func getReflection(input []string, errorCount int) (int, error) {
	t, err := findHorizontalReflection(input, errorCount)
	if err != nil {
		t, err = findVertialReflection(input, errorCount)
		if err != nil {
			return -1, err

		} else {
			return t, nil
		}
	} else {
		return t * 100, nil
	}
}

func createInputs(lines []string) [][]string {
	var current []string
	var result [][]string

	for _, line := range lines {
		if len(line) == 0 {
			temp := make([]string, len(current))
			copy(temp, current)

			result = append(result, temp)
			current = current[0:0]
		} else {

			current = append(current, line)
		}
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}

func part1(lines []string, errorCount int) int {
	/*
		A couple of notes
		1. Brute force would take way too long O(n^n) it think.  However as we go through the input as
		   normal can count the mismatches, when we find a match that has 0 expected errors, we found
		   the valid reflection
		2. There _should_ be two elements that can be flipped.  In the examples, if the mirrored item
		   is flipped it would have also resulted in a valid reflection
		3.
	*/
	result := 0
	inputs := createInputs(lines)

	for i, input := range inputs {
		t, err := getReflection(input, errorCount)
		if err != nil {
			fmt.Printf("Failed to get reflection on input %d\n", i)
		}
		result += t
	}

	return result
}

func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)

	start := time.Now()
	result := part1(output, 0)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	result = part1(output, 1)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %d. It took %s\n", result, elapsed)
}
