package day7

import (
	"aoc2024/utils"
	"strconv"
	"strings"
)

type Operator string

const (
	add Operator = "+"
	mul Operator = "*"
	cat Operator = "||"
)

func BridgeRepairOperators(input []string) string {
	total := 0
	results, values := getCalibrations(input)
	for i, expectedResult := range results {
		found := false
		if isValidCalibration(expectedResult, values[i], 0, 0, add, &found) {
			total += expectedResult
		}
	}
	return strconv.Itoa(total)
}

func isValidCalibration(expectedResult int, pieces []int, total int, index int, op Operator, alreadyFound *bool) bool {
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

	switch op {
	case add:
		total += pieces[index]
	case mul:
		total *= pieces[index]
	case cat:
		total = concatInts(total, pieces[index])
	}

	withAdd := isValidCalibration(expectedResult, pieces, total, index+1, add, alreadyFound)
	withMul := isValidCalibration(expectedResult, pieces, total, index+1, mul, alreadyFound)
	withCat := isValidCalibration(expectedResult, pieces, total, index+1, cat, alreadyFound)
	return withAdd || withMul || withCat
}

func getCalibrations(input []string) (results []int, values [][]int) {
	for _, line := range input {
		split := strings.Split(line, ":")
		result, _ := strconv.Atoi(split[0])
		results = append(results, result)

		testValuesStr := utils.Filter(strings.Split(split[1], " "), func(str string) bool {
			return len(str) > 0
		})
		testValues := utils.Map(testValuesStr, func(val string) int {
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
