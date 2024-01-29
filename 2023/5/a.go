package main

import "strconv"

func SolveA(input string) string {
	almanacData := getAlmanacData(input)

	locations := getSeedsLocations(almanacData.seeds,
		almanacData.seetToSoil,
		almanacData.soilToFertilizer,
		almanacData.fertilizerToWater,
		almanacData.waterToLight,
		almanacData.lightToTemperature,
		almanacData.temperatureToHumidity,
		almanacData.humidityToLocation)

	lowest := int(^uint(0) >> 1)

	for _, location := range locations {
		if location < lowest {
			lowest = location
		}
	}

	return strconv.Itoa(lowest)
}
