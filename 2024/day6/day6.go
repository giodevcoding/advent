package day6

import (
	"bufio"
	"fmt"
	"maps"
	"math"
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
	x, y, dir := getGuardStart(input)
	positionsWalked[Coordinate{x, y}] = true

	for {
		nextX, nextY := x+dir.X, y+dir.Y
		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			break
		}
		nextStep := rune(input[nextY][nextX])

		for nextStep == '#' {
			dir = turnRight(dir)
			nextX, nextY = x+dir.X, y+dir.Y
			if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
				break
			}
			nextStep = rune(input[nextY][nextX])
		}

		x, y = nextX, nextY
		positionsWalked[Coordinate{x, y}] = true
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
	positionsWalked := make(map[Coordinate][]Direction)
	width, height := len(input[0]), len(input)
	x, y, dir := getGuardStart(input)
	startX, startY := x, y
	c := 0
    positionsWalked[Coordinate{x, y}] = append(positionsWalked[Coordinate{x, y}], dir)

	if _, obstructionFound := getGuardPossibleObstruction(startX, startY, x, y, dir, positionsWalked, input); obstructionFound {
		c++
	}

	for {

		nextX, nextY := x+dir.X, y+dir.Y
		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			break
		}

		nextStep := rune(input[nextY][nextX])

		for nextStep == '#' {
			positionsWalked[Coordinate{x, y}] = append(positionsWalked[Coordinate{x, y}], dir)
			dir = turnRight(dir)
			positionsWalked[Coordinate{x, y}] = append(positionsWalked[Coordinate{x, y}], dir)
			nextX, nextY = x+dir.X, y+dir.Y
			if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
				break
			}
			nextStep = rune(input[nextY][nextX])
		}

		x, y = nextX, nextY
		positionsWalked[Coordinate{x, y}] = append(positionsWalked[Coordinate{x, y}], dir)

		if _, obstructionFound := getGuardPossibleObstruction(startX, startY, x, y, dir, positionsWalked, input); obstructionFound {
			c++
		}
	}

	return strconv.Itoa(c)
}

func getGuardPossibleObstruction(startX, startY, x, y int, dir Direction, positionsWalked map[Coordinate][]Direction, input []string) (Coordinate, bool) {
	internalPositionsWalked := maps.Clone(positionsWalked)
	width, height := len(input[0]), len(input)
	obsX, obsY := x+dir.X, y+dir.Y

	// Don't try obstrcution that is outside of map
	if obsX < 0 || obsX >= width || obsY < 0 || obsY >= height {
		return Coordinate{}, false
	}

	// Don't try obstruction that is that on guard starting spot
	if obsX == startX && obsY == startY {
		return Coordinate{}, false
	}

	// If obstruction is in front, turn instead
	if input[obsY][obsX] == '#' {
		internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
		dir = turnRight(dir)
		internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
		obsX, obsY = x+dir.X, y+dir.Y
	}

	// Don't try obstruction on previously traversed path
	if dirs := positionsWalked[Coordinate{x + dir.X, y + dir.Y}]; dirs != nil {
		return Coordinate{}, false
	}

	internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
	dir = turnRight(dir)
	internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
outer:
	for {
		nextX, nextY := x+dir.X, y+dir.Y
		if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
			break
		}

		nextStep := rune(input[nextY][nextX])

		for nextStep == '#' {
			internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
			dir = turnRight(dir)
			internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
			nextX, nextY = x+dir.X, y+dir.Y
			if nextX < 0 || nextX >= width || nextY < 0 || nextY >= height {
				break outer
			}
			nextStep = rune(input[nextY][nextX])
		}

		x, y = nextX, nextY

		for _, posDir := range internalPositionsWalked[Coordinate{x, y}] {
			if posDir.X == dir.X && posDir.Y == dir.Y {
				fmt.Println("FOUND POSSIBLE OBSTRUCTION AT", obsX, obsY)
				showObstruction(x, y, obsX, obsY, internalPositionsWalked, input)
				return Coordinate{obsX, obsY}, true
			}
		}

		internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)
		//showObstruction(x, y, obsX, obsY, internalPositionsWalked, input)
	}
	return Coordinate{}, false
}

func showObstruction(guardX, guardY, obsX, obsY int, positionsWalked map[Coordinate][]Direction, input []string) {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	for coord, dirs := range positionsWalked {
		runes := []rune(inputCopy[coord.Y])
		combDir := Direction{}
		for _, dir := range dirs {
			combDir.X += int(math.Abs(float64(dir.X)))
			combDir.Y += int(math.Abs(float64(dir.Y)))
		}
		if combDir.X == 0 && combDir.Y > 0 {
			runes[coord.X] = '|'
		} else if combDir.X > 0 && combDir.Y == 0 {
			runes[coord.X] = '-'
		} else {
			runes[coord.X] = '+'
		}
		inputCopy[coord.Y] = string(runes)
	}

	runes := []rune(inputCopy[obsY])
	runes[obsX] = 'O'
	inputCopy[obsY] = string(runes)

	runes = []rune(inputCopy[guardY])
	runes[guardX] = '*'
	inputCopy[guardY] = string(runes)

	for _, line := range inputCopy {
		fmt.Println(line)
	}

	//buf.ReadBytes('\n')
}

func getGuardStart(input []string) (x, y int, dir Direction) {
	guardChars := []rune{'^', '>', 'v', '<'}
	for i, row := range input {
		for j, char := range row {
			if slices.Contains(guardChars, char) {
				x, y, dir = j, i, getStartingDirection(char)
				return
			}
		}
	}
	return
}

func getStartingDirection(char rune) Direction {
	switch char {
	case '^':
		return Direction{0, -1}
	case '>':
		return Direction{1, 0}
	case 'v':
		return Direction{0, 1}
	case '<':
		return Direction{-1, 0}
	}

	return Direction{0, 0}
}

func turnRight(currentDir Direction) (newDirection Direction) {
	flip := currentDir.X == 0
	newDirection = Direction{currentDir.Y, currentDir.X}
	if flip {
		newDirection.X *= -1
	}
	return
}
