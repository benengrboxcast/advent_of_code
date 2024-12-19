package main

import (
	"fmt"
	"math"
	"path/filepath"
	"slices"
	"sync"
	"time"

	"advent2023/pkg/file"
	"advent2023/pkg/numerics"
)

func parseWinningNumbers(line string) ([]int, int) {
	idx := 6

	// Put the index on the first number of the 'winning' numbers
	for {
		if line[idx] == ':' {
			idx += 2
			break
		}
		idx++
	}

	var rval []int
	for {
		var thisNumber int
		thisNumber, idx = numerics.GetNumeric(line, idx)
		idx++
		rval = append(rval, thisNumber)
		if line[idx] == '|' {
			idx += 2
			break
		}
	}

	slices.Sort(rval)
	return rval, idx
}

func part1Partial(lines []string, startIndex int, endIndex int) int {
	result := 0
	for i := startIndex; i < endIndex; i++ {
		win, index := parseWinningNumbers(lines[i])
		count := 0.0
		for {
			var current int
			current, index = numerics.GetNumeric(lines[i], index)
			index++
			if slices.Contains(win, current) {
				count += 1
			}

			if index >= len(lines[i]) {
				if count > 0 {
					result += int(math.Pow(2, count-1))
				}

				break
			}
		}
	}
	return result
}

func part1(lines []string) int {
	return part1Partial(lines, 0, len(lines))
}

func part1MultiThread(lines []string) int {
	var p1 int
	var p2 int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		p1 = part1Partial(lines, 0, len(lines)/2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		p2 = part1Partial(lines, len(lines)/2, len(lines))
	}()

	wg.Wait()

	return p1 + p2
}

func part2(lines []string) int {
	cardCounts := make(map[int]int)
	result := 0
	for i := 0; i < len(lines); i++ {
		cardCounts[i]++
		win, index := parseWinningNumbers(lines[i])
		count := 0
		for {
			var current int
			current, index = numerics.GetNumeric(lines[i], index)
			index++
			if slices.Contains(win, current) {
				count += 1
			}

			if index >= len(lines[i]) {
				thisCount := cardCounts[i]
				for copyIndex := 0; copyIndex < count; copyIndex++ {
					cardCounts[i+copyIndex+1] += thisCount
				}
				result += thisCount
				break
			}
		}
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
	result = part1MultiThread(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 1 (Multithread) done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %d. It took %s\n", result, elapsed)
}
