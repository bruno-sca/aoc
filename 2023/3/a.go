package main

import (
	"strconv"
)

func SolveA(input string) string {
	lines := getLines(input)

	numbersStartCoord := make([]Coord, 0)

	for y, line := range lines {
		symbols := getSymbols(line, y)
		for _, s := range symbols {
			numbersStartCoord = append(numbersStartCoord, getAdjacenstNumbersStartCoord(lines, s.Coord)...)
		}
	}

	// Remove duplicates
	seen := make(map[Coord]bool)
	uniqueNumbersStartCoord := make([]Coord, 0)

	for _, c := range numbersStartCoord {
		if _, ok := seen[c]; !ok {
			uniqueNumbersStartCoord = append(uniqueNumbersStartCoord, c)
			seen[c] = true
		}
	}

	sum := 0

	for _, c := range uniqueNumbersStartCoord {
		sum += parseCordToInt(c, lines)
	}

	return strconv.Itoa(sum)
}
