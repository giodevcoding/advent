package ten

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

type (
	Direction int
	PipeShape [2]Direction
	Pipe      struct {
		char       rune
		x, y       int
		directions PipeShape
	}
)

const (
	UP    Direction = 0
	LEFT  Direction = 1
	DOWN  Direction = 2
	RIGHT Direction = 3
)

func RunPartOne(input []string) {
	fmt.Println(FarthestDistance(input))
}

func RunPartTwo(input []string) {
	fmt.Println(EnclosedTiles(input))
}

func FarthestDistance(input []string) int {
	runeMap := ConvertStringsToRuneMap(input)
	pipeLoop := PipeLoop(runeMap)
	length := len(pipeLoop)
	return length/2 + length%2
}

func EnclosedTiles(input []string) int {
	pipeLoop := PipeLoop(ConvertStringsToRuneMap(input))

	area := PipeLoopInternalArea(pipeLoop)
	internalPoints := area + 1.0 - (float64(len(pipeLoop)) / 2)

	return int(internalPoints)
}

func PipeLoopInternalArea(pipeLoop []Pipe) float64 {

	leftLace, rightLace := 0.0, 0.0

	for i := 0; i < len(pipeLoop); i++ {
		var pipe1, pipe2 Pipe
		pipe1 = pipeLoop[i]
		if i >= len(pipeLoop)-1 {
			pipe2 = pipeLoop[0]
		} else {
			pipe2 = pipeLoop[i+1]
		}

		leftLace += float64(pipe1.x * pipe2.y)
		rightLace += float64(pipe1.y * pipe2.x)
	}

	return math.Abs((leftLace - rightLace) / 2)
}

func CleanMap(runeMap [][]rune) ([][]rune, []Pipe) {
	cleanMap := runeMap

	pipeLoop := PipeLoop(runeMap)
	for y := range cleanMap {
		for x := range cleanMap[y] {
			if !slices.ContainsFunc(pipeLoop, func(pipe Pipe) bool {
				return pipe.x == x && pipe.y == y
			}) {
				cleanMap[y][x] = '.'
			}
		}
	}
	return cleanMap, pipeLoop
}

func PipeLoop(runeMap [][]rune) []Pipe {
	x, y := GetStartingCoordinates(runeMap)
	traversedPipes := []Pipe{}

	startingPipe := Pipe{
		x: x,
		y: y,
	}

	currentX, currentY, prevDir := FirstConnectedPipeLocation(x, y, runeMap)
	startingPipe.directions[0] = prevDir

	for runeMap[currentY][currentX] != 'S' {
		pipe := GetPipe(currentX, currentY, runeMap)
		traversedPipes = append(traversedPipes, pipe)

		nextDir := PipeOtherSide(pipe.directions, OppositeDirection(prevDir))

		nextX, nextY := NextPipeLocation(currentX, currentY, nextDir)
		currentX, currentY, prevDir = nextX, nextY, nextDir
	}

	startingPipe.directions[1] = OppositeDirection(prevDir)
	startingPipe.char = PipeCharFromShape(startingPipe.directions)

	traversedPipes = append([]Pipe{startingPipe}, traversedPipes...)

	return traversedPipes
}

func FirstConnectedPipeLocation(x, y int, runeMap [][]rune) (int, int, Direction) {
	if x > 0 {
		if pipe, err := GetPipeShape(runeMap[y][x-1]); err == nil {
			if PipeContains(pipe, RIGHT) {
				return x - 1, y, LEFT
			}
		}
	}

	if x < len(runeMap[0])-1 {
		if pipe, err := GetPipeShape(runeMap[y][x+1]); err == nil {
			if PipeContains(pipe, LEFT) {
				return x + 1, y, RIGHT
			}
		}
	}

	if y > 0 {
		if pipe, err := GetPipeShape(runeMap[y-1][x]); err == nil {
			if PipeContains(pipe, DOWN) {
				return x, y - 1, UP
			}
		}
	}

	if y < len(runeMap)-1 {
		if pipe, err := GetPipeShape(runeMap[y+1][x]); err == nil {
			if PipeContains(pipe, UP) {
				return x, y + 1, DOWN
			}
		}
	}

	return x, y, -1
}

func GetStartingCoordinates(pipeMap [][]rune) (x, y int) {
	for i, rowPipes := range pipeMap {
		for j, columnPipe := range rowPipes {
			if columnPipe == 'S' {
				x, y = j, i
				return x, y
			}
		}
	}
	return
}

func ConvertStringsToRuneMap(input []string) [][]rune {
	result := make([][]rune, len(input))

	for i := range input {
		result[i] = []rune(input[i])
	}

	return result
}

func GetPipeShape(char rune) (PipeShape, error) {
	switch char {
	case '-':
		return PipeShape{LEFT, RIGHT}, nil
	case '|':
		return PipeShape{UP, DOWN}, nil
	case '7':
		return PipeShape{LEFT, DOWN}, nil
	case 'J':
		return PipeShape{LEFT, UP}, nil
	case 'L':
		return PipeShape{RIGHT, UP}, nil
	case 'F':
		return PipeShape{RIGHT, DOWN}, nil
	}

	return PipeShape{}, errors.New(fmt.Sprintf("%s is not a pipe", string(char)))
}

func GetPipe(x, y int, runeMap [][]rune) Pipe {
	directions, _ := GetPipeShape(runeMap[y][x])
	return Pipe{
		x:          x,
		y:          y,
		directions: directions,
		char:       runeMap[y][x],
	}
}

func PipeOtherSide(pipe PipeShape, currentSide Direction) Direction {
	if pipe[0] == currentSide {
		return pipe[1]
	}
	if pipe[1] == currentSide {
		return pipe[0]
	}
	return -1
}

func PipeContains(pipe PipeShape, dir Direction) bool {
	return pipe[0] == dir || pipe[1] == dir
}

func OppositeDirection(dir Direction) Direction {
	switch dir {
	case UP:
		return DOWN
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	}
	return -1
}

func PipeCharFromShape(pipeShape PipeShape) rune {
	if PipeContains(pipeShape, LEFT) && PipeContains(pipeShape, RIGHT) {
		return '-'
	}
	if PipeContains(pipeShape, UP) && PipeContains(pipeShape, DOWN) {
		return '|'
	}
	if PipeContains(pipeShape, LEFT) && PipeContains(pipeShape, DOWN) {
		return '7'
	}
	if PipeContains(pipeShape, LEFT) && PipeContains(pipeShape, UP) {
		return 'J'
	}
	if PipeContains(pipeShape, RIGHT) && PipeContains(pipeShape, UP) {
		return 'L'
	}
	if PipeContains(pipeShape, RIGHT) && PipeContains(pipeShape, DOWN) {
		return 'F'
	}
	return '.'
}

func NextPipeLocation(currentX, currentY int, dir Direction) (int, int) {
	newX, newY := currentX, currentY
	switch dir {
	case UP:
		newY--
	case DOWN:
		newY++
	case LEFT:
		newX--
	case RIGHT:
		newX++
	}
	return newX, newY
}

func NextDirection(currentX, currentY int, dir Direction, runeMap [][]rune) Direction {
	pipe := GetPipe(currentX, currentY, runeMap)
	return PipeOtherSide(pipe.directions, OppositeDirection(dir))
}
