package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"advent2023/pkg/file"
	"advent2023/pkg/math"
)

type NetworkNode struct {
	Label int //
	Human string
	Right *NetworkNode
	Left  *NetworkNode
}

type NetworkMap map[int][2]int

func stringToInt(val string) int {
	return int(val[0])<<16 | int(val[1])<<8 | int(val[2])
}

func parseNodes(lines []string, startIndex int) *NetworkNode {
	var root *NetworkNode
	rootId := stringToInt("AAA")

	createdNodes := make(map[int]*NetworkNode)

	for i := startIndex; i < len(lines); i++ {
		current := stringToInt(lines[i][0:3])
		left := stringToInt(lines[i][7:10])
		right := stringToInt(lines[i][12:])

		currentNode := createdNodes[current]
		if currentNode == nil {
			newNode := NetworkNode{Label: current}
			currentNode = &newNode
			createdNodes[current] = currentNode
		}

		leftNode := createdNodes[left]
		if leftNode == nil {
			newNode := NetworkNode{Label: left}
			leftNode = &newNode
			createdNodes[left] = leftNode
		}
		currentNode.Left = leftNode

		rightNode := createdNodes[right]
		if rightNode == nil {
			newNode := NetworkNode{Label: right}
			rightNode = &newNode
			createdNodes[right] = rightNode
		}
		currentNode.Right = rightNode

		if current == rootId {
			root = currentNode
		}
	}
	return root
}

func part1(lines []string) int {
	inst := lines[0]

	endLabel := stringToInt("ZZZ")
	currentNode := parseNodes(lines, 2)
	instIndex := 0
	stepCount := 0
	for currentNode.Label != endLabel {
		if inst[instIndex] == 'L' {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
		stepCount++
		instIndex++
		if instIndex >= len(inst) {
			instIndex = 0
		}
	}
	return stepCount
}

func linesToMap(lines []string, startIndex int) (NetworkMap, []int) {
	result := make(NetworkMap)
	var starts []int
	for i := startIndex; i < len(lines); i++ {
		current := stringToInt(lines[i][0:3])
		left := stringToInt(lines[i][7:10])
		right := stringToInt(lines[i][12:])

		result[current] = [2]int{left, right}

		if current&0xFF == int('A') {
			starts = append(starts, current)
		}
	}

	return result, starts
}

func parseNodes2(lines []string, startIndex int) []*NetworkNode {
	var roots []*NetworkNode
	rootId := int('A')

	createdNodes := make(map[int]*NetworkNode)

	for i := startIndex; i < len(lines); i++ {
		current := stringToInt(lines[i][0:3])
		left := stringToInt(lines[i][7:10])
		right := stringToInt(lines[i][12:])

		currentNode := createdNodes[current]
		if currentNode == nil {
			newNode := NetworkNode{Label: current}
			currentNode = &newNode
			createdNodes[current] = currentNode
			currentNode.Human = lines[i][0:3]
		}

		leftNode := createdNodes[left]
		if leftNode == nil {
			newNode := NetworkNode{Label: left}
			leftNode = &newNode
			createdNodes[left] = leftNode
			leftNode.Human = lines[i][7:10]
		}
		currentNode.Left = leftNode

		rightNode := createdNodes[right]
		if rightNode == nil {
			newNode := NetworkNode{Label: right}
			rightNode = &newNode
			createdNodes[right] = rightNode
			rightNode.Human = lines[i][12:15]
		}
		currentNode.Right = rightNode

		if current&0xFF == rootId {
			roots = append(roots, currentNode)
		}
	}
	return roots
}
func part2(lines []string) int {

	var inst []bool
	for _, c := range lines[0] {
		inst = append(inst, c == 'L')
	}
	var wg sync.WaitGroup
	graph, starts := linesToMap(lines, 2)
	freqSteps := make([]int, len(starts))

	/*
		    Brute force here (just following the graph until all end nodes had been reached.  It can go through about
		    4 million steps/second which ends up being ~41 days of computation.

			I only got this by looking at other people's work.  It seems that if you make the assumption that everything is
		    a loop that only hits one endpoint and that the loop is a multiple of the first time you hit that endpoint
			you can then calculate the LCM.  In general, I _think_ you would need to do something like this:
		        1. Keep a map of step indexes and which nodes were landed on at that step index.
		        2. Once you've landed on the same node on the same index, you have a loop and can calculate at what steps
		           you land on the end node for each loop (there may be some before the loop starts as well)

		    Once you have done that for all the starting indexes, you can find which step all paths will have an end.
	*/
	for i := 0; i < len(starts); i++ {
		wg.Add(1)
		go func(start int, freqStep *int) {
			defer wg.Done()
			idx := 0
			position := start

			*freqStep = 0
			for {
				if inst[idx] {
					position = graph[position][0]
				} else {
					position = graph[position][1]
				}
				*freqStep++
				if position&0xFF == int('Z') {
					fmt.Println("Final position found for start ", start, ": ", position, " after ", *freqStep)
					break
				}
				idx++
				if idx == len(inst) {
					idx = 0
				}
			}
		}(starts[i], &freqSteps[i])
	}
	wg.Wait()
	return math.LCM(freqSteps[0], freqSteps[1:])
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
