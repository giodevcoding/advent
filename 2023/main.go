package main

import (
	"advent/helpers"
	"advent/eleven"
	"fmt"
	"path/filepath"
)

func main() {
	var absPath, _ = filepath.Abs("./eleven/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	//eleven.RunPartOne(lines)
	eleven.RunPartTwo(lines)
}
