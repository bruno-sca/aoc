package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type AlmanacData struct {
	seeds                 []int
	seetToSoil            []AlmanacMap
	soilToFertilizer      []AlmanacMap
	fertilizerToWater     []AlmanacMap
	waterToLight          []AlmanacMap
	lightToTemperature    []AlmanacMap
	temperatureToHumidity []AlmanacMap
	humidityToLocation    []AlmanacMap
}

type AlmanacMap struct {
	destStart   int
	targetStart int
	mapRange    int
}

func getAlmanacData(input string) AlmanacData {
	maps := strings.Split(input, "\n\n")

	data := AlmanacData{
		seeds:                 getSeeds(maps[0]),
		seetToSoil:            getMap(maps[1]),
		soilToFertilizer:      getMap(maps[2]),
		fertilizerToWater:     getMap(maps[3]),
		waterToLight:          getMap(maps[4]),
		lightToTemperature:    getMap(maps[5]),
		temperatureToHumidity: getMap(maps[6]),
		humidityToLocation:    getMap(maps[7]),
	}

	return data
}

func getMap(input string) []AlmanacMap {

	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lines = lines[1:]

	m := make([]AlmanacMap, 0)

	for _, line := range lines {
		split := strings.Split(line, " ")

		destStart := ParseInt(split[0])
		targetStart := ParseInt(split[1])
		mapRange := ParseInt(split[2])

		if mapRange == 0 {
			continue
		}

		m = append(m, AlmanacMap{
			destStart:   destStart,
			targetStart: targetStart,
			mapRange:    mapRange,
		})
	}

	return m
}

func getSeeds(input string) []int {
	var seeds []int

	numbers := strings.Split(strings.Trim(strings.Split(input, ":")[1], " "), " ")

	for _, number := range numbers {
		seeds = append(seeds, ParseInt(number))
	}

	return seeds
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error parseint: %e", err)
	}
	return i
}

func getSeedsLocations(
	seeds []int, seetToSoil []AlmanacMap,
	soilToFertilizer []AlmanacMap,
	fertilizerToWater []AlmanacMap,
	waterToLight []AlmanacMap,
	lightToTemperature []AlmanacMap,
	temperatureToHumidity []AlmanacMap,
	humidityToLocation []AlmanacMap,
) []int {
	var seedsLocations []int

	for _, seed := range seeds {

		seedSoil := getTargetFromMap(seetToSoil, seed)
		seedFertilizer := getTargetFromMap(soilToFertilizer, seedSoil)
		seedWater := getTargetFromMap(fertilizerToWater, seedFertilizer)
		seedLight := getTargetFromMap(waterToLight, seedWater)
		seedTemperature := getTargetFromMap(lightToTemperature, seedLight)
		seedHumidity := getTargetFromMap(temperatureToHumidity, seedTemperature)
		seedLocation := getTargetFromMap(humidityToLocation, seedHumidity)

		// fmt.Printf("Seed %d, soil: %d, fertilizer: %d, water: %d, light: %d, temperature: %d, humidity: %d, location,: %d\n", seed, seedSoil, seedFertilizer, seedWater, seedLight, seedTemperature, seedHumidity, seedLocation)

		seedsLocations = append(seedsLocations, seedLocation)
	}

	return seedsLocations
}

func getTargetFromMap(m []AlmanacMap, target int) int {
	for _, mapItem := range m {
		if target >= mapItem.targetStart && target < mapItem.targetStart+mapItem.mapRange {
			return mapItem.destStart + (target - mapItem.targetStart)
		}
	}

	return target
}

func getSeedRange(start int, r int) []int {
	var seedRange []int

	for i := 0; i < r; i++ {
		seedRange = append(seedRange, start+i)
	}

	return seedRange
}
