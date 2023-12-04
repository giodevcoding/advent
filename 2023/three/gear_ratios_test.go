package three

import (
	"advent/helpers"
	"fmt"
	"testing"
)

func TestEngineSchematicSum(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
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
			expected: 4361,
		},
		{
			input: []string{
				".479........155..............944.....622..............31.........264.......................532..........................254.........528.....",
				"..............-...............%.....+...................=....111*.................495.......+.......558..................../..........*.....",
				"....................791*..62.....$.............847........&........-..........618.*...........818....&..642.........................789.....",
				"....520.58......405......#....542.../587.............*....198.......846.........*..............*.......*....................647.............",
				".........*........./.964..........................474.302.....................786...43..............505..436...................*.....#51....",
			},
			expected: 13615,
		},
		{
			input: []string{
				"........@.........960..................................*...........966.321...925............926...................*.947..&.............574..",
				"..............%.....$.........$......*.......479....909.339..........*..............803........*17......284$...657.......587......*.........",
				"...........772............&....345..93...465*................419......676...............-.@521.....-...........................399.662......",
			},
			expected: 10057,
		},
		{
			input: []string{
                ".......+...38..........*...506.........811.....+188......766...623..363....*......*.914............#.@..92.365.........../...694..312..156..",
                "199...745......189=.389.....676......+........*..=........442....810........................................*.......................795.....",
                "........59...*.......405...*..........*......%..........&.........*.........515.586.......239@...571.80..................852...........*....",
                "....737.....608..........362...336....642....606..................262......................................209.........................617..",
                "566..186./......*277..18....*779..*.........+...........226..........184.....=......696..........+............344...550...........*.....65..",
                ".....*........................$20................*..............................901......388...........$.............25...........*......157",
                "..904........659............$.....762........808...%..........*18.............................#.878..........44..693............611..566*...",
			},
			expected: 11347,
		},
	}

	for i, test := range tests {
		testName := fmt.Sprint(i)
		t.Run(testName, func(t *testing.T) {
			result := EngineSchematicSum(test.input)
			if result != test.expected {
				t.Errorf("EngineSchematicSum(%T) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func TestGearRatioSum(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
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
			expected: 467835,
		},
    }

	for i, test := range tests {
		testName := fmt.Sprint(i)
		t.Run(testName, func(t *testing.T) {
			result := GearRatioSum(test.input)
			if result != test.expected {
				t.Errorf("GearRatioSum(%T) = %d; expected %d", test.input, result, test.expected)
			}
		})
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
