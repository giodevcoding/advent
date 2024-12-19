package day15

import (
	. "aoc2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var buf = bufio.NewReader(os.Stdin)

var UP = Vec2{X: 0, Y: -1}
var RIGHT = Vec2{X: 1, Y: 0}
var DOWN = Vec2{X: 0, Y: 1}
var LEFT = Vec2{X: -1, Y: 0}

func WarehouseRobotWork(input []string) string {
	warehouseMap, moves := parseInputData(input)
	player := findPlayerStart(warehouseMap)
	printMap(warehouseMap)

	for _, move := range moves {
		swapCoord := player
		x, y := player.X, player.Y

	checkLoop:
		for {
			x += move.X
			y += move.Y
			switch warehouseMap[y][x] {
			case '#':
				break checkLoop
			case 'O':
				continue checkLoop
			default: // '.'
				swapCoord.X, swapCoord.Y = x, y
				break checkLoop
			}
		}

		playerMoved := !swapCoord.Equals(player)

		for !swapCoord.Equals(player) {
			moveTowards := getMoveTowards(swapCoord, player)
			nextX, nextY := swapCoord.X+moveTowards.X, swapCoord.Y+moveTowards.Y
			warehouseMap[swapCoord.Y][swapCoord.X], warehouseMap[nextY][nextX] = warehouseMap[nextY][nextX], warehouseMap[swapCoord.Y][swapCoord.X]
			swapCoord.X = nextX
			swapCoord.Y = nextY
		}

		if playerMoved {
			player.X += move.X
			player.Y += move.Y
		}

	}
	printMap(warehouseMap)

	sum := getGPSCoordinates(warehouseMap)

	return strconv.Itoa(sum)
}

func WarehouseWideRobot(input []string) string {
	warehouseMap, moves := parseInputDataWide(input)
	player := findPlayerStart(warehouseMap)
	printMap(warehouseMap)

	for _, move := range moves {
		if shouldMoveWideBoxes(player, move, warehouseMap) {
			moveWideBox(player, move, &warehouseMap)
			player.AddMut(move)
		}
	}
	printMap(warehouseMap)

	sum := getGPSWideCoordinates(warehouseMap)

	return strconv.Itoa(sum)
}

func shouldMoveWideBoxes(player Vec2, move Vec2, warehouseMap [][]rune) bool {
	// Horizontal
	if move.Equals(LEFT) || move.Equals(RIGHT) {
		next := player.Add(move)

		switch warehouseMap[next.Y][next.X] {
		case '[', ']':
			// Add other side of the box
			next.AddMut(move)
			return shouldMoveWideBoxes(next, move, warehouseMap)
		case '.':
			return true
		case '#':
			return false
		}
		return false
	}
	// Vertical
	next := player.Add(move)
	switch warehouseMap[next.Y][next.X] {
	case '[':
		// Add other side of the box
		left := shouldMoveWideBoxes(next, move, warehouseMap)
		right := shouldMoveWideBoxes(next.Add(RIGHT), move, warehouseMap)
		return left && right
	case ']':
		// Add other side of the box
		left := shouldMoveWideBoxes(next.Add(LEFT), move, warehouseMap)
		right := shouldMoveWideBoxes(next, move, warehouseMap)
		return left && right
	case '.':
		return true
	case '#':
		return false
	}
	return false
}

func moveWideBox(player Vec2, move Vec2, warehouseMap *[][]rune) {
	// Horizontal
	if move.Equals(LEFT) || move.Equals(RIGHT) {
		next := player.Add(move)

		if (*warehouseMap)[next.Y][next.X] != '.' {
			moveWideBox(next, move, warehouseMap)
		}
		(*warehouseMap)[player.Y][player.X], (*warehouseMap)[next.Y][next.X] = (*warehouseMap)[next.Y][next.X], (*warehouseMap)[player.Y][player.X]
		return
	}

	// Vertical
	next := player.Add(move)
    
    if (*warehouseMap)[next.Y][next.X] != '.' {
        switch (*warehouseMap)[next.Y][next.X] {
        case '[':
            moveWideBox(next, move, warehouseMap)
            moveWideBox(next.Add(RIGHT), move, warehouseMap)
        case ']':
            moveWideBox(next.Add(LEFT), move, warehouseMap)
            moveWideBox(next, move, warehouseMap)
        }
    }
    (*warehouseMap)[player.Y][player.X], (*warehouseMap)[next.Y][next.X] = (*warehouseMap)[next.Y][next.X], (*warehouseMap)[player.Y][player.X]
}

func getMoveTowards(v1 Vec2, v2 Vec2) Vec2 {
	diffX, diffY := v2.X-v1.X, v2.Y-v1.Y
	if diffY < 0 {
		return UP
	}
	if diffY > 0 {
		return DOWN
	}
	if diffX < 0 {
		return LEFT
	}
	if diffX > 0 {
		return RIGHT
	}
	return Vec2{}
}

func findPlayerStart(warehouseMap [][]rune) Vec2 {
	for y := range warehouseMap {
		for x := range warehouseMap {
			if warehouseMap[y][x] == '@' {
				return Vec2{X: x, Y: y}
			}
		}
	}
	panic("Could not find player start!!")
}

func parseInputData(input []string) (warehouseMap [][]rune, moves []Vec2) {
	buildMap := true

	for _, line := range input {
		if len(line) == 0 {
			buildMap = !buildMap
		}

		runes := []rune(line)

		if buildMap {
			warehouseMap = append(warehouseMap, runes)
		} else {
			moves = append(moves, Map(runes, runeToMove)...)
		}
	}

	return
}

func parseInputDataWide(input []string) (warehouseMap [][]rune, moves []Vec2) {
	buildMap := true

	for _, line := range input {
		if len(line) == 0 {
			buildMap = !buildMap
		}

		if buildMap {
			wideLine := ""

			for _, char := range line {
				switch char {
				case '#':
					wideLine += "##"
				case 'O':
					wideLine += "[]"
				case '.':
					wideLine += ".."
				case '@':
					wideLine += "@."
				}
			}

			runes := []rune(wideLine)
			warehouseMap = append(warehouseMap, runes)
		} else {
			runes := []rune(line)
			moves = append(moves, Map(runes, runeToMove)...)
		}
	}

	return
}

func runeToMove(r rune) Vec2 {
	switch r {
	case '^':
		return UP
	case '>':
		return RIGHT
	case 'v':
		return DOWN
	case '<':
		return LEFT
	}
	return Vec2{}
}

func printMap(warehouseMap [][]rune) {
	fmt.Println("WAREHOUSE:")
	for _, line := range warehouseMap {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func getGPSCoordinates(warehouseMap [][]rune) (sum int) {
	for y := range warehouseMap {
		for x, r := range warehouseMap[y] {
			if r == 'O' {
				sum += (100 * y) + x
			}
		}
	}
	return
}

func getGPSWideCoordinates(warehouseMap [][]rune) (sum int) {
	for y := range warehouseMap {
		for x, r := range warehouseMap[y] {
			if r == '[' {
				sum += (100 * y) + x
			}
		}
	}
	return
}
