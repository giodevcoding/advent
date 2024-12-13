package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	X, Y int
}

type Region struct {
	RegionType rune
	Area       int
	Perimeter  int
	Coords     []Coordinate
}

var buf = bufio.NewReader(os.Stdin)

func CalculateFenceCost(input []string) string {
	alreadyFoundSpots := make(map[Coordinate]bool)
	regions := []Region{}
	for y := range input {
		for x := range input[0] {
			regionType := rune(input[y][x])
			coords, area, perimeter := inspectRegion(regionType, x, y, input, &alreadyFoundSpots)
			if len(coords) > 0 {
				regions = append(regions, Region{regionType, area, perimeter, coords})
			}
		}
	}

	fenceCosts := 0
	for _, region := range regions {
        // Part 1
		// fenceCosts += region.Area * region.Perimeter

        // Part 2
        sides := getRegionSideCount(region, input)
        regionCost := region.Area * sides
        printRegion(region, sides, regionCost)
        fenceCosts += regionCost
	}

	return strconv.Itoa(fenceCosts)
}

func getRegionSideCount(region Region, gardenMap []string) (sides int) {
	for _, coord := range region.Coords {
		x, y, regionType := coord.X, coord.Y, region.RegionType
		// Top Left
		if !isInRegion(x-1, y, regionType, gardenMap) &&
			!isInRegion(x, y-1, regionType, gardenMap) {
			sides++
		} else if isInRegion(x-1, y, regionType, gardenMap) &&
			!isInRegion(x-1, y-1, regionType, gardenMap) &&
			isInRegion(x, y-1, regionType, gardenMap) {
			sides++
		}
		// Top Right
		if !isInRegion(x+1, y, regionType, gardenMap) &&
			!isInRegion(x, y-1, regionType, gardenMap) {
			sides++
		} else if isInRegion(x+1, y, regionType, gardenMap) &&
			!isInRegion(x+1, y-1, regionType, gardenMap) &&
			isInRegion(x, y-1, regionType, gardenMap) {
			sides++
		}
		// Bottom Left
		if !isInRegion(x-1, y, regionType, gardenMap) &&
			!isInRegion(x, y+1, regionType, gardenMap) {
			sides++
		} else if isInRegion(x-1, y, regionType, gardenMap) &&
			!isInRegion(x-1, y+1, regionType, gardenMap) &&
			isInRegion(x, y+1, regionType, gardenMap) {
			sides++
		}
		// Bottom Right
		if !isInRegion(x+1, y, regionType, gardenMap) &&
			!isInRegion(x, y+1, regionType, gardenMap) {
			sides++
		} else if isInRegion(x+1, y, regionType, gardenMap) &&
			!isInRegion(x+1, y+1, regionType, gardenMap) &&
			isInRegion(x, y+1, regionType, gardenMap) {
			sides++
		}
	}
	return
}

func printRegion(region Region, sides, regionCost int) {
    smallestX, smallestY, largestX, largestY := math.MaxInt, math.MaxInt, 0, 0
    for _, coord := range region.Coords {
        smallestX = int(math.Min(float64(smallestX), float64(coord.X)))
        smallestY = int(math.Min(float64(smallestY), float64(coord.Y)))
        largestX = int(math.Max(float64(largestX), float64(coord.X)))
        largestY = int(math.Max(float64(largestY), float64(coord.Y)))
    }

    width, height := (largestX - smallestX)+1, (largestY - smallestY)+1
    regionMap := make([][]rune, height)
    for i := range regionMap {
        line := make([]rune, width)
        for j := range line {
            line[j] = '.'
        }
        regionMap[i] = line
    }

    for _, coord := range region.Coords {
        regionMap[coord.Y-smallestY][coord.X-smallestX] = region.RegionType
    }

    fmt.Println("--- Region: ", string(region.RegionType), " ---")
    fmt.Println()
    for _, line := range regionMap {
        fmt.Println(string(line))
    }

    fmt.Println("Area: ", region.Area)
    fmt.Println("Sides: ", sides)
    fmt.Println("Cost: ", regionCost)
    fmt.Println()
    buf.ReadBytes('\n')
}

func isInRegion(x, y int, regionType rune, gardenMap []string) bool {
	width, height := len(gardenMap[0]), len(gardenMap)
	if x < 0 || x >= width || y < 0 || y >= height {
		return false
	}

	return rune(gardenMap[y][x]) == regionType
}

func inspectRegion(regionType rune, x, y int, gardenMap []string, alreadyFoundSpots *map[Coordinate]bool) (coords []Coordinate, area, perimeter int) {
	// End of bounds
	if x < 0 || x >= len(gardenMap[0]) || y < 0 || y >= len(gardenMap) {
		perimeter = 1
		return coords, area, perimeter
	}

	locToCheck := rune(gardenMap[y][x])
	if locToCheck != regionType {
		perimeter = 1
		return coords, area, perimeter
	}

	if _, ok := (*alreadyFoundSpots)[Coordinate{x, y}]; ok {
		return
	}

	area = 1
	coords = append(coords, Coordinate{x, y})
	(*alreadyFoundSpots)[Coordinate{x, y}] = true

	//Up
	upCoords, upArea, upPerimeter := inspectRegion(regionType, x, y-1, gardenMap, alreadyFoundSpots)
	coords = append(coords, upCoords...)
	area += upArea
	perimeter += upPerimeter
	//Right
	rightCoords, rightArea, rightPerimeter := inspectRegion(regionType, x+1, y, gardenMap, alreadyFoundSpots)
	coords = append(coords, rightCoords...)
	area += rightArea
	perimeter += rightPerimeter
	//Down
	downCoords, downArea, downPerimeter := inspectRegion(regionType, x, y+1, gardenMap, alreadyFoundSpots)
	coords = append(coords, downCoords...)
	area += downArea
	perimeter += downPerimeter
	//Left
	leftCoords, leftArea, leftPerimeter := inspectRegion(regionType, x-1, y, gardenMap, alreadyFoundSpots)
	coords = append(coords, leftCoords...)
	area += leftArea
	perimeter += leftPerimeter
	return
}
