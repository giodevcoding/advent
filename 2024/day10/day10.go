package day10

import (
	"aoc2024/utils"
	"bufio"
	"os"
	"strconv"
)

var buf = bufio.NewReader(os.Stdin)

type Coordinate struct {
    X, Y int
}

func TrailheadScores(input []string) string {
    score := 0
    topMap := getTopographicalMap(input)
	width, height := len(topMap[0]), len(topMap)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if topMap[y][x] == 0 {
                score += len(utils.SliceToSet(getTrailheadReachableNines(x, y, -1, topMap)))
            }
        }
    }

	return strconv.Itoa(score)
}

func TrailheadRatings(input []string) string {
    score := 0
    topMap := getTopographicalMap(input)
	width, height := len(topMap[0]), len(topMap)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if topMap[y][x] == 0 {
                score += len(getTrailheadReachableNines(x, y, -1, topMap))
            }
        }
    }

	return strconv.Itoa(score)
}

func getTopographicalMap(input []string) [][]int {
	width, height := len(input[0]), len(input)

	topMap := make([][]int, height)
	for y := range topMap {
		topMap[y] = make([]int, width)
		for x := range topMap[y] {
			topMap[y][x] = utils.RuneToInt(rune(input[y][x]))
		}
	}

	return topMap
}

func getTrailheadReachableNines(x, y, lastHeight int, topMap [][]int) []Coordinate {
    if x < 0 || x >= len(topMap[0]) || y < 0 || y >= len(topMap) {
        return []Coordinate{}
    }


    currentHeight := topMap[y][x]
    if currentHeight - lastHeight != 1 {
        return []Coordinate{}
    } 

    if currentHeight == 9 {
        return []Coordinate{Coordinate{x, y}}
    }

    nines := []Coordinate{}
    nines = append(nines, getTrailheadReachableNines(x, y-1, currentHeight, topMap)...)
    nines = append(nines, getTrailheadReachableNines(x+1, y, currentHeight, topMap)...)
    nines = append(nines, getTrailheadReachableNines(x, y+1, currentHeight, topMap)...)
    nines = append(nines, getTrailheadReachableNines(x-1, y, currentHeight, topMap)...)
    return nines
}
