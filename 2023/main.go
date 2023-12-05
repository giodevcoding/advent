package main

import (
	"advent/helpers"
	"advent/five"
	"fmt"
	"path/filepath"
)

func main() {
	var absPath, _ = filepath.Abs("./five/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

    //five.RunPartOne(lines)
    five.RunPartTwo(lines)
}
