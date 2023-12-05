package helpers

import (
	"strconv"
	"strings"
)

func PrefixExistsInList(sub string, list []string) bool {
	for _, str := range list {
		if strings.HasPrefix(str, sub) {
			return true
		}
	}
	return false
}

func IntListStringToSlice(numList string) []int {
	textNumbers := strings.Split(strings.TrimSpace(numList), " ")
	actualNumbers := []int{}

	for _, textNum := range textNumbers {
        if (len(textNum) == 0) {
            continue
        }
		actualNum, err := strconv.Atoi(textNum)
		if err == nil {
			actualNumbers = append(actualNumbers, actualNum)
		}
	}

	return actualNumbers
}

