package utils

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Comparator[T any] func(a, b T) int

func MergeSort[T any](arr []T, comp Comparator[T]) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := int(math.Floor(float64(len(arr)) / 2))

	leftSlice := MergeSort(arr[0:mid], comp)
	rightSlice := MergeSort(arr[(mid):len(arr)], comp)
	return merge(leftSlice, rightSlice, comp)
}

func merge[T any](leftSlice []T, rightSlice []T, comp Comparator[T]) []T {
	leftSize, rightSize := len(leftSlice), len(rightSlice)
	merged := make([]T, 0, leftSize+rightSize)

	leftIndex, rightIndex := 0, 0

	for leftIndex < leftSize && rightIndex < rightSize {
		leftItem, rightItem := leftSlice[leftIndex], rightSlice[rightIndex]
		sort := comp(leftItem, rightItem)
		if sort <= 0 {
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

type Number interface {
	constraints.Integer | constraints.Float
}

func NumberAsc[T Number](a, b T) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func NumberDesc[T interface {
	constraints.Integer | constraints.Float
}](a, b T) int {
	switch {
	case a < b:
		return 1
	case a > b:
		return -1
	default:
		return 0
	}
}
