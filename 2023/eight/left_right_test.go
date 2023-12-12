package eight

import "testing"

func TestLeftRightMinSteps(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			2,
		},
		{
			[]string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			6,
		},
	}

	for _, test := range tests {
		testName := test.input[0]
		t.Run(testName, func(t *testing.T) {
			result := LeftRightMinSteps(test.input)
			if result != test.expected {
				t.Errorf("LeftRightMinSteps(%s) = %d; expected %d.", testName, result, test.expected)
			}
		})
	}
}

func TestLeftRightMinStepsGhost(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		},
		6,
	}

    result := LeftRightMinStepsGhost(test.input)
    if (result != test.expected) {
        t.Errorf("LeftRightMinStepsGhost(%T) = %d; expected %d", test.input, result, test.expected)
    }
}
