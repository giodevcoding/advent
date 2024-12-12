package main

import (
	"aoc2024/day11"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_11, day11.MagicStonesBlinkingFast, "65601038650482"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
