package main

import (
	"aoc2024/day16"
	. "aoc2024/utils"
)

func main() {
    day, answerFunc, testExpected := DAY_16, day16.ReindeerMazeLowestScore, "11048"

	runner := AnswerRunner{}
	runner.SetDay(day)
	runner.TestAnswerFunc(answerFunc, testExpected)
	runner.RunAnswerFunc(answerFunc)
}
