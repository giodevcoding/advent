package main

import (
	"aoc2024/day6"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_6, day6.GuardPossibleObstructions, "6"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
