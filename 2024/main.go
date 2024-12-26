package main

import (
	"aoc2024/day17"
	. "aoc2024/utils"
)

func main() {
    day, answerFunc, _ := DAY_17, day17.ChronospatialComputer, "4,6,3,5,6,3,5,2,1,0"

	runner := AnswerRunner{}
	runner.SetDay(day)
	//runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
