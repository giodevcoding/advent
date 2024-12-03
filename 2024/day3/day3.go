package day3

import (
	"fmt"
	"strconv"
	"unicode"
)

func CorruptedMemory(input []string) string {
	total := 0
    memory := ""
    for _, memoryPart := range input {
        memory += memoryPart
    }
	sourceStr := "mul("
	currentMatch := ""


	for i := 0; i < len(memory); i++ {
		currentChar := rune(memory[i])
		checkIndex := len(currentMatch)

		// checking for prefix
		if currentMatch != sourceStr {
			if currentChar == rune(sourceStr[checkIndex]) {
				currentMatch += string(sourceStr[checkIndex])
			} else {
				currentMatch = ""
			}
			//fmt.Println(currentMatch)

			//parsing number side
		} else {
			aStr, bStr := "", ""
			a, b := -1, -1
            numParsing:
			for j := i; a == -1 || b == -1; j++ {
				currentChar = rune(memory[j])
				switch {
				case unicode.IsDigit(currentChar):
					if a == -1 {
						aStr += string(currentChar)
					} else if b == -1 {
						bStr += string(currentChar)
					}
				case currentChar == ',':
					if a == -1 {
						num, err := strconv.Atoi(aStr)
						if err != nil {
							currentMatch = ""
							i = j
							break numParsing
						}
						a = num
					} else {
						currentMatch = ""
						i = j
						break numParsing
					}
				case currentChar == ')':
					if b == -1 {
						num, err := strconv.Atoi(bStr)
						if err != nil {
							currentMatch = ""
							i = j
							break numParsing
						}
						b = num
					} else {
						currentMatch = ""
						i = j
						break numParsing
					}
				default:
             //       fmt.Println("breaking at " + string(currentChar))
					currentMatch = ""
					i = j
					break numParsing

				}

				//Printing
				switch {
				case a == -1:
			//		fmt.Println(currentMatch + aStr)
				case b == -1:
			//		fmt.Println(currentMatch + aStr + "," + bStr)
				}
			}
			if a != -1 && b != -1 {
                fmt.Println(currentMatch + aStr + "," + bStr + ")")
				total += a * b
                fmt.Println(fmt.Sprintf("Adding %d to total, now equaling: %d", a*b, total))
				currentMatch = ""
			}
		}
	}

	return strconv.Itoa(total)
}
