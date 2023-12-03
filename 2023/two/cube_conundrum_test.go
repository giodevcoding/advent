package two

import (
	"fmt"
	"testing"
)

func TestPossibleGames(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		},
		8,
	}

	result := PossibleGames(test.input)
	if result != test.expected {
		t.Errorf("PossibleGames got %d; expected %d", result, test.expected)
	}
}

func TestParseGame(t *testing.T) {
	type Test struct {
		input    string
		expected CubeGame
	}
	tests := []Test{
		{
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: CubeGame{
				id: 1,
				rounds: []Round{
					{red: 4, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
			},
		},
		{
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: CubeGame{
				id: 2,
				rounds: []Round{
					{red: 0, green: 2, blue: 1},
					{red: 1, green: 3, blue: 4},
					{red: 0, green: 1, blue: 1},
				},
			},
		},
		{
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: CubeGame{
				id: 3,
				rounds: []Round{
					{red: 20, green: 8, blue: 6},
					{red: 4, green: 13, blue: 5},
					{red: 1, green: 5, blue: 0},
				},
			},
		},
		{
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: CubeGame{
				id: 4,
				rounds: []Round{
					{red: 3, green: 1, blue: 6},
					{red: 6, green: 3, blue: 0},
					{red: 14, green: 3, blue: 15},
				},
			},
		},
		{
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: CubeGame{
				id: 5,
				rounds: []Round{
					{red: 6, green: 3, blue: 1},
					{red: 1, green: 2, blue: 2},
				},
			},
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Game %d", test.expected.id)
		t.Run(testName, func(t *testing.T) {
			result := ParseGame(test.input)
			if !result.Equals(test.expected) {
				t.Errorf("ParseGame(%s) = %v; expected %v", test.input, result, test.expected)
			}
		})
	}
}

func TestCubeGameEquals(t *testing.T) {
	type (
		TestInput struct {
			gameA CubeGame
			gameB CubeGame
		}
		Test struct {
			input    TestInput
			expected bool
		}
	)

	tests := []Test{
		{
			input: TestInput{
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 4, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 4, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
			},
			expected: true,
		},
		{
			input: TestInput{
				CubeGame{
					id: 2,
					rounds: []Round{
						{red: 4, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 4, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
			},
			expected: false,
		},
		{
			input: TestInput{
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 3, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 4, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
			},
			expected: false,
		},
		{
			input: TestInput{
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 3, green: 0, blue: 3},
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
				CubeGame{
					id: 1,
					rounds: []Round{
						{red: 1, green: 2, blue: 6},
						{red: 0, green: 2, blue: 0},
					},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v == %v", test.input.gameA, test.input.gameB)
		t.Run(testName, func(t *testing.T) {
			result := test.input.gameA.Equals(test.input.gameB)
			if result != test.expected {
				t.Errorf("%v.Equals(%v) = %t; expected %t", test.input.gameA, test.input.gameB, result, test.expected)
			}
		})
	}
}

func TestRoundSliceEquals(t *testing.T) {
	type (
		TestInput struct {
			sliceA, sliceB RoundSlice
		}
		Test struct {
			input    TestInput
			expected bool
		}
	)

	tests := []Test{
		{
			TestInput{
				sliceA: []Round{
					{red: 4, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
				sliceB: []Round{
					{red: 4, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
			},
			true,
		},
		{
			TestInput{
				sliceA: []Round{
					{red: 3, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
				sliceB: []Round{
					{red: 4, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
			},
			false,
		},
		{
			TestInput{
				sliceA: []Round{
					{red: 4, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
				sliceB: []Round{
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
			},
			false,
		},
		{
			TestInput{
				sliceA: []Round{
					{red: 4, green: 0, blue: 3},
					{red: 1, green: 2, blue: 6},
					{red: 0, green: 2, blue: 0},
				},
				sliceB: []Round{
					{red: 1, green: 2, blue: 6},
					{red: 4, green: 0, blue: 3},
					{red: 0, green: 2, blue: 0},
				},
			},
			false,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v == %v", test.input.sliceA, test.input.sliceB)
		t.Run(testName, func(t *testing.T) {
			result := test.input.sliceA.Equals(test.input.sliceB)
			if result != test.expected {
				t.Errorf("%v.Equals(%v) = %t; expected %t", test.input.sliceA, test.input.sliceB, result, test.expected)
			}
		})
	}
}

func TestExtractGameId(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 3},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 4},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
		{"Game 59: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 59},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("ExtractGameId(\"%s\") == %d", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result := ExtractGameId(test.input)
			if result != test.expected {
				t.Errorf("ExtractGameId(\"%s\") = %d; expected %d", test.input, result, test.expected)
			}
		})

	}
}

func TestExtractRounds(t *testing.T) {
	tests := []struct {
		input    string
		expected RoundSlice
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", RoundSlice{
			{red: 4, green: 0, blue: 3},
			{red: 1, green: 2, blue: 6},
			{red: 0, green: 2, blue: 0},
		}},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", RoundSlice{
			{red: 0, green: 2, blue: 1},
			{red: 1, green: 3, blue: 4},
			{red: 0, green: 1, blue: 1},
		}},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", RoundSlice{
			{red: 20, green: 8, blue: 6},
			{red: 4, green: 13, blue: 5},
			{red: 1, green: 5, blue: 0},
		}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("ExtractRound(\"%s\") == %v", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result := ExtractRounds(test.input)
			if !result.Equals(test.expected) {
				t.Errorf("ExtractRounds(\"%s\") = %v; expected %v", test.input, result, test.expected)
			}
		})

	}
}

func TestParseRound(t *testing.T) {
	tests := []struct {
		input    string
		expected Round
	}{
		{"3 blue, 4 red", Round{4, 0, 3}},
		{"3 green, 4 blue, 1 red", Round{1, 3, 4}},
		{" 8 green, 6 blue, 20 red", Round{20, 8, 6}},
		{"3 green, 15 blue, 14 red", Round{14, 3, 15}},
		{"2 blue, 1 red, 2 green", Round{1, 2, 2}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("ParseRound(\"%s\") == %v", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result := ParseRound(test.input)
			if result != test.expected {
				t.Errorf("ParseRound(\"%s\") = %v; expected %v", test.input, result, test.expected)
			}
		})

	}
}

func TestIsRoundValid(t *testing.T) {
	maxCount := Round{12, 13, 14}
	tests := []struct {
		input    Round
		expected bool
	}{
		{Round{4, 0, 3}, true},
		{Round{1, 3, 4}, true},
		{Round{20, 8, 6}, false},
		{Round{14, 3, 15}, false},
		{Round{1, 2, 2}, true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("IsRoundValid(\"%v\") == %t", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result := IsRoundValid(test.input, maxCount)
			if result != test.expected {
				t.Errorf("IsRoundValid(\"%v\") = %t; expected %t", test.input, result, test.expected)
			}
		})

	}
}

func TestCalculateMinimumPowerSet(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			48,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			12,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			1560,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			630,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			36,
		},
	}

    for _, test := range tests {
        game := ParseGame(test.input)
        testName := fmt.Sprintf("Game %d", game.id)
        t.Run(testName, func(t *testing.T){
            result := CalculateMinimumPowerSet(game)
            if (result != test.expected) {
                t.Errorf("CalculateMinimumPowerSet(%s) = %d; expected %d", test.input, result, test.expected)
            }
        })
    }
}
