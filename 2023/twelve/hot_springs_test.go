package twelve

import "testing"

func TestPossibleArrangements(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"???.### 1,1,3",
			".??..??...?##. 1,1,3",
			"?#?#?#?#?#?#?#? 1,3,1,6",
			"????.#...#... 4,1,1",
			"????.######..#####. 1,6,5",
			"?###???????? 3,2,1",
		},
		21,
	}

	result := PossibleArrangements(test.input)
	if result != test.expected {
		t.Errorf("PossibleArrangements(%T) = %d; expected %d", test.input, result, test.expected)
	}

}
