package eleven

import (
	"fmt"
	"math"
	"strings"
)

type (
	Galaxy struct {
		id, x, y int
	}
)

func RunPartOne(input []string) {
	fmt.Println(ShortestGalaxyPathsSum(input, 2))
}

func RunPartTwo(input []string) {
	fmt.Println(ShortestGalaxyPathsSum(input, 1000000))
}

func ShortestGalaxyPathsSum(input []string, expansionMultiplier int) int {
	galaxies := GalaxiesInUniverse(input)
	emptyRows, emptyCols := EmptyRowsAndCols(input)

	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := DistanceBetweenGalaxiesAfterExpansion(galaxies[i], galaxies[j], emptyRows, emptyCols, expansionMultiplier)
			sum += distance
		}
	}

	return sum
}

func DistanceBetweenGalaxiesAfterExpansion(galaxyA Galaxy, galaxyB Galaxy, emptyRows []int, emptyCols []int, expansionMultiplier int) int {
	emptyColsBetween, emptyRowsBetween := 0, 0
	for _, colIndex := range emptyCols {
		if colIndex > galaxyA.x && colIndex < galaxyB.x {
			emptyColsBetween++
		} else if colIndex > galaxyB.x && colIndex < galaxyA.x {
			emptyColsBetween++
        }
	}

	for _, rowIndex := range emptyRows {
		if rowIndex > galaxyA.y && rowIndex < galaxyB.y {
			emptyRowsBetween++
		} else if rowIndex > galaxyB.y && rowIndex < galaxyA.y {
			emptyRowsBetween++
        }
	}

	xDistance := int(math.Abs(float64(galaxyA.x - galaxyB.x))) + emptyColsBetween * (expansionMultiplier - 1)	
	yDistance := int(math.Abs(float64(galaxyA.y - galaxyB.y))) + emptyRowsBetween * (expansionMultiplier - 1)	

    return xDistance+yDistance

}

func EmptyRowsAndColsBetweenGalaxies(galaxyA Galaxy, galaxyB Galaxy, input []string) {

}

func GalaxiesInUniverse(input []string) (galaxies []Galaxy) {
	latestGalaxyId := 1

	for y := range input {
		for x, char := range input[y] {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{
					id: latestGalaxyId,
					x:  x,
					y:  y,
				})
				latestGalaxyId++
			}
		}
	}
	return galaxies
}

func ExpandUniverse(input []string) []string {
	expanded := input
	emptyRows, emptyCols := EmptyRowsAndCols(input)

	yOffset, xOffset := 0, 0
	_ = xOffset

	for y := 0; y < len(expanded); y++ {
		if len(emptyRows) > yOffset {
			if y == emptyRows[yOffset]+yOffset {
				expanded = AddEmptyRowAtIndex(expanded, y)
				yOffset++
			}
		}
	}

	for x := 0; x < len(expanded[0]); x++ {
		if len(emptyCols) > xOffset {
			if x == emptyCols[xOffset]+xOffset {
				expanded = AddEmptyColAtIndex(expanded, x)
				xOffset++
			}
		}
	}

	for _, line := range expanded {
		fmt.Println(line)
	}

	return expanded
}

func EmptyRowsAndCols(input []string) ([]int, []int) {
	emptyRowIndexes, emptyColIndexes := []int{}, []int{}

	// Empty Rows
	for i, row := range input {
		if !strings.Contains(row, "#") {
			emptyRowIndexes = append(emptyRowIndexes, i)
		}
	}

	// Empty Columns
	for i := 0; i < len(input[0]); i++ {
		empty := true
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				empty = false
				break
			}
		}

		if empty {
			emptyColIndexes = append(emptyColIndexes, i)
		}
	}

	return emptyRowIndexes, emptyColIndexes
}

func AddEmptyColAtIndex(slice []string, index int) []string {
	newCols := slice
	for i, line := range slice {
		newLine := fmt.Sprintf("%s%s%s", line[:index], ".", line[index:])
		newCols[i] = newLine
	}
	return newCols
}

func AddEmptyRowAtIndex(slice []string, index int) []string {
	emptyRow := strings.Repeat(".", len(slice[0]))
	newRows := append(slice[:index], emptyRow)
	newRows = append(newRows, slice[index:]...)
	return newRows
}
