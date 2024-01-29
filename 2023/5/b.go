package main

import "strconv"

func SolveB(input string) string {
	almanacData := getAlmanacData(input)

	lowest := int(^uint(0) >> 1)

	for i := 0; i < len(almanacData.seeds); i = i + 2 {

		seedRange := getSeedRange(almanacData.seeds[i], almanacData.seeds[i+1])
		locations := getSeedsLocations(seedRange,
			almanacData.seetToSoil,
			almanacData.soilToFertilizer,
			almanacData.fertilizerToWater,
			almanacData.waterToLight,
			almanacData.lightToTemperature,
			almanacData.temperatureToHumidity,
			almanacData.humidityToLocation,
		)

		for _, location := range locations {
			if location < lowest {
				lowest = location
			}
		}

		almanacData.seeds = make([]int, 0)
		seedRange = nil
	}

	return strconv.Itoa(lowest)
}
