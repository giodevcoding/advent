package four

import (
	"fmt"
	"strings"
	"testing"
)

func TestScratchcardsPointsSum(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		},
		13,
	}

	result := ScratchcardsPointsSum(test.input)
	if result != test.expected {
		t.Errorf("ScratchcardsPointsSum(%T) = %d; expected %d", test.input, result, test.expected)
	}

}

func TestAllScratchcardCopies(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		},
		30,
	}

	result := AllScratchcardCopies(test.input)
	if result != test.expected {
		t.Errorf("ScratchcardsPointsSum(%T) = %d; expected %d", test.input, result, test.expected)
	}

}

func TestGetScratchcardsPoints(t *testing.T) {
	tests := [8]struct {
		input    string
		expected int
	}{
		{"Card   1: 58  6 71 93 96 38 25 29 17  8 | 79 33 93 58 53 96 71  8 67 90 17  6 46 85 64 25 73 32 18 52 77 16 63  2 38", 64},
		{"Card   2: 34 79 17 22 15 73 61 58 46 32 | 36 53 22 32 56 15 71  7 17 19 79 81 44 59 46 34 52 61 24 73 54 28 88 50 58", 256},
		{"Card 171:  6 12 87 66 97 65 74 75  9 88 |  5 26 49 14 34 81 64 46 55 72 39 30 98 82 76 31 23 95 96 68 54 56 86 15 69", 0},
		{"Card 172: 54 13 27 98  1  9 53 71 17 11 | 92 62 55 54  1 17 24 40 39 98 11  5 80 75 71 99 53 20 50 81  9 89 27 13 86", 512},
		{"Card 119: 72 99 25 49 16 68  6 89 31 81 | 80 97 11 62 41 96 27 89 49  3 25  6 99 94 26 34 55 81 48 75 53 72 68 16 28", 256},
		{"Card 120: 78 77 54 70  9  6 22 43 40 16 | 48 20 54 25 75 91 33 67  3  5 95 37 29 90  1 24 32 39 12 51 71 16 36 76 63", 2},
		{"Card 121: 76 19 51 52  4 18 32 43 34 55 |  4 52 51  1 32 37 80  6 76 73 88 43 82 19 89 34 55 18 25 58 85 23  9  5 60", 512},
		{"Card  88: 44  8 41 84 38 91 70 31  1 50 | 60 83  4 51 49 88 90 14 34 45 70 25 56 23 91 11 38 41 48  7  2 19 28  9 27", 8},
	}

	for _, test := range tests {
		testName := strings.Split(test.input, ":")[0]
		t.Run(testName, func(t *testing.T) {
			card := GetScratchcard(test.input)
			fmt.Println(card)
			result := GetScratchcardPoints(card)
			if result != test.expected {
				t.Errorf("GetScratchcardPoints(%v) = %d; expected %d", card, result, test.expected)
			}
		})
	}
}
