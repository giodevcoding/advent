package main

import (
	"aoc2024/day10"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_10, day10.TrailheadRatings, "81"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
