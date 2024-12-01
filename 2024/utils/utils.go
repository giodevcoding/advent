package utils
import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type AnswerFunc func([]string) string

type AnswerRunner struct {
    day int
}

func (runner *AnswerRunner) SetDay(newDay int) {
    runner.day = newDay
}

func (runner AnswerRunner) getInputs(path string) []string {
    path, pathErr :=filepath.Abs(path)

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
    for  scanner.Scan() {
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

func Filter[T any](arr []T, filterFunc func(T) bool) []T {
    result := []T{}
    for _, value := range arr {
        if filterFunc(value) {
            result = append(result, value)
        }
    }
    return result
}

func Map[T, U any](arr []T, mapFunc func(T) U) []U {
    result := make([]U, len(arr))
    for i, value := range arr {
        result[i] = mapFunc(value)
    }
    return result
}

func StringToNumList(str string) []int {
    split := strings.Split(str, " ")
    split = Filter(split, func(s string) bool {
        return len(strings.TrimSpace(s)) > 0
    })
    nums := Map(split, func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	})
    return nums
}
