package five

import (
	"advent/helpers"
	"fmt"
	"math"
	"strings"
	"unicode"
)

type (
	AlmanacMap struct {
		name   string
		ranges []Range
	}

	Range struct {
		destinationStart int
		sourceStart      int
		length           int
	}
)

func RunPartOne(input []string) {
	fmt.Println(LowestLocationSingleSeeds(input))
}

func RunPartTwo(input []string) {
	fmt.Println(LowestLocationSeedRanges(input))
}

func LowestLocationSingleSeeds(input []string) int {
	seeds := ParseSeeds(input)
	maps := ParseMaps(input)
	lowest := math.MaxInt

	for _, seed := range seeds {
		location := GetLocationForSeed(seed, maps)
		if location < lowest {
			lowest = location
		}
	}

	return lowest
}

func LowestLocationSeedRanges(input []string) int {
	seedRanges := ParseSeeds(input)
	maps := ParseMaps(input)
	lowest := math.MaxInt

	for i := 0; i < len(seedRanges); i += 2 {
		rangeStart, rangeLength := seedRanges[i], seedRanges[i+1]
		rangeEnd := rangeStart + rangeLength

		for seed := rangeStart; seed < rangeEnd; seed++ {
			location := GetLocationForSeed(seed, maps)
			if location < lowest {
				lowest = location
			}
		}
	}

    return lowest
}

func GetLocationForSeed(seed int, maps []AlmanacMap) (result int) {
	result = seed
	for _, almanacMap := range maps {
		result = GetMappedOutput(result, almanacMap)
	}
	return result
}

func GetMappedOutput(input int, almanacMap AlmanacMap) int {
	for _, mapRange := range almanacMap.ranges {
		if input >= mapRange.sourceStart && input < mapRange.sourceStart+mapRange.length {
			return input + (mapRange.destinationStart - mapRange.sourceStart)
		}
	}

	return input
}

func ParseSeeds(input []string) []int {
	return helpers.IntListStringToSlice(strings.Split(input[0], ":")[1])
}

func ParseMaps(input []string) (result []AlmanacMap) {
	mapsData := input
	currentMap := AlmanacMap{}

	for _, almanacMapLine := range mapsData {
		if len(strings.TrimSpace(almanacMapLine)) <= 0 {
			continue
		}
		if !unicode.IsDigit(rune(almanacMapLine[0])) && len(currentMap.ranges) == 0 {
			currentMap.name = almanacMapLine
		} else if !unicode.IsDigit(rune(almanacMapLine[0])) && len(currentMap.ranges) > 0 {
			result = append(result, currentMap)
			currentMap = AlmanacMap{}
		} else if unicode.IsDigit(rune(almanacMapLine[0])) {
			ints := helpers.IntListStringToSlice(almanacMapLine)
			currentMap.ranges = append(currentMap.ranges, Range{
				destinationStart: ints[0],
				sourceStart:      ints[1],
				length:           ints[2],
			})
		}
	}

	result = append(result, currentMap)

	return result
}
