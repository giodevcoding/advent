package day7

import (
	"aoc2024/utils"
	"strconv"
	"strings"
)

func BridgeRepairOperators(input []string) string {
	total := 0
	results, values := getCalibrations(input)
	for i, expectedResult := range results {
		found := false
		if isValidCalibration(expectedResult, values[i], 0, 0, &found) {
			total += expectedResult
		}
	}
	return strconv.Itoa(total)
}

func isValidCalibration(expectedResult int, pieces []int, total int, index int, alreadyFound *bool) bool {
	if *alreadyFound {
		return *alreadyFound
	}

	if index == len(pieces) {
		if total == expectedResult {
			*alreadyFound = true
			return true
		}
		return false
	}

	return isValidCalibration(expectedResult, pieces, total+pieces[index], index+1, alreadyFound) ||
		isValidCalibration(expectedResult, pieces, total*pieces[index], index+1, alreadyFound) ||
		isValidCalibration(expectedResult, pieces, concatInts(total, pieces[index]), index+1, alreadyFound)
}

func getCalibrations(input []string) (results []int, values [][]int) {
	for _, line := range input {
		split := strings.Split(line, ":")
		result, _ := strconv.Atoi(split[0])
		results = append(results, result)

		testValues := utils.Map(strings.Split(split[1], " "), func(val string) int {
			num, _ := strconv.Atoi(val)
			return num
		})

		values = append(values, testValues)
	}

	return
}

func concatInts(a int, b int) (concatInt int) {
	concatStr := strconv.Itoa(a) + strconv.Itoa(b)
	concatInt, _ = strconv.Atoi(concatStr)
	return
}
