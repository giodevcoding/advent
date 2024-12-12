package main

import (
	"aoc2024/day9"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_9, day9.UpdateChecksumWholeFiles, "2858"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
