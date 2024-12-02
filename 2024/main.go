package main

import (
	"aoc2024/day2"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_2, day2.ProblemDampener, "4"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
