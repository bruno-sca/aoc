package main

import (
	"bufio"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func checkIfContainsNumber(numbers []int, target int) bool {
	return slices.Contains(numbers, target)
}

func checkNumbers(numbers []int, targets []int) int {
	count := 0

	for _, target := range targets {
		if checkIfContainsNumber(numbers, target) {
			count++
		}
	}

	return count
}

func getLines(input string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getInfoFromLine(line string) ([]int, []int, int) {
	info := strings.Split(line, ":")

	cardNumber := regexp.MustCompile(`\d+`).FindAllString(info[0], 1)[0]

	numbers := strings.Split(info[1], "|")

	cardNumbers := parseIntSlice(deleteEmpty(strings.Split(strings.Trim(numbers[0], " "), " ")))

	targetNumbers := parseIntSlice(deleteEmpty(strings.Split(strings.Trim(numbers[1], " "), " ")))

	return cardNumbers, targetNumbers, ParseInt(cardNumber)
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func parseIntSlice(s []string) []int {
	var result []int
	for _, i := range s {
		result = append(result, ParseInt(i))
	}
	return result
}
