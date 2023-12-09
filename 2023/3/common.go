package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type (
	Coord struct {
		X, Y int
	}
	Symbol struct {
		Coord
		Symbol rune
	}
)

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func getLines(input string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getSymbols(line string, y int) []Symbol {
	symbols := make([]Symbol, 0)
	for x, c := range line {
		if c == '.' || unicode.IsDigit(c) {
			continue
		}

		symbols = append(symbols, Symbol{Coord{x, y}, c})
	}

	return symbols
}

func getAsterisks(line string, y int) []Coord {
	asterisks := make([]Coord, 0)
	for x, c := range line {
		if c != '*' {
			continue
		}

		asterisks = append(asterisks, Coord{x, y})
	}

	return asterisks
}

func getAdjacenstNumbersStartCoord(lines []string, sc Coord) []Coord {
	coordsToCheck := []Coord{
		{X: sc.X - 1, Y: sc.Y - 1}, {X: sc.X, Y: sc.Y - 1}, {X: sc.X + 1, Y: sc.Y - 1},
		{X: sc.X - 1, Y: sc.Y}, {X: sc.X + 1, Y: sc.Y},
		{X: sc.X - 1, Y: sc.Y + 1}, {X: sc.X, Y: sc.Y + 1}, {X: sc.X + 1, Y: sc.Y + 1},
	}

	numberStartCoord := make([]Coord, 0)

	for _, c := range coordsToCheck {
		if c.X >= 0 && c.X < len(lines[0]) && c.Y >= 0 && c.Y < len(lines) && unicode.IsDigit(rune(lines[c.Y][c.X])) {
			x := c.X
			y := c.Y

			for x > 0 && unicode.IsDigit(rune(lines[y][x-1])) {
				x--
			}

			numberStartCoord = append(numberStartCoord, Coord{x, y})
		}
	}

	return numberStartCoord
}

func parseCordToInt(c Coord, lines []string) int {
	x := c.X

	numberString := ""

	for x < len(lines[c.Y]) && unicode.IsDigit(rune(lines[c.Y][x])) {
		numberString += string(lines[c.Y][x])
		x++
	}

	if len(numberString) == 0 {
		return 0
	}

	return ParseInt(numberString)
}
