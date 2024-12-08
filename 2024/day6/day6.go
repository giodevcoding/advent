package day6

import (
	"bufio"
	"fmt"
	"maps"
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
	possibleObstructions := make(map[Coordinate]bool)
	positionsWalked := make(map[Coordinate][]Direction)
	width, height := len(input[0]), len(input)
	x, y, dir := getGuardStart(input)
	positionsWalked[Coordinate{x, y}] = append(positionsWalked[Coordinate{x, y}], dir)
	startX, startY := x, y

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

		if x+dir.X >= 0 && x+dir.X < width && y+dir.Y >= 0 && y+dir.Y < height {
			if input[y+dir.Y][x+dir.X] != '#' {
				if (x+dir.X == startX && y+dir.Y == startY) == false {
					if _, ok := positionsWalked[Coordinate{x + dir.X, y + dir.Y}]; !ok {
						possibleObstruction := getGuardPossibleObstruction(x, y, dir, input, positionsWalked)
						if possibleObstruction != nil {
							possibleObstructions[*possibleObstruction] = true
						}
					}
				}
			}
		}

		x, y = nextX, nextY
		positionsWalked[Coordinate{x, y}] = append(positionsWalked[Coordinate{x, y}], dir)
	}

	return strconv.Itoa(len(possibleObstructions))
}

func getGuardPossibleObstruction(startX, startY int, startDir Direction, input []string, positionsWalked map[Coordinate][]Direction) *Coordinate {
	internalPositionsWalked := maps.Clone(positionsWalked)
	width, height := len(input[0]), len(input)
	x, y, dir := startX, startY, turnRight(startDir)

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

		if slices.Contains(internalPositionsWalked[Coordinate{nextX, nextY}], dir) {
			fmt.Println("VALID OBSTRUCTION AT", startX+startDir.X, startY+startDir.Y)
			return &Coordinate{startX + startDir.X, startY + startDir.Y}
		}

		x, y = nextX, nextY
		internalPositionsWalked[Coordinate{x, y}] = append(internalPositionsWalked[Coordinate{x, y}], dir)

		inputCopy := make([]string, len(input))
		copy(inputCopy, input)

		runes := []rune(inputCopy[startY+startDir.Y])
		runes[startX+startDir.X] = '0'
		inputCopy[startY+startDir.Y] = string(runes)

		runes = []rune(inputCopy[y])
		runes[x] = '*'
		inputCopy[y] = string(runes)

		/*for _, line := range inputCopy {
			fmt.Println(line)
		}
		buf.ReadBytes('\n')*/
	}

	return nil
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
