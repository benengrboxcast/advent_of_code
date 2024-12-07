package main

import (
	"advent2023/pkg/file"
	"advent2023/pkg/numerics"
	"fmt"
	"path/filepath"
	"time"
)

func findNextValue(values []float64) float64 {
	if len(values) == 2 {
		return values[1] + values[1] - values[0]
	}

	digit := values[1] - values[0]
	done := true
	var next []float64
	for i := 1; i < len(values); i++ {
		delta := values[i] - values[i-1]
		next = append(next, delta)
		if done && delta != digit {
			done = false
		}
	}

	if done {
		return values[len(values)-1] + digit
	} else {
		return values[len(values)-1] + findNextValue(next)
	}
}

func findFirstValue(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	digit := values[1] - values[0]
	done := true
	var next []float64
	for i := 1; i < len(values); i++ {
		delta := values[i] - values[i-1]
		next = append(next, delta)
		if done && delta != digit {
			done = false
		}
	}

	if done {
		return values[0] - digit
	} else {
		return values[0] - findFirstValue(next)
	}
}

func part1(lines []string) float64 {
	result := 0.0
	for _, line := range lines {
		index := 0
		var values []float64
		for index < len(line) {
			var v int
			v, index = numerics.GetNumeric(line, index)
			index++
			values = append(values, float64(v))
		}
		result += findNextValue(values)
	}
	return result
}

func part2(lines []string) float64 {
	result := 0.0
	for _, line := range lines {
		index := 0
		var values []float64
		for index < len(line) {
			var v int
			v, index = numerics.GetNumeric(line, index)
			index++
			values = append(values, float64(v))
		}
		result += findFirstValue(values)
	}
	return result
}
func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)

	start := time.Now()
	result := part1(output)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %f. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %f. It took %s\n", result, elapsed)
}
