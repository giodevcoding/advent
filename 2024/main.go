package main

import (
	"aoc2024/day8"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_8, day8.UniqueAntinodes, "34"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
