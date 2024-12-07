package main

import (
	"fmt"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
)

func part1(lines []string) int {

	firstRowCost := len(lines)
	result := 0
	for col := 0; col < len(lines[0]); col++ {
		rockCount := 0
		lastStationary := 0
		colCost := 0
		for row := 0; row < len(lines); row++ {
			if lines[row][col] == '#' {
				rockCount = 1
				lastStationary = row
			} else if lines[row][col] == 'O' {
				colCost += firstRowCost - lastStationary - rockCount
				rockCount++
			}
		}
		result += colCost
	}
	return result
}

func rotateNorth(lines [][]byte) [][]byte {
	for col := 0; col < len(lines[0]); col++ {
		var empties []int
		for row := 0; row < len(lines); row++ {
			c := lines[row][col]
			if c == '.' {
				empties = append(empties, row)
			} else if c == '#' {
				empties = empties[:0]
			} else if c == 'O' && len(empties) > 0 {
				var dest int
				dest, empties = empties[0], empties[1:]
				lines[dest][col] = 'O'
				lines[row][col] = '.'
				empties = append(empties, row)
			}
		}
	}
	return lines
}

func rotateSouth(lines [][]byte) [][]byte {
	for col := len(lines[0]) - 1; col >= 0; col-- {
		var empties []int
		for row := len(lines) - 1; row >= 0; row-- {
			c := lines[row][col]
			if c == '.' {
				empties = append(empties, row)
			} else if c == '#' {
				empties = empties[:0]
			} else if c == 'O' && len(empties) > 0 {
				var dest int
				dest, empties = empties[0], empties[1:]
				lines[dest][col] = 'O'
				lines[row][col] = '.'
				empties = append(empties, row)
			}
		}
	}
	return lines
}

func rotateWest(lines [][]byte) [][]byte {
	for row := 0; row < len(lines); row++ {
		var empties []int
		for col := 0; col < len(lines[0]); col++ {
			c := lines[row][col]
			if c == '.' {
				empties = append(empties, col)
			} else if c == '#' {
				empties = empties[:0]
			} else if c == 'O' && len(empties) > 0 {
				var dest int
				dest, empties = empties[0], empties[1:]
				lines[row][dest] = 'O'
				lines[row][col] = '.'
				empties = append(empties, col)
			}
		}
	}
	return lines
}

func rotateEast(lines [][]byte) [][]byte {
	for row := 0; row < len(lines); row++ {
		var empties []int
		for col := len(lines[0]) - 1; col >= 0; col-- {
			c := lines[row][col]
			if c == '.' {
				empties = append(empties, col)
			} else if c == '#' {
				empties = empties[:0]
			} else if c == 'O' && len(empties) > 0 {
				var dest int
				dest, empties = empties[0], empties[1:]
				lines[row][dest] = 'O'
				lines[row][col] = '.'
				empties = append(empties, col)
			}
		}
	}
	return lines
}

func gridToInt(lines [][]byte) int {
	result := 0
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			result = result << 2
			if lines[row][col] == 'O' {
				result = result | 1
			} else if lines[row][col] == '#' {
				result = result | 2
			}
		}
	}
	return result
}

var cache = make(map[int][][]byte)

func transform(lines [][]byte) [][]byte {
	lines = rotateNorth(lines)
	lines = rotateWest(lines)
	lines = rotateSouth(lines)
	lines = rotateEast(lines)
	return lines
}

func printGrid(lines [][]byte) {
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}

func cost(lines [][]byte) int {
	firstRowCost := len(lines)
	result := 0
	for col := 0; col < len(lines[0]); col++ {
		for row := 0; row < len(lines); row++ {
			if lines[row][col] == 'O' {
				result += firstRowCost - row
			}
		}
	}
	return result
}
func part2(lines []string) int {
	data := make([][]byte, len(lines))
	for i, line := range lines {
		data[i] = []byte(line)
	}
	current := gridToInt(data)
	for i := 0; i < 1000000000; i++ {
		data = transform(data)
		next := gridToInt(data)
		if current == next {
			fmt.Printf("No changes on line %d\n", i)
			printGrid(data)
			return cost(data)
		}
		current = next
		if i%1000000 == 0 {
			fmt.Printf("Cycle %d\n", i)
		}
	}
	return cost(data)
}

func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)
	fmt.Printf("Grid size is %d x %d\n", len(output[0]), len(output))
	start := time.Now()
	result := part1(output)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %d. It took %s\n", result, elapsed)
}
