package day1

import (
	"aoc2024/utils"
	"math"
	"strconv"
)

func ReconcileLists(input []string) string {
	var listLen = len(input)
	var list1, list2 = make([]int, listLen), make([]int, listLen)

	for i, line := range input {
		nums := utils.StringToNumList(line)
		list1[i] = nums[0]
		list2[i] = nums[1]
	}

	list1 = utils.MergeSortInt(list1)
	list2 = utils.MergeSortInt(list2)

	differencesSum := 0

	for i := range list1 {
		differencesSum += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return strconv.Itoa(differencesSum)
}

func SimilarityScores(input []string) string {
	var listLen = len(input)
	var list1, list2 = make([]int, listLen), make([]int, listLen)

	for i, line := range input {
		nums := utils.StringToNumList(line)
		list1[i] = nums[0]
		list2[i] = nums[1]
	}

    similarityScore := 0

    for _, left := range list1 {
        count := 0
        for _, right := range list2 {
            if left == right {
                count++
            }
        }
        similarityScore += left*count
    }

    return strconv.Itoa(similarityScore)
}
