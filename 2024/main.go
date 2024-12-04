package main

import (
	"aoc2024/day4"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_4, day4.XDashMasWordSearch, "9"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
