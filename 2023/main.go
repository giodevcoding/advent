package main

import (
	"advent/helpers"
	"advent/three"
	"fmt"
	"path/filepath"
)

func main() {
	var absPath, _ = filepath.Abs("./three/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

    //three.RunPartOne(lines)
    three.RunPartTwo(lines)
}
