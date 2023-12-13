package eleven

import "testing"

func TestShortestGalaxyPathsSum(t *testing.T) {
	test := struct {
		input    []string
		expected int
	}{

		[]string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		},
		374,
	}

	result := ShortestGalaxyPathsSum(test.input, 2)
	if result != test.expected {
		t.Errorf("ShortestGalaxyPathsSum(%T) = %d; expected %d", test.input, result, test.expected)
	}
}

/*
....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......
*/
