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
    
}

func EngineSchematicSum(input []string) (sum int) {
    partNumbers, symbols := GetPartNumbersAndSymbols(input) 
    for y, line := range partNumbers {
        for _, partNumber := range line {
           if IsPartNumberValid (partNumber, y, symbols) {
               sum += partNumber.value
           }
        }
    }

    return sum
}

func IsPartNumberValid(partNumber PartNumber, partNumberY int, symbols [][]int) bool {
    // TODO: Check if symbols contains an index that would make the PartNumber valid
    return false
}

func GetPartNumbersAndSymbols(input []string) (partNumbers [][]PartNumber, symbols [][]int) {
    for _, line := range input {
        partNumbers = append(partNumbers, ExtractPartNumbersFromLine(line))
        symbols = append(symbols, ExtractSymbolIndexesFromLine(line))
    } 

    return partNumbers, symbols
}

func ExtractPartNumbersFromLine(line string) PartNumberSlice {
	partNumbers := []PartNumber{}
	currentNumber := ""
	currentIndex := 0

    for i := 0; i < len(line); i++ {
        char := rune(line[i])
		if unicode.IsDigit(char) {
			currentIndex = i
            fmt.Println(currentIndex)
			for j := i; j < len(line); j++ {
				subChar := rune(line[j])

				if !unicode.IsDigit(subChar) {
					value, _ := strconv.Atoi(currentNumber)
					partNumber := PartNumber{
						value:  value,
						index:  currentIndex,
						length: len(currentNumber),
					}

					partNumbers = append(partNumbers, partNumber)

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
