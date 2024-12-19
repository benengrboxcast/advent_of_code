package main

import (
	"fmt"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
	"advent2023/pkg/numerics"
)

type SchematicNumber struct {
	Value      int
	StartIndex int
	EndIndex   int
}

func parse_line(line string) ([]bool, []SchematicNumber) {
	idx := 0

	var symbols []bool
	var numbers []SchematicNumber
	for idx < len(line) {
		if line[idx] == '.' {
			idx++
			symbols = append(symbols, false)
		} else if line[idx] >= 0x30 && line[idx] <= 0x39 {
			var value SchematicNumber
			value.StartIndex = idx
			value.Value, idx = numerics.GetNumeric(line, idx)
			value.EndIndex = idx - 1
			numbers = append(numbers, value)

			for i := value.StartIndex; i <= value.EndIndex; i++ {
				symbols = append(symbols, false)
			}

		} else {
			symbols = append(symbols, true)
			idx++
		}
	}

	return symbols, numbers
}

func parse_line_2(line string) ([]int, []SchematicNumber) {
	idx := 0

	var gears []int
	var numbers []SchematicNumber
	for idx < len(line) {
		if line[idx] == '*' {
			gears = append(gears, idx)
			idx++
		} else if line[idx] >= 0x30 && line[idx] <= 0x39 {
			var value SchematicNumber
			value.StartIndex = idx
			value.Value, idx = numerics.GetNumeric(line, idx)
			value.EndIndex = idx - 1
			numbers = append(numbers, value)

		} else {

			idx++
		}
	}

	return gears, numbers
}

func getMinMaxIdx(n SchematicNumber, max int) (int, int) {
	min_idx := 0
	if n.StartIndex-1 > min_idx {
		min_idx = n.StartIndex - 1
	}

	max_idx := max
	if n.EndIndex+1 < max_idx {
		max_idx = n.EndIndex + 1
	}
	return min_idx, max_idx
}

func part1(lines []string) int {
	rval := 0
	symbols := make(map[int][]bool)
	numbers := make(map[int][]SchematicNumber)

	for idx, line := range lines {
		s, n := parse_line(line)
		symbols[idx] = s
		numbers[idx] = n
	}

	for k, v := range numbers {
		for _, n := range v {
			found := false
			if n.StartIndex > 0 {
				if symbols[k][n.StartIndex-1] {
					found = true
				}
			}

			if !found && n.EndIndex < len(symbols[k])-1 {
				if symbols[k][n.EndIndex+1] {
					found = true
				}
			}

			if !found && k > 0 {
				min, max := getMinMaxIdx(n, len(symbols[k-1])-1)
				for i := min; i <= max; i++ {
					if symbols[k-1][i] {
						found = true
						break
					}
				}
			}

			if !found && k < len(symbols)-1 {
				min, max := getMinMaxIdx(n, len(symbols[k+1])-1)
				for i := min; i <= max; i++ {
					if symbols[k+1][i] {
						found = true
						break
					}
				}
			}

			if found {
				rval += n.Value
			}
		}
	}

	return rval
}

func findAdjacent(numbers []SchematicNumber, idx int) []int {
	var found []int
	for _, n := range numbers {
		// There are three times a number can be adjacent to a gear
		//  1. When it occupies the same index (StartIndex <= idx && EndIndex >= idx)
		//  2. It ends on the index before this gear (EndIndex == idx - 1)
		//  3. It start on the next index after this gear (StartIndex == idx + 1)
		if (n.StartIndex <= idx && n.EndIndex >= idx) || (n.EndIndex == (idx - 1)) || (n.StartIndex == (idx + 1)) {
			found = append(found, n.Value)
		}

	}
	return found
}

func part2(lines []string) int {
	var prevNumbers []SchematicNumber
	var thisNumbers []SchematicNumber
	var nextNumbers []SchematicNumber
	var gears []int
	var nextGears []int
	result := 0

	gears, thisNumbers = parse_line_2(lines[0])
	for i := 0; i < len(lines)-1; i++ {
		nextGears, nextNumbers = parse_line_2(lines[i+1])

		for _, gear := range gears {

			// Any in the previous line
			adjacent := findAdjacent(prevNumbers, gear)
			adjacent = append(adjacent, findAdjacent(thisNumbers, gear)...)
			adjacent = append(adjacent, findAdjacent(nextNumbers, gear)...)

			if len(adjacent) == 2 {
				result += (adjacent[0] * adjacent[1])
			}
		}

		prevNumbers = thisNumbers
		thisNumbers = nextNumbers
		gears = nextGears

	}
	return result
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
