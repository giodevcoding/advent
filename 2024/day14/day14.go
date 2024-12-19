package day14

import (
	. "aoc2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	Position, Velocity Vec2
}

const (
	SECONDS = 100
	WIDTH   = 101
	HEIGHT  = 103
)

var buf = bufio.NewReader(os.Stdin)

func BathroomRobotSafetyFactor(input []string) string {
	robots := getRobots(input)
    printRobots(robots)
    buf.ReadBytes('\n')

	for second := 1; second < math.MaxInt; second++ {
		for i := range robots {
			bot := &robots[i]

			newX := bot.Position.X + (bot.Velocity.X)
			newX = newX % WIDTH
			if newX < 0 {
				newX = WIDTH + newX
			}
			bot.Position.X = newX

			newY := bot.Position.Y + (bot.Velocity.Y)
			newY = newY % HEIGHT
			if newY < 0 {
				newY = HEIGHT + newY
			}
			bot.Position.Y = newY
		}
        if second == 77 || (second - 77) % WIDTH == 0 {
            printRobots(robots)
            fmt.Println("AFTER ", second, " SECONDS ^^^^")
            buf.ReadBytes('\n')
        }

	}
	quadrants := [4][]Robot{}
	for i := range quadrants {
		quadrants[i] = []Robot{}
	}

	for _, robot := range robots {
		xMid, yMid := WIDTH/2, HEIGHT/2
		x, y := robot.Position.X, robot.Position.Y

		switch {
		// Top Left
		case x < xMid && y < yMid:
			quadrants[0] = append(quadrants[0], robot)
		case x > xMid && y < yMid:
			quadrants[1] = append(quadrants[1], robot)
		case x < xMid && y > yMid:
			quadrants[2] = append(quadrants[2], robot)
		case x > xMid && y > yMid:
			quadrants[3] = append(quadrants[3], robot)
		}
	}

	safetyFactor := len(quadrants[0]) * len(quadrants[1]) * len(quadrants[2]) * len(quadrants[3])

	return strconv.Itoa(safetyFactor)
}

func getRobots(input []string) []Robot {
	robots := make([]Robot, len(input))
	for i, info := range input {
		regex := regexp.MustCompile(`(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
		matches := regex.FindAllStringSubmatch(info, -1)

		pX, _ := strconv.Atoi(matches[0][1])
		pY, _ := strconv.Atoi(matches[0][2])
		vX, _ := strconv.Atoi(matches[0][3])
		vY, _ := strconv.Atoi(matches[0][4])

		robots[i] = Robot{
			Position: Vec2{X: pX, Y: pY},
			Velocity: Vec2{X: vX, Y: vY},
		}
	}

	return robots
}

func printRobots(robots []Robot) {
	robotPositionCounts := map[Vec2]int{}
	for count, robot := range robots {
		count = robotPositionCounts[robot.Position]
		robotPositionCounts[robot.Position] = count + 1
	}

	robotMap := make([][]string, HEIGHT)
	for y := range robotMap {
		robotMap[y] = make([]string, WIDTH)
		for x := range robotMap[y] {
			if count, ok := robotPositionCounts[Vec2{X: x, Y: y}]; ok {
				robotMap[y][x] = strconv.Itoa(count)
			} else {
				robotMap[y][x] = "."
			}
		}
		fmt.Println(strings.Join(robotMap[y], ""))
	}

}
