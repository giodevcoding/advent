package day5

import (
	"aoc2024/utils"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func PrintUpdateOrdering(input []string) string {
    total := 0
    orderings, updates := extractOrderAndUpdates(input)
    
    for _, update := range updates {
        updateSorted := utils.MergeSort(update, printUpdateComparator(orderings))
        correct := reflect.DeepEqual(update, updateSorted)

        if correct {
            total += update[len(update)/2]
        }
    }

	return strconv.Itoa(total)
}

func PrintUpdateFixOrdering(input []string) string {
    total := 0
    orderings, updates := extractOrderAndUpdates(input)
    
    for _, update := range updates {
        updateSorted := utils.MergeSort(update, printUpdateComparator(orderings))
        correct := reflect.DeepEqual(update, updateSorted)

        if !correct {
            total += updateSorted[len(update)/2]
        }
    }

	return strconv.Itoa(total)
}

func printUpdateComparator(orderings map[int][]int) (func(a int, b int) int){
    return func(a int, b int) int {
        if (slices.Contains(orderings[b], a)) {
            return -1
        } else if (slices.Contains(orderings[a], b)) {
            return 1
        }
        return 0
    }
}

func extractOrderAndUpdates(input []string) (map[int][]int, [][]int) {
	orders := make(map[int][]int)
	switchIndex := 0

	for i, line := range input {
		if len(line) == 0 {
			switchIndex = i + 1
			break
		}

		orderingInfo := utils.Map(strings.Split(line, "|"), func(info string) int {
			num, err := strconv.Atoi(info)
			if err != nil {
				panic(err)
			}
			return num
		})

        X, Y := orderingInfo[0], orderingInfo[1]

        orders[Y] = append(orders[Y], X)
	}

	updates := utils.Map(input[switchIndex:len(input)], func(updatesRaw string) []int {
		return utils.Map(strings.Split(updatesRaw, ","), func(update string) int {
			num, err := strconv.Atoi(update)
			if err != nil {
				panic(err)
			}
			return num
		})

	})

	return orders, updates
}
