package main

import (
	"strconv"
)

func SolveB(input string) string {
	lines := getLines(input)

	sum := 0

	for y, line := range lines {
		symbols := getAsterisks(line, y)
		for _, s := range symbols {
			numbersStartCoord := getAdjacenstNumbersStartCoord(lines, s)

			seen := make(map[Coord]bool)
			uniqueNumbersStartCoord := make([]Coord, 0)

			for _, c := range numbersStartCoord {
				if _, ok := seen[c]; !ok {
					uniqueNumbersStartCoord = append(uniqueNumbersStartCoord, c)
					seen[c] = true
				}
			}

			if len(uniqueNumbersStartCoord) != 2 {
				continue
			}

			n := 1

			for _, c := range uniqueNumbersStartCoord {
				n *= parseCordToInt(c, lines)
			}

			sum += n
		}
	}

	return strconv.Itoa(sum)
}
