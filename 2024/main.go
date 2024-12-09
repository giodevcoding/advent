package main

import (
	"aoc2024/day7"
	. "aoc2024/utils"
)

func main() {
	var day, answerFunc, testExpected = DAY_7, day7.BridgeRepairOperators, "11387"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
