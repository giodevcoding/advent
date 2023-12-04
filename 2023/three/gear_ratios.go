package three

import (
	"fmt"
	"strconv"
	"unicode"
)

type (
	PartNumber struct {
		value  int
		index  int
		length int
	}
	PartNumberSlice []PartNumber
)

func RunPartOne(input []string) {
	fmt.Println(EngineSchematicSum(input))
}

func RunPartTwo(input []string) {
    fmt.Println(GearRatioSum(input))
}

func EngineSchematicSum(input []string) (sum int) {
	partNumbers, symbolIndexes := GetPartNumbersAndSymbolIndexes(input)
	for y, line := range partNumbers {
		for _, partNumber := range line {
			if IsPartNumberValid(partNumber, y, symbolIndexes) {
				sum += partNumber.value
			}
		}
	}

	return sum
}

func GearRatioSum(input []string) (sum int) {
	partNumbers, gearIndexes := GetPartNumbersAndGearSymbolIndexes(input)
	_ = partNumbers

    for y, gearLine := range gearIndexes {
        for _, gearIndex := range gearLine {
            valid, ratio := GearValid(gearIndex, y, partNumbers)
            if valid {
                sum += ratio
            }
        }
    }
	return sum
}

func GearValid(gearIndex int, gearY int, allPartNumbers [][]PartNumber) (valid bool, ratio int) {
	partsTouching := []PartNumber{}

	if gearY > 0 {
		for _, partNumber := range allPartNumbers[gearY-1] {
			if SymbolInXAxisOfPartNumber([]int{gearIndex}, partNumber) {
				partsTouching = append(partsTouching, partNumber)
			}
		}
	}

	for _, partNumber := range allPartNumbers[gearY] {
		if SymbolInXAxisOfPartNumber([]int{gearIndex}, partNumber) {
			partsTouching = append(partsTouching, partNumber)
		}
	}

	if gearY < len(allPartNumbers)-1 {
		for _, partNumber := range allPartNumbers[gearY+1] {
			if SymbolInXAxisOfPartNumber([]int{gearIndex}, partNumber) {
				partsTouching = append(partsTouching, partNumber)
			}
		}
	}

	if len(partsTouching) == 2 {
		valid = true
        ratio = partsTouching[0].value*partsTouching[1].value
	}

	return valid, ratio
}

func IsPartNumberValid(partNumber PartNumber, partNumberY int, symbolIndexes [][]int) bool {
	// Line above
	if partNumberY > 0 {
		if SymbolInXAxisOfPartNumber(symbolIndexes[partNumberY-1], partNumber) {
			return true
		}
	}

	if SymbolInXAxisOfPartNumber(symbolIndexes[partNumberY], partNumber) {
		return true
	}

	// Line below
	if partNumberY < len(symbolIndexes)-1 {
		if SymbolInXAxisOfPartNumber(symbolIndexes[partNumberY+1], partNumber) {
			return true
		}
	}
	return false
}

func SymbolInXAxisOfPartNumber(symbolIndexes []int, partNumber PartNumber) bool {
	for _, symbolIndex := range symbolIndexes {
		if symbolIndex >= partNumber.index-1 && symbolIndex <= partNumber.index+partNumber.length {
			return true
		}
	}

	return false
}

func GetPartNumbersAndSymbolIndexes(input []string) (partNumbers [][]PartNumber, symbols [][]int) {
	for _, line := range input {
		partNumbers = append(partNumbers, ExtractPartNumbersFromLine(line))
		symbols = append(symbols, ExtractSymbolIndexesFromLine(line))
	}

	return partNumbers, symbols
}

func GetPartNumbersAndGearSymbolIndexes(input []string) (partNumbers [][]PartNumber, symbolIndexes [][]int) {
	for _, line := range input {
		partNumbers = append(partNumbers, ExtractPartNumbersFromLine(line))
		symbolIndexes = append(symbolIndexes, ExtractGearSymbolIndexesFromLine(line))
	}

	return partNumbers, symbolIndexes
}

func ExtractPartNumbersFromLine(line string) PartNumberSlice {
	partNumbers := []PartNumber{}
	currentNumber := ""
	currentIndex := 0

	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		if unicode.IsDigit(char) {
			currentIndex = i
			for j := i; j < len(line); j++ {
				subChar := rune(line[j])

				if !unicode.IsDigit(subChar) || j >= len(line)-1 {
					if j >= len(line)-1 && unicode.IsDigit(subChar) {
						currentNumber += string(subChar)
					}
					value, _ := strconv.Atoi(currentNumber)
					partNumber := PartNumber{
						value:  value,
						index:  currentIndex,
						length: len(currentNumber),
					}

					partNumbers = append(partNumbers, partNumber)
					currentNumber = ""

					i = j
					break
				}

				currentNumber += string(subChar)
			}
		}
	}

	return partNumbers
}

func ExtractSymbolIndexesFromLine(line string) []int {
	indexes := []int{}
	for i, char := range line {
		if IsSymbol(char) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func ExtractGearSymbolIndexesFromLine(line string) []int {
	indexes := []int{}
	for i, char := range line {
		if char == '*' {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func IsSymbol(input rune) bool {
	return input != '.' && !unicode.IsDigit(input)
}

func (slice PartNumberSlice) Equals(otherSlice PartNumberSlice) bool {
	if len(slice) != len(otherSlice) {
		return false
	}

	for i := range slice {
		if slice[i] != otherSlice[i] {
			return false
		}
	}
	return true
}
