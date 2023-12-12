package main

import (
	"advent/helpers"
	"advent/ten"
	"fmt"
	"path/filepath"
)

func main() {
	var absPath, _ = filepath.Abs("./ten/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	//ten.RunPartOne(lines)
	ten.RunPartTwo(lines)
}
