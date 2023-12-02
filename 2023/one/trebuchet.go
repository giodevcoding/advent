package one

import (
	"advent/helpers"
	"fmt"
	"path/filepath"
	"strconv"
	"unicode"
)

func RunPartOne() {
	var absPath, _ = filepath.Abs("./one/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(CalibrationValues(lines))
}

func RunPartTwo() {

}

func CalibrationValues(input []string) int {
	var sum = 0

	for _, str := range input {
		var calibrationValue = GetCalibrationValue(str)
		sum += calibrationValue
	}

	return sum
}

func GetCalibrationValue(str string) int {
	var first, last rune
	var foundFirst = false

	for _, char := range str {
		if unicode.IsDigit(char) {
			if !foundFirst {
				first = char
				foundFirst = true
			}
			last = char
		}
	}

	var calibrationValue, _ = strconv.Atoi(string(first) + string(last))

	return calibrationValue
}

func ProperCalibrationValues(input []string) int {
	var sum = 0

	for _, str := range input {
		var calibrationValue = GetProperCalibrationValue(str)
		sum += calibrationValue
	}

	return sum
}

func GetProperCalibrationValue(str string) int {

	return 0
}

func GetDigit(str string) int {
	var values = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var result, err = strconv.Atoi(str)

	if err != nil {
		var result, ok = values[str]
		if ok {
			return result
		}
	} else {
        fmt.Println(result)
		return result
	}

	return 0
}
