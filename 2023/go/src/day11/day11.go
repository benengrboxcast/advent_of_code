package main

import (
	"advent2023/pkg/file"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type Position struct {
	x int
	y int
}

func (p Position) distance(other Position) int {
	var xs [2]int
	if p.x < other.x {
		xs[0] = p.x
		xs[1] = other.x
	} else {
		xs[0] = other.x
		xs[1] = p.x
	}

	var ys [2]int
	if p.y < other.y {
		ys[0] = p.y
		ys[1] = other.y
	} else {
		ys[0] = other.y
		ys[1] = p.y
	}

	return xs[1] - xs[0] + ys[1] - ys[0]
}

func findDistances(lines []string, expandCost int) int {
	// Expand the galaxy
	// Do columns firs because that harder
	var expandCols []int
	for col := 0; col < len(lines[0]); col++ {
		expand := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] == '#' {
				expand = false
			}
		}
		if expand {
			expandCols = append(expandCols, col)
		}
	}

	// Expand Rows
	var expandRows []int
	for row := 0; row < len(lines); row++ {
		if !strings.Contains(lines[row], "#") {
			expandRows = append(expandRows, row)
		}
	}

	var galaxies []Position
	rowIndex := 0
	expandRowIndex := 0
	for row := 0; row < len(lines); row++ {
		if expandRowIndex < len(expandRows) && row == expandRows[expandRowIndex] {
			expandRowIndex++
			rowIndex += expandCost
			continue
		}

		colIndex := 0
		expandColIndex := 0
		r := lines[row]
		for col := 0; col < len(r); col++ {
			if expandColIndex < len(expandCols) && expandCols[expandColIndex] == col {
				expandColIndex++
				colIndex += expandCost
				continue
			}
			if r[col] == '#' {
				galaxies = append(galaxies, Position{colIndex, rowIndex})
			}
			colIndex++

		}
		rowIndex++
	}

	result := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += galaxies[i].distance(galaxies[j])
		}
	}

	return result
}
func part1(lines []string) int {
	return findDistances(lines, 2)
}

func part2(lines []string) int {
	return findDistances(lines, 1000000)
}

func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)

	start := time.Now()
	result := part1(output)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %d. It took %s\n", result, elapsed)
}
