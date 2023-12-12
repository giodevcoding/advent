package six

import "testing"

func TestRecordBeatingPossibilities(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		input: []string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		},
		expected: 288,
	}

	result := RecordBeatingPossibilities(test.input)
	if result != test.expected {
        t.Errorf("RecordBeatingPossibilities(%v) = %d; expected %d", test.input, result, test.expected)
	}
}

func TestSingleRecordBeatingPossibilities(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		input: []string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		},
		expected: 71503,
	}

	result := SingleRecordBeatingPossibilities(test.input)
	if result != test.expected {
        t.Errorf("RecordBeatingPossibilities(%v) = %d; expected %d", test.input, result, test.expected)
	}
}

/*
1 = 0 
      1
2 = 1
      3
3 = 4
      6
4 = 10
      10
5 = 20
      15
6 = 35
      21
7 = 56
      28
8 = 84
*/
