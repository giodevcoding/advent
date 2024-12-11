package day8

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"strconv"
)

type Coordinate struct {
	X int
	Y int
}

func UniqueAntinodes(input []string) string {
    width, height := len(input[0]), len(input)
	antennae := findAntennae(input)
	antinodes := make([]Coordinate, 0)
	for _, coords := range antennae {
        newAntinodes := getAllAntinodes([]Coordinate{}, coords, 0, width, height)
        antinodes = append(antinodes, newAntinodes...)
	}

    antinodes = utils.Filter(antinodes, func(a Coordinate) bool {
        if a.X < 0 || a.X >= width {
            return false
        }
        if a.Y < 0 || a.Y >= height {
            return false
        }
        return true
    })
    antinodes = utils.RemoveDuplicates(antinodes)

    fmt.Println(antinodes)

	return strconv.Itoa(len(antinodes))
}

func getAllAntinodes(antinodes []Coordinate, antennae []Coordinate, index, width, height int) []Coordinate {
	if index == len(antennae)-1 {
		return antinodes
	}

	mainAntennae := antennae[index]
	for i := index + 1; i < len(antennae); i++ {
		checkAntennae := antennae[i]
        // Part 1
        //newAntinodes := getAntinodes(mainAntennae, checkAntennae)

        // Part 2
        newAntinodes := getAntinodesLine(mainAntennae, checkAntennae, width, height)
		antinodes = append(antinodes, newAntinodes[:]...)
	}
 
    return getAllAntinodes(antinodes, antennae, index+1, width, height)
}

func getAntinodes(antennae1 Coordinate, antennae2 Coordinate) (antinodes [2]Coordinate) {
	xDiff, yDiff := utils.IntAbs(antennae1.X-antennae2.X), utils.IntAbs(antennae1.Y-antennae2.Y)

	xModifier, yModifier := 1, 1
	if math.Signbit(float64(antennae1.X - antennae2.X)) {
		xModifier = -1
	}
	if math.Signbit(float64(antennae1.Y - antennae2.Y)) {
		yModifier = -1
	}

	antinodes[0].X = antennae1.X + (xDiff * xModifier)
	antinodes[1].X = antennae2.X - (xDiff * xModifier)

	antinodes[0].Y = antennae1.Y + (yDiff * yModifier)
	antinodes[1].Y = antennae2.Y - (yDiff * yModifier)
	return
}

func getAntinodesLine(antennae1 Coordinate, antennae2 Coordinate, width, height int) (antinodes []Coordinate) {
	xDiff, yDiff := utils.IntAbs(antennae1.X-antennae2.X), utils.IntAbs(antennae1.Y-antennae2.Y)

	xModifier, yModifier := 1, 1
	if math.Signbit(float64(antennae1.X - antennae2.X)) {
		xModifier = -1
	}
	if math.Signbit(float64(antennae1.Y - antennae2.Y)) {
		yModifier = -1
	}

    antinodePrev := antennae1
    antinodeNext := antennae2

    for antinodePrev.X >= 0 && antinodePrev.Y >= 0 {
        antinodes = append(antinodes, antinodePrev)
        antinodePrev.X += xDiff * xModifier
        antinodePrev.Y += yDiff * yModifier
    }

    for antinodeNext.X < width && antinodeNext.Y < height {
        antinodes = append(antinodes, antinodeNext)
        antinodeNext.X -= xDiff * xModifier
        antinodeNext.Y -= yDiff * yModifier
    }
    
	return
}

func findAntennae(input []string) map[rune][]Coordinate {
	antennae := make(map[rune][]Coordinate)
	for y, row := range input {
		for x, char := range row {
			if char != '.' {
				antennae[char] = append(antennae[char], Coordinate{x, y})
			}
		}
	}
	return antennae
}
