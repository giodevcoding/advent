package two

import (
	"advent/helpers"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	CubeGame struct {
		id     int
		rounds RoundSlice
	}
	Round struct {
		red   int
		green int
		blue  int
	}
	RoundSlice []Round
)

func (game CubeGame) Equals(other CubeGame) bool {
	if game.id != other.id {
		return false
	}
	if !game.rounds.Equals(other.rounds) {
		return false
	}
	return true
}

func (roundSlice RoundSlice) Equals(otherSlice RoundSlice) bool {
	if len(roundSlice) != len(otherSlice) {
		return false
	}

	for i := range roundSlice {
		if roundSlice[i] != otherSlice[i] {
			return false
		}
	}

	return true
}

func RunPartOne() {
	var absPath, _ = filepath.Abs("./two/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(PossibleGames(lines))
}

func RunPartTwo() {
	var absPath, _ = filepath.Abs("./two/input.txt")
	var lines, err = helpers.ReadFileLines(absPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(AllMinimumPowerSets(lines))
}

func PossibleGames(input []string) int {
	maxCount := Round{12, 13, 14}
	games := ParseGamesList(input)
	idSums := 0

	for _, game := range games {
		possible := true

		for _, round := range game.rounds {
			if !IsRoundValid(round, maxCount) {
				possible = false
			}
		}

		if possible {
			idSums += game.id
		}
	}

	return idSums
}

func AllMinimumPowerSets(input []string) int {
	games := ParseGamesList(input)
	sumOfPowers := 0

	for _, game := range games {
		sumOfPowers += CalculateMinimumPowerSet(game)
	}

	return sumOfPowers
}

func IsRoundValid(round Round, maxCount Round) bool {
	if round.red > maxCount.red {
		return false
	}

	if round.green > maxCount.green {
		return false
	}
	if round.blue > maxCount.blue {
		return false
	}

	return true
}

func CalculateMinimumPowerSet(game CubeGame) int {
	minimums := Round{0, 0, 0}
	for _, round := range game.rounds {
		if round.red > minimums.red {
            minimums.red = round.red
		}

		if round.green > minimums.green {
            minimums.green = round.green
		}

		if round.blue > minimums.blue {
            minimums.blue = round.blue
		}
	}
    return minimums.red * minimums.green * minimums.blue
}

func ParseGamesList(inputs []string) []CubeGame {
	games := make([]CubeGame, len(inputs))
	for i := range inputs {
		games[i] = ParseGame(inputs[i])
	}

	return games
}

func ParseGame(input string) CubeGame {
	gameId := ExtractGameId(input)
	rounds := ExtractRounds(input)

	return CubeGame{gameId, rounds}
}

func ExtractGameId(input string) int {
	colonSplit := strings.Split(input, ":")
	gameId, err := strconv.Atoi(strings.Split(colonSplit[0], " ")[1])
	if err != nil {
		fmt.Println(fmt.Errorf("Could not extract GameId from string: %s", input))
		return -1
	}

	return gameId
}

func ExtractRounds(input string) RoundSlice {
	colonSplit := strings.Split(input, ":")
	roundStrings := strings.Split(colonSplit[1], ";")
	rounds := make(RoundSlice, len(roundStrings))

	for i, roundStr := range roundStrings {
		rounds[i] = ParseRound(roundStr)
	}
	return rounds
}

func ParseRound(input string) Round {
	diceCounts := strings.Split(input, ",")
	round := Round{0, 0, 0}

	for _, diceCount := range diceCounts {
		parts := strings.Split(strings.TrimSpace(diceCount), " ")
		count, _ := strconv.Atoi(parts[0])

		switch color := parts[1]; color {
		case "red":
			round.red = count
		case "green":
			round.green = count
		case "blue":
			round.blue = count
		}
	}
	return round
}
