package twelve

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type (
	ConditionRecord struct {
		damagedSprings string
		springGroups   []int
	}
)

func RunPartOne(input []string) {

}

func RunPartTwo(input []string) {

}

func PossibleArrangements(input []string) int {
	records := ConditionRecords(input)
	variations := GetPossibleVariations(records[0])
    _ = variations

    binarySlices := GetAllBinarySlices(3)

    for i := range binarySlices {
        fmt.Println(binarySlices[i])
    }

	return 0
}

func GetPossibleVariations(record ConditionRecord) (result []string) {
	/*
	   - ???.### variations
	   1 = ....###
	   2 = #...###
	   3 = .#..###
	   4 = ..#.###
	   5 = ##..###
	   6 = .##.###
	   7 = #.#.###
	   8 = ###.###
	*/
	return
}

func GetAllBinarySlices(length int) [][]int {
	result := [][]int{}
    return result
}

func ConditionRecords(input []string) []ConditionRecord {
	records := make([]ConditionRecord, len(input))
	for i, line := range input {
		records[i] = ParseConditionRecord(line)
	}
	return records
}

func ParseConditionRecord(input string) ConditionRecord {
	split := strings.Split(input, " ")
	damagedSprings := split[0]
	springGroups := []int{}

	for _, char := range split[1] {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			springGroups = append(springGroups, num)
		}
	}
	return ConditionRecord{
		damagedSprings,
		springGroups,
	}
}
