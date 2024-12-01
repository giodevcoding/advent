package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
    path, pathErr :=filepath.Abs("./input.txt")

    if pathErr != nil {
        return
    }

    file, fileErr := os.Open(path)
    if fileErr != nil {
        return
    }
    defer file.Close()

    var lines []string
    var scanner = bufio.NewScanner(file)
    for  scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    for _, input := range lines {
        fmt.Println(NotQuiteLisp(input))
    }
}

func NotQuiteLisp(input string) string {
    floor := 0
    for i, c := range input {
        if c == '(' {
            floor += 1
        }
        if c == ')' {
            floor -= 1
        }

        if floor < 0 {
            return fmt.Sprintf("%d", i + 1)
        }
    }
    return fmt.Sprintf("%d", floor)
}
