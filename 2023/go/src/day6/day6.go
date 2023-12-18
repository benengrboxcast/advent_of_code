package main

import (
	"advent2023/pkg/numerics"
	"fmt"
	"math"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
)

func waysToWin(time float64, dist float64) float64 {

	sq := math.Sqrt(time*time - 4*(dist))
	r1 := (-time + sq) / (-2)
	// It's possible this is a tie
	if r1*(time-r1) == dist {
		r1 += 0.1
	}
	middle := time / 2

	if r1 > middle {
		return (math.Floor(r1)-middle)*2 + 1
	} else {
		return (middle-math.Ceil(r1))*2 + 1
	}
}

func part1(lines []string) float64 {

	var times []float64
	var dists []float64
	index := 0

	for index < len(lines[0]) {
		var val int
		val, index = numerics.GetNumeric(lines[0], index)
		times = append(times, float64(val))
	}

	index = 0
	for index < len(lines[1]) {
		var val int
		val, index = numerics.GetNumeric(lines[1], index)
		dists = append(dists, float64(val))
	}

	ways := waysToWin(times[0], dists[0])
	for i := 1; i < len(dists); i++ {
		ways = ways * waysToWin(times[i], dists[i])
	}
	return ways
}

func part2(lines []string) float64 {
	return waysToWin(55826490.0, 246144110121111.0)
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
