package main

import (
	"fmt"
	"path/filepath"
	"time"

	"advent2023/pkg/file"
	"advent2023/pkg/numerics"
)

type GameDraw struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id       int
	Draws    []GameDraw
	MaxRed   int
	MaxGreen int
	MaxBlue  int
	MinRed   int
	MinGreen int
	MinBlue  int
}

func stringToGame(line string) Game {
	var rval Game
	var count int
	var index int

	rval.Id, index = numerics.GetNumeric(line, 5)
	var draw GameDraw
	for {

		// Read in number of color
		count, index = numerics.GetNumeric(line, index)
		index++
		// get the color
		switch line[index] {
		case 'r':
			draw.Red = count
			if count > rval.MaxRed {
				rval.MaxRed = count
			}
			if rval.MinRed == 0 || rval.MinRed < count {
				rval.MinRed = count
			}
			index += 3
		case 'g':
			draw.Green = count
			if count > rval.MaxGreen {
				rval.MaxGreen = count
			}
			if rval.MinGreen == 0 || rval.MinGreen < count {
				rval.MinGreen = count
			}
			index += 5
		case 'b':
			draw.Blue = count
			if count > rval.MaxBlue {
				rval.MaxBlue = count
			}
			if rval.MinBlue == 0 || rval.MinBlue < count {
				rval.MinBlue = count
			}
			index += 4
		}

		// Is this the end of the line
		if index >= len(line) {
			rval.Draws = append(rval.Draws, draw)
			return rval
		}

		// Is this another round
		if line[index] == ';' {
			rval.Draws = append(rval.Draws, draw)
			draw.Blue = 0
			draw.Red = 0
			draw.Green = 0
		}
		index += 2
	}
}

func part1(lines []string) int {

	rval := 0
	for _, line := range lines {
		game := stringToGame(line)
		if game.MaxBlue <= 14 && game.MaxGreen <= 13 && game.MaxRed <= 12 {
			rval += game.Id
		}
	}
	return rval
}

func part2(lines []string) int {

	rval := 0
	for _, line := range lines {
		game := stringToGame(line)
		power := game.MinBlue * game.MinGreen * game.MinRed
		rval += power
	}
	return rval
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
