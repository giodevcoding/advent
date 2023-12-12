package six

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	RaceRecord struct {
		time     int
		distance int
	}
)

func RunPartOne(input []string) {
	fmt.Println(RecordBeatingPossibilities(input))
}

func RunPartTwo(input []string) {
	fmt.Println(SingleRecordBeatingPossibilities(input))
}

func RecordBeatingPossibilities(input []string) (result int) {
    result = 1
	records := ParseRaceRecords(input)
	for _, record := range records {
		result *= GetRecordBeatingPossibilities(record)
	}
	return result
}

func SingleRecordBeatingPossibilities(input []string) int {
    return GetRecordBeatingPossibilities(ParseSingleRaceRecord(input))
}

func GetRecordBeatingPossibilities(record RaceRecord) int {
	for i := 0; i < record.time; i++ {
        travel := i * (record.time - i)

		if travel > record.distance {
			possibilities := (record.time - (i*2))+1
			return possibilities
		}
	}

	return 0
}

func ParseSingleRaceRecord(input []string) RaceRecord {
	timeSplit := strings.Split(strings.Split(input[0], ":")[1], " ")
	distanceSplit := strings.Split(strings.Split(input[1], ":")[1], " ")

    timeStrings := ""
    distanceStrings := ""

	for _, timeStr := range timeSplit {
		if len(timeStr) > 0 {
			timeStrings += timeStr
		}
	}

	for _, distanceStr := range distanceSplit {
		if len(distanceStr) > 0 {
			distanceStrings += distanceStr
		}
	}

    time, _ := strconv.Atoi(timeStrings)
    distance, _ := strconv.Atoi(distanceStrings)

    return RaceRecord{
        time: time,
        distance: distance,
    }
}

func ParseRaceRecords(input []string) (records []RaceRecord) {
	timeSplit := strings.Split(strings.Split(input[0], ":")[1], " ")
	distanceSplit := strings.Split(strings.Split(input[1], ":")[1], " ")

	times := []int{}
	distances := []int{}

	for _, timeStr := range timeSplit {
		if len(timeStr) > 0 {
			time, _ := strconv.Atoi(timeStr)
			times = append(times, time)
		}
	}

	for _, distanceStr := range distanceSplit {
		if len(distanceStr) > 0 {
			distance, _ := strconv.Atoi(distanceStr)
			distances = append(distances, distance)
		}
	}

	for i := range times {
		records = append(records, RaceRecord{times[i], distances[i]})
	}

	return records
}
