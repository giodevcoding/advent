package main

import (
	"aoc2024/day13"
	. "aoc2024/utils"
)

func main() {
    var day, answerFunc, testExpected = DAY_13, day13.ClawMachineMatrixMath, "875318608908"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
