package nine

import "testing"

func TestOasisReportExtrapolation(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		},
		114,
	}

	result := OasisReportExtrapolation(test.input)
	if result != test.expected {
		t.Errorf("OasisReportExtrapolation(%T) = %d; expected %d", test.input, result, test.expected)
	}
}

func TestOasisReportBackwardsExtrapolation(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{
		[]string{
			"0 3 6 9 12 15",
			"1 3 6 10 15 21",
			"10 13 16 21 30 45",
		},
		2,
	}

	result := OasisReportBackwardsExtrapolation(test.input)
	if result != test.expected {
		t.Errorf("OasisReportExtrapolation(%T) = %d; expected %d", test.input, result, test.expected)
	}
}
