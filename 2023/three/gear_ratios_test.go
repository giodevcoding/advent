package three

import (
	"advent/helpers"
	"fmt"
	"testing"
)

func TestEngineSchematicSum(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
		4361,
	}

    result := EngineSchematicSum(test.input)
    if (result != test.expected) {
        t.Errorf("EngineSchematicSum(%T) = %d; expected %d", test.input, result, test.expected)
    }
}

func TestExtractSymbolIndexesFromLIne(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"467..114..", []int{}},
		{"...*......", []int{3}},
		{"..35..633.", []int{}},
		{"......#...", []int{6}},
		{"617*......", []int{3}},
		{".....+.58.", []int{5}},
		{"..592.....", []int{}},
		{"......755.", []int{}},
		{"...$.*....", []int{3, 5}},
		{".664.598..", []int{}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s = %v", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result := ExtractSymbolIndexesFromLine(test.input)
			if !helpers.IntSlicesEqual(result, test.expected) {
				t.Errorf("ExtractSymbolIndexesFromLine(%s) = %v; expected %v", test.input, result, test.expected)
			}
		})
	}
}

func TestExtractPartNumbersFromLine(t *testing.T) {
	tests := []struct {
		input    string
		expected PartNumberSlice
	}{
		{"467..114..", PartNumberSlice{
			{467, 0, 3},
			{114, 5, 3},
		}},
		{"...*......", PartNumberSlice{}},
		{"..35..633.", PartNumberSlice{
			{35, 2, 2},
			{633, 6, 3},
		}},
		{"......#...", PartNumberSlice{}},
		{"617*......", PartNumberSlice{
			{617, 0, 3},
		}},
		{".....+.58.", PartNumberSlice{
			{58, 7, 2},
		}},
		{"..592.....", PartNumberSlice{
			{592, 2, 3},
		}},
		{"......755.", PartNumberSlice{
			{755, 6, 3},
		}},
		{"...$.*....", PartNumberSlice{}},
		{".664.598..", PartNumberSlice{
			{664, 1, 3},
			{598, 5, 3},
		}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s = %v", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result := ExtractPartNumbersFromLine(test.input)
			fmt.Printf("result = %v; expected = %v\n", result, test.expected)
			if result.Equals(test.expected) {
				t.Errorf("ExtractSymbolIndexesFromLine(%s) = %v; expected %v", test.input, result, test.expected)
			}
		})
	}
}
