package day4

import (
	"strconv"
)

type Direction struct {
	Name string
	X    int
	Y    int
}

type DirectionPair [2]Direction

var UP = Direction{
	Name: "UP",
	X:    0,
	Y:    -1,
}
var UP_RIGHT = Direction{
	Name: "UP_RIGHT",
	X:    1,
	Y:    -1,
}
var RIGHT = Direction{
	Name: "RIGHT",
	X:    1,
	Y:    0,
}
var DOWN_RIGHT = Direction{
	Name: "DOWN_RIGHT",
	X:    1,
	Y:    1,
}
var DOWN = Direction{
	Name: "DOWN",
	X:    0,
	Y:    1,
}
var DOWN_LEFT = Direction{
	Name: "DOWN_LEFT",
	X:    -1,
	Y:    1,
}
var LEFT = Direction{
	Name: "LEFT",
	X:    -1,
	Y:    0,
}
var UP_LEFT = Direction{
	Name: "UP_LEFT",
	X:    -1,
	Y:    -1,
}

var directions = [8]Direction{UP, UP_RIGHT, RIGHT, DOWN_RIGHT, DOWN, DOWN_LEFT, LEFT, UP_LEFT}
var xDirections = [2]DirectionPair{{UP_RIGHT, DOWN_LEFT}, {UP_LEFT, DOWN_RIGHT}}

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

					if nextRn != searchRn {
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

func XDashMasWordSearch(input []string) string {
	total := 0
	for y, row := range input {
	    for x, rn := range row {
            if rn != 'A' {
                continue
            }

            found := true
            for _, dirPair := range xDirections {
                x1, y1 := x+dirPair[0].X, y+dirPair[0].Y
                x2, y2 := x+dirPair[1].X, y+dirPair[1].Y
                
                if x1 < 0 || x1 >= len(row) {
                    found = false
                    break
                }

                if y1 < 0 || y1 >= len(input) {
                    found = false
                    break
                }

                if x2 < 0 || x2 >= len(row) {
                    found = false
                    break
                }

                if y2 < 0 || y2 >= len(input) {
                    found = false
                    break
                }

                rn1, rn2 := rune(input[y1][x1]), rune(input[y2][x2])

                isMas := (rn1 == 'M' && rn2 == 'S') || (rn1 == 'S' && rn2 == 'M')
                
                if !isMas {
                    found = false
                    break
                }
            }

            if found {
                total++
            }
        }
    }
	return strconv.Itoa(total)
}
