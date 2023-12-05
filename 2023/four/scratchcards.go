package four

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
    "advent/helpers"
)

type (
	Scratchcard struct {
		id             int
		cardNumbers    []int
		winningNumbers []int
	}
)

func RunPartOne(input []string) {
	fmt.Println(ScratchcardsPointsSum(input))
}

func RunPartTwo(input []string) {
    fmt.Println(AllScratchcardCopies(input))
}

func ScratchcardsPointsSum(input []string) (sum int) {
	cards := GetAllScratchcards(input)
	for i := range cards {
		sum += GetScratchcardPoints(cards[i])
	}
	return sum
}

func AllScratchcardCopies(input []string) (sum int) {
	cards := GetAllScratchcards(input)
    counts := make([]int, len(input))

    fmt.Println(counts)

    for _, card := range cards {
        AddScratchcardCopies(card, cards, &counts)
    }

    for _, count := range counts {
        sum += count+1
    }

    return sum
}

func AddScratchcardCopies(card Scratchcard, allCards []Scratchcard, copiesList *[]int) {
    winningNumbers := GetScratchcardWinningNumbers(card)
    for i := card.id+1; i < card.id+1+winningNumbers; i++ {
        (*copiesList)[i-1]++
        AddScratchcardCopies(allCards[i-1], allCards, copiesList)
    }
}

func GetAllScratchcards(input []string) []Scratchcard {
	cards := make([]Scratchcard, len(input))

	for i := range input {
		cards[i] = GetScratchcard(input[i])
	}

	return cards
}

func GetScratchcardWinningNumbers(card Scratchcard) (matches int) {

	for _, num := range card.cardNumbers {
		if slices.Contains(card.winningNumbers, num) {
			matches++
		}
	}

	if matches == 0 {
		return matches
	}

	return matches
}

func GetScratchcardPoints(card Scratchcard) int {
    matches := GetScratchcardWinningNumbers(card)
	return int(math.Pow(2, float64(matches-1)))
}

func GetScratchcard(input string) (card Scratchcard) {
	card.id = ExtractCardId(input)
	card.cardNumbers, card.winningNumbers = GetNumberLists(input)
	return card
}

func ExtractCardId(input string) int {
    gameTitleParts := strings.Split(strings.Split(input, ":")[0], " ")
	id, err := strconv.Atoi(gameTitleParts[len(gameTitleParts)-1])

    if err != nil {
        fmt.Println(err)
		return -1
	}
	return id
}

func GetNumberLists(input string) (cardNumbers []int, winningNumbers []int) {
	listsString := strings.Split(input, ":")[1]
	lists := strings.Split(listsString, "|")

	cardNumbers = helpers.IntListStringToSlice(lists[0])
	winningNumbers = helpers.IntListStringToSlice(lists[1])

	return cardNumbers, winningNumbers
}
