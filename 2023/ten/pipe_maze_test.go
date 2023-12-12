package ten

import (
	"fmt"
	"testing"
)

func TestFarthestDistance(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			8,
		},
		{

			[]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			}, 4,
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			result := FarthestDistance(test.input)
			if result != test.expected {
				t.Errorf("FarthestDistance(%T) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func TestEnclosedTiles(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			},
			1,
		},
		{

			[]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			}, 1,
		},
		{
			[]string{
				"...........",
				".S-------7.",
				".|F-----7|.",
				".||.....||.",
				".||.....||.",
				".|L-7.F-J|.",
				".|..|.|..|.",
				".L--J.L--J.",
				"..........."},
			4,
		},
		{

			[]string{
				"FF7FSF7F7F7F7F7F---7",
				"L|LJ||||||||||||F--J",
				"FL-7LJLJ||||||LJL-77",
				"F--JF--7||LJLJ7F7FJ-",
				"L---JF-JLJ.||-FJLJJ7",
				"|F|F-JF---7F7-L7L|7|",
				"|FFJF7L7F-JF7|JL---7",
				"7-L-JL7||F7|L7F-7F7|",
				"L.L7LFJ|||||FJL7||LJ",
				"L7JLJL-JLJLJL--JLJ.L",
			}, 10,
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			result := EnclosedTiles(test.input)
			if result != test.expected {
				t.Errorf("FarthestDistance(%T) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}
