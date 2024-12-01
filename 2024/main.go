package main

import (
	"aoc2024/day1"
	. "aoc2024/utils"
)

func main() {
	runner := AnswerRunner{}
	runner.SetDay(1)
	runner.TestAnswerFunc(day1.SimilarityScores, "31")
	runner.RunAnswerFunc(day1.SimilarityScores)
}
