package main

import (
	"strconv"
)

func SolveB(input string) string {
	lines := getLines(input)

	deck := make([]int, len(lines))

	for i := 0; i < len(deck); i++ {
		deck[i] = 1
	}

	for i, line := range lines {
		cardNumbers, numbers, _ := getInfoFromLine(line)

		hits := checkNumbers(cardNumbers, numbers)

		for j := i + 1; j <= i+hits; j++ {
			deck[j] = deck[i] + deck[j]
		}
	}

	total := 0

	for _, v := range deck {
		total += v
	}

	return strconv.Itoa(total)
}
