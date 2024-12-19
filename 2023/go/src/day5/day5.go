package main

import (
	"advent2023/pkg/numerics"
	"cmp"
	"fmt"
	"path/filepath"
	"slices"
	"sort"
	"time"

	"advent2023/pkg/file"
)

type MapRange struct {
	Source int
	Dest   int
	Size   int
}

func stringToMapRange(line string) MapRange {
	index := 0
	var result MapRange
	result.Dest, index = numerics.GetNumeric(line, index)
	result.Source, index = numerics.GetNumeric(line, index+1)
	result.Size, index = numerics.GetNumeric(line, index+1)

	return result
}

func combineMaps(first []MapRange, second []MapRange) []MapRange {
	result := []MapRange{}

	for si := 0; si < len(second); si++ {
		overlap := false
		for fi := 0; fi < len(first); fi++ {
			if first[fi].Dest == second[si].Source {
				result = append(result, MapRange{Source: first[fi].Source, Dest: second[si].Dest, Size: first[fi].Size})
				overlap = true
			}
		}
		if !overlap {
			result = append(first, second...)
		}
	}

	return result
}

func mapInputs(inputs []int, maps []MapRange) []int {
	slices.SortFunc(maps, func(a, b MapRange) int {
		return cmp.Compare(a.Source, b.Source)
	})

	for i := 0; i < len(inputs); i++ {
		for mapIndex := 0; mapIndex < len(maps); mapIndex++ {
			m := maps[mapIndex]
			if m.Source > inputs[i] {
				break
			} else if inputs[i] >= m.Source && inputs[i] < (m.Source+m.Size) {
				inputs[i] = inputs[i] + (m.Dest - m.Source)
				break
			}
		}
	}
	return inputs
}

func part1(lines []string) int {
	index := 6
	var sources []int
	for index < len(lines[0]) {
		var current int
		current, index = numerics.GetNumeric(lines[0], index)
		sources = append(sources, current)
		index++
	}

	lineIndex := 3
	var maps []MapRange
	for lineIndex < len(lines) {
		if len(lines[lineIndex]) == 0 || lines[lineIndex][0] == ' ' {
			lineIndex += 2
			sources = mapInputs(sources, maps)
			maps = maps[:0]
		} else {
			maps = append(maps, stringToMapRange(lines[lineIndex]))
			lineIndex++
		}
	}

	sort.Ints(sources)
	return sources[0]
}

func part2(lines []string) int {
	index := 6
	var sources []int
	for index < len(lines[0]) {
		var current int
		current, index = numerics.GetNumeric(lines[0], index)
		var count int
		count, index = numerics.GetNumeric(lines[0], index)
		for i := 0; i < count; i++ {
			sources = append(sources, current+i)
		}
	}

	lineIndex := 3
	var maps []MapRange
	for lineIndex < len(lines) {
		if len(lines[lineIndex]) == 0 || lines[lineIndex][0] == ' ' {
			lineIndex += 2
			sources = mapInputs(sources, maps)
			maps = maps[:0]
		} else {
			maps = append(maps, stringToMapRange(lines[lineIndex]))
			lineIndex++
		}
	}

	sort.Ints(sources)
	return sources[0]
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
