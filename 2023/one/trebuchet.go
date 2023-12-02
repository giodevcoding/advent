package one

import (
	"advent/helpers"
	"fmt"
	"path/filepath"
	"slices"
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
	var absPath, _ = filepath.Abs("./one/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ProperCalibrationValues(lines))
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
    foundFirst := false

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
    sum := 0

	for _, str := range input {
        calibrationValue := GetProperCalibrationValue(str)
		sum += calibrationValue
	}

	return sum
}

func GetProperCalibrationValue(str string) int {
    digits := NumbersInString(str)
    firstDigit := digits[0]
    lastDigit := digits[len(digits)-1]

	return (firstDigit*10)+lastDigit
}

func NumbersInString(str string) []int {
    numberWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
    wordStartIndex := 0
    currentWord := ""
    result := []int{}
    stringLength := len(str)

	for i := 0; i < stringLength; i++ {
		if len(currentWord) == 0 {
			wordStartIndex = i
		}

		currentWord += string(str[i])

        // Check if this is an actual digit rather than a word
		if len(currentWord) == 1 {
			var digit, err = strconv.Atoi(currentWord)
			if err == nil {
				result = append(result, digit)
				currentWord = ""
				continue
			}
		}

		if helpers.PrefixExistsInList(currentWord, numberWords) {
			if slices.Contains(numberWords, currentWord) {
				result = append(result, GetDigit(currentWord))
				currentWord = ""
				i = wordStartIndex
			}
		} else {
            currentWord = ""
			i = wordStartIndex
		}
	}

    return result
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
		return result
	}

	return 0
}
