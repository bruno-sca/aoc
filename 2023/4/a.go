package main

import (
	"math"
	"strconv"
)

func SolveA(input string) string {
	lines := getLines(input)

	totalPoints := 0

	for _, line := range lines {
		cardNumbers, numbers, _ := getInfoFromLine(line)

		totalPoints += int(math.Pow(float64(2), float64(checkNumbers(cardNumbers, numbers)-1)))
	}

	return strconv.Itoa(totalPoints)
}
