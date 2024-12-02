package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const (
    DAY_1 = 1
    DAY_2 = 2
    DAY_3 = 3
    DAY_4 = 4
    DAY_5 = 5
    DAY_6 = 6
    DAY_7 = 7
    DAY_8 = 8
    DAY_9 = 9
    DAY_10 = 10
    DAY_11 = 11
    DAY_12 = 12
    DAY_13 = 13
    DAY_14 = 14
    DAY_15 = 15
    DAY_16 = 16
    DAY_17 = 17
    DAY_18 = 18
    DAY_19 = 19
    DAY_20 = 20
    DAY_21 = 21
    DAY_22 = 22
    DAY_23 = 23
    DAY_24 = 24
    DAY_25 = 25
)

type AnswerFunc func([]string) string

type AnswerRunner struct {
	day int
}

func (runner *AnswerRunner) SetDay(newDay int) {
	runner.day = newDay
}

func (runner AnswerRunner) getInputs(path string) []string {
	path, pathErr := filepath.Abs(path)

	if pathErr != nil {
		panic(pathErr)
	}

	file, fileErr := os.Open(path)
	if fileErr != nil {
		panic(fileErr)
	}
	defer file.Close()

	var lines []string
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (runner AnswerRunner) RunAnswerFunc(answerFunc AnswerFunc) {
	inputs := runner.getInputs(fmt.Sprintf("./day%d/input.txt", runner.day))
	out := answerFunc(inputs)
	fmt.Println(out)
}

func (runner AnswerRunner) TestAnswerFunc(answerFunc AnswerFunc, expected string) {
	inputs := runner.getInputs(fmt.Sprintf("./day%d/test.txt", runner.day))
	out := answerFunc(inputs)
	fmt.Printf("Expected: %s\n", expected)
	fmt.Printf("Actual: %s\n", out)
	if out != expected {
		panic(fmt.Errorf("Expected and actual did not match!! Expected: %s | Actual: %s", expected, out))
	}
}
