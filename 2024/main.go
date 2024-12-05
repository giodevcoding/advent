package main

import (
	"aoc2024/day5"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_5, day5.PrintUpdateFixOrdering, "123"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
