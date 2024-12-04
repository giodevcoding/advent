package day4

import (
	"strconv"
)

type Direction struct {
    Name string
	X int
	Y int
}

var UP = Direction{
    Name: "UP",
	X: 0,
	Y: -1,
}
var UP_RIGHT = Direction{
    Name: "UP_RIGHT",
	X: 1,
	Y: -1,
}
var RIGHT = Direction{
    Name: "RIGHT",
	X: 1,
	Y: 0,
}
var DOWN_RIGHT = Direction{
    Name: "DOWN_RIGHT",
	X: 1,
	Y: 1,
}
var DOWN = Direction{
    Name: "DOWN",
	X: 0,
	Y: 1,
}
var DOWN_LEFT = Direction{
    Name: "DOWN_LEFT",
	X: -1,
	Y: 1,
}
var LEFT = Direction{
    Name: "LEFT",
	X: -1,
	Y: 0,
}
var UP_LEFT = Direction{
    Name: "UP_LEFT",
	X: -1,
	Y: -1,
}

var directions = [8]Direction{UP, UP_RIGHT, RIGHT, DOWN_RIGHT, DOWN, DOWN_LEFT, LEFT, UP_LEFT}

func XmasWordSearch(input []string) string {
	total := 0
	searchStr := [4]rune{'X', 'M', 'A', 'S'}
	for y, row := range input {
		for x, rn := range row {
			if rn != searchStr[0] {
				continue
			}

            //fmt.Printf("Found 'X' at x: %d y: %d\n", x, y)
			// Found X, look for rest of word in all 8 directions
            for _, dir := range directions { 
                found := true
				for searchIdx := 1; searchIdx < len(searchStr); searchIdx++ {
					nextY, nextX := y+(dir.Y*searchIdx), x+(dir.X*searchIdx)
                    searchRn := rune(searchStr[searchIdx])

                    // fmt.Printf("Searching for '%c' in direction %s at x: %d y: %d\n", searchRn, dir.Name, nextX, nextY)

					if nextY < 0 || nextY >= len(input) {
                        found = false
						break
					}

					if nextX < 0 || nextX >= len(row) {
                        found = false
						break
					}

                    nextRn := rune(input[nextY][nextX])

                    if (nextRn != searchRn) {
                        found = false
                        break
                    }
				}

                if found {
                    total++
                }
			}
		}
	}
	return strconv.Itoa(total)
}
