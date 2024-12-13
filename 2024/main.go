package main

import (
	"aoc2024/day12"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_12, day12.CalculateFenceCost, "1206"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
