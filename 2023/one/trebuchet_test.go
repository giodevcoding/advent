package one

import (
	"fmt"
	"testing"
    "advent/helpers"
)

func TestCalibrationValues(t *testing.T) {
	var sampleInput = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	expected := 142
	result := CalibrationValues(sampleInput)

	if result != expected {
		t.Errorf("CalibrationValues(string[]) = %d; expected %d", result, expected)
	}
}

func TestGetCalibrationValue(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s=%d", test.input, test.expected)
		t.Run(testname, func(t *testing.T) {
			result := GetCalibrationValue(test.input)
			if result != test.expected {
				t.Errorf("GetCalibrationValue(%s) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func TestProperCalibrationValues(t *testing.T) {
	var sampleInput = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	expected := 281
	result := ProperCalibrationValues(sampleInput)

	if result != expected {
		t.Errorf("ProperCalibrationValues(string[]) = %d; expected %d", result, expected)
	}
}

func TestGetProperCalibrationValue(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s=%d", test.input, test.expected)
		t.Run(testname, func(t *testing.T) {
			result := GetProperCalibrationValue(test.input)
			if result != test.expected {
				t.Errorf("GetCalibrationValue(%s) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func TestNumbersInString(t *testing.T) {
	var tests = []struct {
		input    string
		expected []int
	}{
		{"pqr3stu8vwx", []int{3, 8}},
		{"a1b2c3d4e5f", []int{1, 2, 3, 4, 5}},
		{"treb7uchet", []int{7}},
		{"two1nine", []int{2, 1, 9}},
		{"eightwothree", []int{8, 2, 3}},
		{"abcone2threexyz", []int{1, 2, 3}},
		{"xtwone3four", []int{2, 1, 3, 4}},
		{"4nineeightseven2", []int{4, 9, 8, 7, 2}},
		{"zoneight234", []int{1, 8, 2, 3, 4}},
		{"7pqrstsixteen", []int{7, 6}},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s=%v", test.input, test.expected)
		t.Run(testname, func(t *testing.T) {
			result := NumbersInString(test.input)
			if !helpers.IntSlicesEqual(result, test.expected) {
				t.Errorf("GetCalibrationValue(%s) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func TestGetDigit(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"zero", 0},
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s=%d", test.input, test.expected)
		t.Run(testname, func(t *testing.T) {
			result := GetDigit(test.input)
			if result != test.expected {
				t.Errorf("GetDigit(%s) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}
