package main

import (
	"aoc2024/day3"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_3, day3.CorruptedMemoryRegexEnabling, "48"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
