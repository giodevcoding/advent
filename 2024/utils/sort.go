package utils

import "math"

func MergeSortInt(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := int(math.Floor(float64(len(arr)) / 2))

    leftSlice := MergeSortInt(arr[0:mid])
    rightSlice := MergeSortInt(arr[(mid):len(arr)])
    return mergeInt(leftSlice, rightSlice)
}

func mergeInt(leftSlice []int, rightSlice []int) []int {
	leftSize, rightSize := len(leftSlice), len(rightSlice)
	merged := make([]int, 0, leftSize+rightSize)

	leftIndex, rightIndex := 0, 0

	for leftIndex < leftSize && rightIndex < rightSize {
		leftItem, rightItem := leftSlice[leftIndex], rightSlice[rightIndex]
		if leftItem <= rightItem {
			merged = append(merged, leftItem)
			leftIndex++
		} else {
			merged = append(merged, rightItem)
			rightIndex++
		}
	}

	for leftIndex < leftSize {
		merged = append(merged, leftSlice[leftIndex])
		leftIndex++
	}

	for rightIndex < rightSize {
		merged = append(merged, rightSlice[rightIndex])
		rightIndex++
	}

	return merged
}
