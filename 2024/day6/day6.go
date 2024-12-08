package day6

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Coordinate struct {
	X int
	Y int
}

type Direction Coordinate

var buf = bufio.NewReader(os.Stdin)

func GuardPath(input []string) string {
	positionsWalked := make(map[Coordinate]bool)
	width, height := len(input[0]), len(input)
	coord, dir := getGuardStart(input)
	positionsWalked[coord] = true

	for {
		nextX, nextY := coord.X+dir.X, coord.Y+dir.Y
		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			break
		}
		nextStep := rune(input[nextY][nextX])

		for nextStep == '#' {
			dir = turnRight(dir)
			nextX, nextY = coord.X+dir.X, coord.Y+dir.Y
			if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
				break
			}
			nextStep = rune(input[nextY][nextX])
		}

		coord = Coordinate{nextX, nextY}
		positionsWalked[coord] = true
	}

	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	for coord := range positionsWalked {
		runes := []rune(inputCopy[coord.Y])
		runes[coord.X] = 'X'
		inputCopy[coord.Y] = string(runes)
	}

	for _, line := range inputCopy {
		fmt.Println(line)
	}

	return strconv.Itoa(len(positionsWalked))
}

func GuardPossibleObstructions(input []string) string {
	obstructions := 0
	loc, dir := getGuardStart(input)

	for y, line := range input {
		for x := range line {
			if (loc == Coordinate{x, y}) {
				continue
			}
			steps := make(map[Coordinate][]Direction)
			inputCopy := make([]string, len(input))
			copy(inputCopy, input)

			runes := []rune(inputCopy[y])
			runes[x] = '0'
			inputCopy[y] = string(runes)

			if doesGuardLoop(loc, dir, inputCopy, steps) {
				obstructions++
			}
		}
	}

	return strconv.Itoa(obstructions)
}

func doesGuardLoop(guardLoc Coordinate, dir Direction, guardMap []string, guardSteps map[Coordinate][]Direction) bool {
	width, height := len(guardMap[0]), len(guardMap)

	if slices.Contains(guardSteps[guardLoc], dir) {
		return true
	}

	guardSteps[guardLoc] = append(guardSteps[guardLoc], dir)

	nextStep := Coordinate{guardLoc.X + dir.X, guardLoc.Y + dir.Y}

	if nextStep.X < 0 || nextStep.X >= width || nextStep.Y < 0 || nextStep.Y >= height {
		return false
	}

	if slices.Contains([]rune{'#', '0'}, rune(guardMap[nextStep.Y][nextStep.X])) {
		dir = turnRight(dir)
	} else {
		guardLoc = Coordinate{guardLoc.X + dir.X, guardLoc.Y + dir.Y}
	}

	return doesGuardLoop(guardLoc, dir, guardMap, guardSteps)
}

func showSteps(guardLoc Coordinate, steps map[Coordinate][]Direction, guardMap []string) {

	mapCopy := make([]string, len(guardMap))
	copy(mapCopy, guardMap)

	for step, dirs := range steps {
		var pathChar rune
		combinedDir := Direction{0, 0}
		for _, d := range dirs {
			combinedDir.X += utils.IntAbs(d.X)
			combinedDir.Y += utils.IntAbs(d.Y)
		}
		if combinedDir.X == 0 && combinedDir.Y > 0 {
			pathChar = '|'
		} else if combinedDir.Y == 0 && combinedDir.X > 0 {
			pathChar = '-'
		} else {
			pathChar = '+'
		}
		runes := []rune(mapCopy[step.Y])
		runes[step.X] = pathChar
		mapCopy[step.Y] = string(runes)
	}

	runes := []rune(mapCopy[guardLoc.Y])
	runes[guardLoc.X] = '*'
	mapCopy[guardLoc.Y] = string(runes)

	for _, line := range mapCopy {
		fmt.Println(line)
	}

	//buf.ReadBytes('\n')

}

func getGuardStart(input []string) (coord Coordinate, dir Direction) {
	for i, row := range input {
		for j, char := range row {
			if char == '^' {
				coord, dir = Coordinate{j, i}, Direction{0, -1}
				return
			}
		}
	}
	return
}

func turnRight(currentDir Direction) (newDirection Direction) {
	flip := currentDir.X == 0
	newDirection = Direction{currentDir.Y, currentDir.X}
	if flip {
		newDirection.X *= -1
	}
	return
}
