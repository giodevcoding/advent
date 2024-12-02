package main

import (
	"aoc2024/day1"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_1, day1.ReconcileLists, "11"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
