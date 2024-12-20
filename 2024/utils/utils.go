package utils

import (
	"slices"
	"strconv"
	"strings"
)

type Vec2 struct {
	X, Y int
}

func (v *Vec2) AddMut(v2 Vec2) {
	(*v).X += v2.X
	(*v).Y += v2.Y
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vec2) AddInt(x int, y int) Vec2 {
	v.X += x
	v.Y += y
	return v
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v Vec2) SubInt(x int, y int) Vec2 {
	v.X -= x
	v.Y -= y
	return v
}

func (v Vec2) Equals(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vec2) EqualsInt(x int, y int) bool {
	return v.X == x && v.Y == y
}

func SliceToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, item := range slice {
		set[item] = true
	}
	return set
}

func SetToSlice[T comparable](set map[T]bool) []T {
	slice := make([]T, 0)
	for item := range set {
		slice = append(slice, item)
	}
	return slice
}

func RemoveDuplicates[T comparable](slice []T) []T {
	return SetToSlice(SliceToSet(slice))
}

func IntAbs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func ForEach[T any](arr []T, forEachFunc func(T)) {
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

func RemoveFromSlice[T comparable](slice []T, element T) []T {
    index := slices.Index(slice, element)
    if index != -1 {
        return append(slice[:index], slice[index+1:]...)
    }
    return slice
}

func RuneToInt(r rune) int {
	return int(r - '0')
}
