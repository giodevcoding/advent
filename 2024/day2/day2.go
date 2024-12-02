package day2

import (
	"aoc2024/utils"
	"math"
	"strconv"
)

func SafeReports(input []string) string {
	var count = 0
	for _, reportStr := range input {
		report := utils.StringToNumList(reportStr)

		if isReportSafe(report) {
			count++
		}

	}
	return strconv.Itoa(count)
}

func ProblemDampener(input []string) string {
	var count = 0
	for _, reportStr := range input {
		report := utils.StringToNumList(reportStr)

		if isReportSafe(report) {
			count++
		} else {
			for i := range report {
                reportDampened := removeReportLevel(report, i)
                if (isReportSafe(reportDampened)) {
                    count++
                    break
                }
			}
		}

	}
	return strconv.Itoa(count)
}

func removeReportLevel(report []int, levelIndex int) []int {
    result := make([]int, 0, len(report)-1)
    for i := range report {
        if (i != levelIndex) {
            result = append(result, report[i])
        }
    }
    return result
}

func isReportSafe(report []int) bool {
	dir := 0
	for i := 0; i < len(report)-1; i++ {
		current, next := report[i], report[i+1]

		diff := next - current
		absDiff := math.Abs(float64(next) - float64(current))

		if absDiff < 1 || absDiff > 3 {
			return false
		}

		currentDir := diff / int(absDiff)

		if dir == 0 {
			dir = currentDir
		} else if dir != currentDir {
			return false
		}

	}
	return true
}
