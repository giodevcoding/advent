package four

import (
	"fmt"
	"strconv"
	"strings"
)

type (
    Scratchcard struct {
        id int
        cardNumbers []int
        winningNumbers []int
    }
)

func ScratchcardsPointsSum(input []string) (sum int) {
    cards := GetAllScratchCards(input)
    for i := range cards {
        fmt.Println(cards[i])
    }
    return sum
}

func GetAllScratchCards(input []string) []Scratchcard {
    cards := make([]Scratchcard, len(input))

    for i := range input {
        cards[i] = GetScratchcard(input[i])
    }

    return cards
}

func GetScratchcard(input string) (card Scratchcard){
    card.id = ExtractCardId(input)
    card.cardNumbers, card.winningNumbers = GetNumberLists(input)
    return card
}

func ExtractCardId(input string) int {
    id, err := strconv.Atoi(strings.Split(strings.Split(input, ":")[0], " ")[1])
    if (err != nil) {
        return -1
    }
    return id;
}

func IntListStringToSlice(numList string) []int{
    textNumbers := strings.Split(strings.TrimSpace(numList), " ")
    actualNumbers := make([]int, len(textNumbers))

    for i, textNum := range textNumbers {
        actualNum, err := strconv.Atoi(textNum)
        if (err == nil) {
            actualNumbers[i] = actualNum
        }
    }

    return actualNumbers
}

func GetNumberLists(input string) (cardNumbers []int, winningNumbers []int) {
    listsString := strings.Split(input, ":")[1]
    lists := strings.Split(listsString, "|")

    cardNumbers = IntListStringToSlice(lists[0])
    winningNumbers = IntListStringToSlice(lists[1])

    return cardNumbers, winningNumbers
}
