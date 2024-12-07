package main

import (
	"advent2023/pkg/numerics"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"advent2023/pkg/file"
)

func HowManyWays(groups [][]byte, required []int) int {

	// If the groups are equal to requred, then just find how many ways to create each one
	if len(groups) == 0 {
		return 0
	}

	if len(required) == 0 {
		return 1
	}

	//if len(groups) == len(required) {
	//	result := 1
	//	for i := 0; i < len(groups); i++ {
	//		result = result * CountContiguousUnknowns(len(groups[i]), required[i:i+1])
	//	}
	//	return result
	//}

	firstGear := -1
	for i := 0; i < len(groups[0]); i++ {
		if groups[0][i] == '#' {
			firstGear = i
			break
		}
	}

	// If we have one group and they are all unknown, calculate
	if firstGear == -1 && len(groups) == 1 {
		return CountContiguousUnknowns(len(groups[0]), required)
	} else if len(groups[0]) < required[0] {
		return HowManyWays(groups[1:], required)
	} else if len(groups[len(groups)-1]) < required[len(required)-1] {
		return HowManyWays(groups[0:len(groups)-1], required)
	}

	// See if we can find a required way to break up the beginning
	if firstGear > -1 && firstGear < required[0]+1 {
		// The first valid gear has to be part of the first group
		var lastGear = firstGear + 1
		for ; lastGear < len(groups[0]) && groups[0][lastGear] == '#'; lastGear++ {
		}
		usable := required[0] - (lastGear - firstGear)
		startIndex := firstGear - usable
		if startIndex < 0 {
			startIndex = 0
		}
		result := 0

		for i := startIndex; i <= firstGear; i++ {
			endIndex := i + required[0] // The is the exclusive end of this group (or the first in the next group)
			if endIndex > len(groups[0]) {
				continue
			}
			updatedGroups := make([][]byte, len(groups))
			copy(updatedGroups, groups)
			// Where we are now + the number needed + 1 space
			beginningIndex := i + required[0] + 1
			if beginningIndex >= len(updatedGroups[0]) {
				updatedGroups = updatedGroups[1:]
			} else {
				updatedGroups[0] = groups[0][beginningIndex:]
			}
			result += HowManyWays(updatedGroups, required[1:])
		}
		return result
	}

	// See if we can remove something from the back
	lastGear := -1
	lastGroup := groups[len(groups)-1]
	for i := len(lastGroup) - 1; i >= 0; i-- {
		if lastGroup[i] == '#' {
			lastGear = i
			break
		}
	}

	needed := required[len(required)-1]
	lastGroupSize := len(lastGroup)
	if lastGear > -1 && (lastGear >= (lastGroupSize - needed)) {
		firstGear = lastGear - 1
		for ; firstGear > 0; firstGear-- {
			if lastGroup[firstGear] != '#' {
				break
			}
		}
		usable := needed - (lastGear - firstGear)
		startIndex := lastGear + usable
		if startIndex >= lastGroupSize {
			startIndex = lastGroupSize - 1
		}
		result := 0

		for i := startIndex; i >= lastGear; i-- {
			endSlice := i - needed
			//if endSlice < 0 || lastGroup[endSlice] == '#' {
			//	continue
			//}
			updateGroups := make([][]byte, len(groups))
			copy(updateGroups, groups)
			if endSlice <= 0 {
				// Just Remove the item
				updateGroups = updateGroups[0 : len(updateGroups)-1]
			} else {
				updateGroups[len(updateGroups)-1] = groups[len(groups)-1][0:endSlice]
			}

			result += HowManyWays(updateGroups, required[0:len(required)-1])
		}

		return result
	}

	// At this point, we need to place the first gear wherever it can go
	spacesAvailable := 0
	for _, group := range groups {
		spacesAvailable += len(group)
	}

	spacesNeeded := 0
	for _, r := range required {
		spacesNeeded += r
	}
	//result := 0
	//for groupIndex := 0; groupIndex < len(groups); groupIndex++ {
	//	if spacesNeeded > spacesAvailable {
	//		break
	//	}
	//	group := groups[groupIndex]
	//	spaceIndex := 0
	//	size := required[0]
	//	for spaceIndex < len(group) {
	//		// Can we place it at this index
	//		if spaceIndex > 0 && spa
	//	}
	//
	//}

	return 0
}

func getFirstSpring(groups [][]byte) (int, int) {
	// If we are not supposed to place anymore springs, then all of the input data must be unknowns
	for gi, group := range groups {
		for ci, char := range group {
			if char == '#' {
				return gi, ci
			}
		}
	}
	return -1, -1
}

func getLastSpring(groups [][]byte) (int, int) {
	for gi := len(groups) - 1; gi >= 0; gi-- {
		group := groups[gi]
		for ci := len(group) - 1; ci >= 0; ci-- {
			char := group[ci]
			if char == '#' {
				return gi, ci
			}
		}
	}
	return -1, -1
}

