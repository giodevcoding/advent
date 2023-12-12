package seven

import (
	"fmt"
	"testing"
)

func TestTotalWinnings(t *testing.T) {
	type Input struct {
		lines []string
		joker bool
	}
	tests := []struct {
		input    Input
		expected int
	}{
		{
			input: Input{
				lines: []string{
					"32T3K 765",
					"T55J5 684",
					"KK677 28",
					"KTJJT 220",
					"QQQJA 483",
				},
				joker: false},
			expected: 6440,
		},
		{
			input: Input{
				lines: []string{
					"32T3K 765",
					"T55J5 684",
					"KK677 28",
					"KTJJT 220",
					"QQQJA 483",
				},
				joker: true},
			expected: 5905,
		},
    }

    for _, test := range tests {
        testName := fmt.Sprintf("TotalWinnings (joker = %t)", test.input.joker)
        t.Run(testName, func(t *testing.T) {
            result := TotalWinnings(test.input.lines, test.input.joker)
            if result != test.expected {
                t.Errorf("TotalWinnings(%T, %t) = %d; expected %d", test.input.lines, test.input.joker, result, test.expected)
            }
        })
    }
}
