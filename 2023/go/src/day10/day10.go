package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
)

type Position struct {
	x int
	y int
}

type TrackStatus struct {
	current    Position
	currentDir Direction
	start      Position
	err        error
	Label      string
}
type Direction int

const (
	Down  Direction = iota
	Up    Direction = iota
	Left  Direction = iota
	Right Direction = iota
)

func findStart(lines []string) (Position, error) {
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				return Position{x, y}, nil
			}
		}
	}
	return Position{-1, -1}, errors.New("could not find start")
}

func getExit(c byte, from Direction) (Direction, error) {
	switch from {
	case Down:
		if c == '|' {
			return Up, nil
		} else if c == '7' {
			return Left, nil
		} else if c == 'F' {
			return Right, nil
		}
	case Up:
		if c == '|' {
			return Down, nil
		} else if c == 'L' {
			return Right, nil
		} else if c == 'J' {
			return Left, nil
		}
	case Left:
		if c == '-' {
			return Right, nil
		} else if c == 'J' {
			return Up, nil
		} else if c == '7' {
			return Down, nil
		}
	case Right:
		if c == '-' {
			return Left, nil
		} else if c == 'L' {
			return Up, nil
		} else if c == 'F' {
			return Down, nil
		}
	}
	return from, errors.New("could not exit")
}

func (d Direction) opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	return d
}
func navigate(lines []string, from Direction, pos Position) (Position, Direction, error) {
	next, err := getExit(lines[pos.y][pos.x], from)
	if err != nil {
		return pos, from, err
	}

	chars := []rune(lines[pos.y])
	chars[pos.y] = '*'

	// Use the exit direction to determine the next grid location
	switch next {
	case Up:
		pos.y--
	case Down:
		pos.y++
	case Left:
		pos.x--
	case Right:
		pos.x++
	}

	// Validate the grid location
	if pos.x < 0 || pos.x >= len(lines[0]) {
		return pos, from, errors.New("out of range")
	} else if pos.y < 0 || pos.y >= len(lines) {
		return pos, from, errors.New("out of range")
	}

	// We are entering the position from the opposite direction we traveled.  i.e. if we went Up we entered that
	// sqaure from Down
	return pos, next.opposite(), nil
}

func drawMap(lines []string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for _, line := range lines {
		fmt.Println(line)
	}
}
func part1(lines []string) int {
	startPosition, err := findStart(lines)
	if err != nil {
		fmt.Errorf(err.Error())
		return 0
	}

	up := Position{startPosition.x, startPosition.y - 1}
	down := Position{startPosition.x, startPosition.y + 1}
	left := Position{startPosition.x - 1, startPosition.y}
	right := Position{startPosition.x + 1, startPosition.y}

	upStatus := TrackStatus{current: up, currentDir: Down, start: up, err: nil, Label: "Up"}
	downStatus := TrackStatus{current: down, currentDir: Up, start: down, err: nil, Label: "Down"}
	leftStatus := TrackStatus{current: left, currentDir: Right, start: left, err: nil, Label: "Left"}
	rightStatus := TrackStatus{current: right, currentDir: Left, start: right, err: nil, Label: "Right"}

	validPaths := []TrackStatus{
		leftStatus, upStatus, downStatus, rightStatus,
	}
	stepCount := 1
	for {
		var nextValid []TrackStatus
		for _, ts := range validPaths {
			var next Position
			next, ts.currentDir, ts.err = navigate(lines, ts.currentDir, ts.current)
			if ts.err == nil {
				chars := []byte(lines[ts.current.y])
				chars[ts.current.x] = '*'
				lines[ts.current.y] = string(chars)
			}
			ts.current = next
			if ts.err == nil {
				nextValid = append(nextValid, ts)
			} else {
				fmt.Println("removing ", ts.Label)
			}
		}
		stepCount++

		// See if any of the valid paths are equal
		for i := 0; i < len(nextValid); i++ {

			for j := i + 1; j < len(nextValid); j++ {
				if nextValid[i].current == nextValid[j].current {
					drawMap(lines)
					return stepCount
				}
			}
		}
		validPaths = nextValid
		//drawMap(lines)
		//fmt.Println()
		//fmt.Println("Steps Done ", stepCount)
	}
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