/*
This method tries to place the first required element in each available space and then calls itself to place the
remaining groups
*/
func bruteForce(groups [][]byte, required []int) int {
	// Check that this input is valid
	spacesAvailable := 0
	for _, group := range groups {
		spacesAvailable += len(group)
	}

	spacesNeeded := 0
	for _, r := range required {
		spacesNeeded += r
	}

	if spacesNeeded > spacesAvailable {
		return 0
	} else if spacesNeeded == spacesAvailable {
		result := 1
		if len(required) == len(groups) {
			for i := 0; i < len(required); i++ {
				if len(groups[i]) != required[i] {
					result = 0
					break
				}
			}
		} else {
			result = 0
		}
		return result
	}

	firstSpringGroup, firstSpringIndex := getFirstSpring(groups)

	// If we have a spring, but require 0, this is invalid
	if len(required) == 0 && firstSpringGroup >= 0 {
		return 0
	}

	if len(groups) == 1 && firstSpringIndex == -1 {
		return CountContiguousUnknowns(len(groups[0]), required)
	}

	if len(required) == 1 {
		// If there is only one required value, the placement has to encompass both the first and last spring

		// If there is NO spring add possibles for all groups
		if firstSpringGroup == -1 {
			result := 0
			for _, group := range groups {
				result += CountContiguousUnknowns(len(group), required)
			}
		} else {
			lastSpringGroup, lastSpringIndex := getLastSpring(groups)
			if lastSpringGroup != firstSpringGroup {
				return 0
			} else {
				start := lastSpringIndex - required[0] + 1
				if start < 0 {
					start = 0
				}
				if start > firstSpringIndex {
					return 0
				}

				lastIndex := firstSpringIndex                     // The last index that we can place this is either where the first spring starts, because it must be included
				end := len(groups[lastSpringGroup]) - required[0] // Or the index that would place this spring on the last element in the group
				if end < lastIndex {
					lastIndex = end
				}
				if lastIndex < 0 {
					return 0
				}

				return lastIndex - start + 1
			}
		}
	}

	// None of the stop conditions have been met, now start trying to place the first spring
	var needed int
	needed, required = required[0], required[1:]
	result := 0
	// We are only able to place the first required spring until we've reached the first spring index.
	lastGroupIndex := firstSpringGroup
	if lastGroupIndex < 0 {
		lastGroupIndex = len(groups) - 1
	}
	for gi := 0; gi <= lastGroupIndex; gi++ {
		group := groups[gi]
		endIndex := len(group)
		if gi == firstSpringGroup {
			endIndex = firstSpringIndex
		}

		for ci := 0; ci <= endIndex; ci++ {
			// Are we able to place
			if ci > (len(group) - needed) {
				break
			}

			if gi == firstSpringGroup && len(group) > ci+needed && group[ci+needed] == '#' {
				// If we placed the spring here, the last spring of this group would touch the first known spring
				continue
			}

			if len(required) == 0 {

				result += 1
			} else if len(group) > (ci + needed + 1) {
				updatedGroup := group[ci+needed+1:]
				updatedGroups := make([][]byte, len(groups[gi:]))
				copy(updatedGroups, groups[gi:])
				updatedGroups[0] = updatedGroup
				result += bruteForce(updatedGroups, required)
			} else {
				result += bruteForce(groups[gi+1:], required)
			}
		}
	}
	return result

}

func parseLine(line string) ([][]byte, []int) {
	var result [][]byte
	var current []byte
	var index int
	for index = 0; index < len(line); index++ {
		c := line[index]
		if c == '.' && len(current) > 0 {
			temp := make([]byte, len(current))
			copy(temp, current)
			result = append(result, temp)
			current = current[0:0]
		} else if c == ' ' {
			break
		} else if c != '.' {
			current = append(current, c)
		}
	}
	if len(current) > 0 {
		result = append(result, current)
	}

	var counts []int
	for index < len(line) {
		var count int
		count, index = numerics.GetNumeric(line, index)
		counts = append(counts, count)
		index++
	}

	return result, counts
}

func CountContiguousUnknowns(count int, groups []int) int {
	if len(groups) == 0 {
		return 0
	} else if len(groups) == 1 {
		return count - groups[0] + 1
	} else if len(groups) == 2 {
		n := count - groups[0] - groups[1]
		return n * (n + 1) / 2
	} else {
		result := 0

		required := 0
		for _, c := range groups[1:] {
			required += c + 1 // There needs to be a space as well
		}
		lastIndex := count - required - groups[0]
		for index := 0; index <= lastIndex; index++ {
			theirIndex := index + groups[0] + 1
			theirCount := count - theirIndex
			result += CountContiguousUnknowns(theirCount, groups[1:])
		}
		return result
	}
}

func part1(lines []string) int {
	result := 0
	for idx, line := range lines {

		groups, required := parseLine(line)
		current := bruteForce(groups, required)
		result += current
		fmt.Printf("%d: %s | %d\n", idx, line, current)

	}
	return result
}

func part2(lines []string) (int, error) {
	result := 0
	for idx, line := range lines {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			return 0, errors.New("Invalid Line")
		}

		springs := split[0]
		counts := split[1]
		for i := 0; i < 4; i++ {
			springs = springs + "?" + split[0]
			counts = counts + "," + split[1]
		}
		line = springs + " " + counts
		groups, required := parseLine(line)
		current := bruteForce(groups, required)
		result += current
		fmt.Printf("%d: %s | %d\n", idx, line, current)

	}
	return result, nil
}

func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)

	start := time.Now()
	// result := part1(output)
	result := part1(output)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	var err error
	result, err = part2(output)
	if err != nil {
		fmt.Printf("Failed with: %s", err.Error())
	}
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %d. It took %s\n", result, elapsed)
}
