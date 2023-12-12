package nine

import (
	"advent/helpers"
	"fmt"
)

func RunPartOne(input []string) {
	fmt.Println(OasisReportExtrapolation(input))
}

func RunPartTwo(input []string) {
	fmt.Println(OasisReportBackwardsExtrapolation(input))
}

func OasisReportExtrapolation(input []string) int {
	valueHistories := ValueHistories(input)

    result := 0

    for _, history := range valueHistories {
        result += GetNextHistoryValue(history)
    }

	return result
}

func OasisReportBackwardsExtrapolation(input []string) int {
	valueHistories := ValueHistories(input)

    result := 0

    for _, history := range valueHistories {
        result += GetPreviousHistoryValue(history)
    }

	return result
}

func GetNextHistoryValue(history []int) int {
	allDifferenceSequences := AllDifferenceSequences([][]int{history})
	result := 0

	for i := len(allDifferenceSequences) - 1; i >= 0; i-- {
		currentSequence := allDifferenceSequences[i]
		result += currentSequence[len(currentSequence)-1]
	}

	return result
}

func GetPreviousHistoryValue(history []int) int {
	allDifferenceSequences := AllDifferenceSequences([][]int{history})
	result := 0

	for i := len(allDifferenceSequences) - 1; i >= 0; i-- {
		currentSequence := allDifferenceSequences[i]
		result = currentSequence[0] - result
	}

	return result
}

func ValueHistories(input []string) [][]int {
	result := make([][]int, len(input))

	for i, historyStr := range input {
		result[i] = helpers.IntListStringToSlice(historyStr)
	}

	return result
}

func AllDifferenceSequences(sequences [][]int) [][]int {
	latestSequence := sequences[len(sequences)-1]

	differences := GetDifferenceSequence(latestSequence)
	result := append(sequences, differences)

	if len(differences) > 1 {
		if SequenceIsFlat(latestSequence) {
			return result
		} else {
			return AllDifferenceSequences(result)
		}
	}
	// else

	return result
}

func SequenceIsFlat(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] != sequence[i+1] {
			return false
		}
	}

	return true
}

func GetDifferenceSequence(sequence []int) []int {
	result := make([]int, len(sequence)-1)
	for i := 0; i < len(result); i++ {
		result[i] = sequence[i+1] - sequence[i]
	}

	return result
}
