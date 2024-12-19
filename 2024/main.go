package main

import (
	"aoc2024/day15"
	. "aoc2024/utils"
)

func main() {
    day, answerFunc, _ := DAY_15, day15.WarehouseWideRobot, "9021"

	runner := AnswerRunner{}
	runner.SetDay(day)
	//runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
