package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"advent2023/pkg/file"
	"advent2023/pkg/pattern_matcher"
)

func parse_str(v string) int {
	count := len(v)
	forward_found := false
	backward_found := false
	forward_val := 0
	backward_val := 0

	for i := 0; i < count; i++ {
		if !forward_found && v[i] >= 0x30 && v[i] <= 0x39 {
			forward_found = true
			forward_val = int(v[i] - 0x30)
		}

		if !backward_found && v[count-i-1] >= 0x30 && v[count-i-1] <= 0x39 {
			backward_found = true
			backward_val = int(v[count-i-1] - 0x30)
		}

		if forward_found && backward_found {
			break
		}
	}

	result := forward_val*10 + backward_val
	return result
}

func parse_string_2_bad(dut string) int {
	first_idx := len(dut)
	first_val := -1
	last_idx := -1
	last_val := -1

	values := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, v := range values {
		idx := strings.Index(dut, v)
		if idx > -1 {
			if idx < first_idx {
				first_val = i + 1
				first_idx = idx
			}
		}
		idx = strings.LastIndex(dut, v)
		if idx > last_idx {
			last_idx = idx
			last_val = i + 1
		}
	}

	for i := 0; i < first_idx; i++ {
		if dut[i] >= 0x30 && dut[i] <= 0x39 {
			first_val = int(dut[i] - 0x30)
			break
		}
	}

	for i := len(dut) - 1; i >= last_idx; i-- {
		if dut[i] >= 0x30 && dut[i] <= 0x39 {
			last_val = int(dut[i] - 0x30)
			break
		}
	}

	result := first_val*10 + last_val
	return result
}

func parse_str_2(v string, left *pattern_matcher.PatternNode, right *pattern_matcher.PatternNode) int {
	count := len(v)
	forward_val := -1
	backward_val := -1
	current_left := []*pattern_matcher.PatternNode{}
	current_right := []*pattern_matcher.PatternNode{}

	for i := 0; i < count; i++ {
		if forward_val < 0 {
			if v[i] > 0x30 && v[i] <= 0x39 {
				forward_val = int(v[i] - 0x30)
			} else {
				temp := []*pattern_matcher.PatternNode{}
				for _, n := range current_left {
					next := n.Next(v[i], left)
					if next.Value > 0 {
						forward_val = next.Value
					} else if next != left {
						temp = append(temp, next)
					}
				}
				current_left = temp
				next := left.Next(v[i], left)
				if next != left {
					current_left = append(current_left, next)
				}
			}
		}

		if backward_val < 0 {
			idx := count - i - 1
			if v[idx] > 0x30 && v[idx] <= 0x39 {
				backward_val = int(v[idx] - 0x30)
			} else {
				temp := []*pattern_matcher.PatternNode{}
				for _, n := range current_right {
					next := n.Next(v[idx], right)
					if next.Value > 0 {
						backward_val = next.Value
					} else if next != right {
						temp = append(temp, next)
					}
				}
				current_right = temp
				next := right.Next(v[idx], right)
				if next != right {
					current_right = append(current_right, next)
				}
			}
		}

		if forward_val >= 0 && backward_val >= 0 {
			break
		}
	}

	result := forward_val*10 + backward_val

	if result < 10 || result > 99 || result%10 == 0 {
		fmt.Println(v, ": ", result)
	}
	return result
}

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		result += parse_str(line)
	}
	return result
}

func part2(lines []string) int {
	forward_patterns := [9]pattern_matcher.Pattern{
		{Label: "one", Value: 1},
		{Label: "two", Value: 2},
		{Label: "three", Value: 3},
		{Label: "four", Value: 4},
		{Label: "five", Value: 5},
		{Label: "six", Value: 6},
		{Label: "seven", Value: 7},
		{Label: "eight", Value: 8},
		{Label: "nine", Value: 9},
	}
	fw_root := pattern_matcher.PatternNode{Label: ' ', Value: -1}
	pattern_matcher.CreateGraph(&fw_root, forward_patterns[:])

	bw_patterns := [9]pattern_matcher.Pattern{
		{Label: "eno", Value: 1},
		{Label: "owt", Value: 2},
		{Label: "eerht", Value: 3},
		{Label: "ruof", Value: 4},
		{Label: "evif", Value: 5},
		{Label: "xis", Value: 6},
		{Label: "neves", Value: 7},
		{Label: "thgie", Value: 8},
		{Label: "enin", Value: 9},
	}
	bw_root := pattern_matcher.PatternNode{Label: ' ', Value: -1}
	pattern_matcher.CreateGraph(&bw_root, bw_patterns[:])

	result := 0

	for _, line := range lines {
		result += parse_str_2(line, &fw_root, &bw_root)
	}

	return result
}

func part2slow(lines []string) int {
	result := 0
	for _, line := range lines {
		result += parse_string_2_bad(line)

	}
	return result
}

func main() {
	abs, _ := filepath.Abs("input")
	output, _ := file.ReadInput(abs)
	result := 0

	start := time.Now()
	result = part1(output)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 done with %d. It took %s\n", result, elapsed)

	start = time.Now()
	result = part2slow(output)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 slow done with %d. It took %s\n", result, elapsed)
}
