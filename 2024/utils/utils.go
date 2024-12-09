package utils

import (
	"strconv"
	"strings"
)

type Slice[T any] []T

func IntAbs (num int) int {
    if num < 0 {
        return num * -1
    }
    return num
}

func ForEach[T any](arr []T, forEachFunc func(T)){
	for _, value := range arr {
        forEachFunc(value)
	}
}

func Filter[T any](arr []T, filterFunc func(T) bool) []T {
	result := []T{}
	for _, value := range arr {
		if filterFunc(value) {
			result = append(result, value)
		}
	}
	return result
}


func Map[T, U any](arr []T, mapFunc func(T) U) []U {
	result := make([]U, len(arr))
	for i, value := range arr {
		result[i] = mapFunc(value)
	}
	return result
}

func StringToNumList(str string) []int {
	split := strings.Split(str, " ")
	split = Filter(split, func(s string) bool {
		return len(strings.TrimSpace(s)) > 0
	})
	nums := Map(split, func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})
	return nums
}
