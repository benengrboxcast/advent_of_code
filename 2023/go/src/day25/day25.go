package main

import (
	"fmt"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
)

func part1(lines []string) int {
	return 1
}

func part2(lines []string) int {
	return 2
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
