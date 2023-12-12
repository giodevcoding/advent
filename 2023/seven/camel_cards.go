package seven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type (
	Hand struct {
		cards string
		bid   int
	}
)

func RunPartOne(input []string) {
	fmt.Println(TotalWinnings(input, false))
}

func RunPartTwo(input []string) {
	fmt.Println(TotalWinnings(input, true))
}

func TotalWinnings(input []string, joker bool) (result int) {
	hands := GetHands(input)
	sort.SliceStable(hands, func(i, j int) bool {
		return HandComparator(hands[i], hands[j], joker)
	})

	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}

	return result
}

func HandComparator(handA Hand, handB Hand, joker bool) bool {
	var handARank, handBRank int
	handARank = GetHandRank(handA, joker)
	handBRank = GetHandRank(handB, joker)

	if handARank < handBRank {
		return true
	} else if handBRank < handARank {
		return false
	} else {
		return SameRankHandComparator(handA, handB, joker)
	}
}

func SameRankHandComparator(handA Hand, handB Hand, joker bool) bool {
	cardValues := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}

	if joker {
		cardValues['J'] = 0
	}

	for i := range handA.cards {
		cardAValue := cardValues[rune(handA.cards[i])]
		cardBValue := cardValues[rune(handB.cards[i])]

		if cardAValue < cardBValue {
			return true
		} else if cardBValue < cardAValue {
			return false
		}
	}

	return false
}

func GetHandRankJoker(hand Hand) int {

	return 0
}

func GetHandRank(hand Hand, joker bool) int {
	cardCounts := make(map[rune]int)

    jokers := 0

	for _, card := range hand.cards {
		if joker && card == 'J' {
            jokers++;
            continue
		}
		_, exists := cardCounts[card]

		if exists {
			cardCounts[card]++
		} else {
			cardCounts[card] = 1
		}
	}

    if joker {
        greatestCount := 0
        var mostFrequentCard rune
        for card, count := range cardCounts {
            if (count > greatestCount) {
                greatestCount = count
                mostFrequentCard = card
            }
        }

        cardCounts[mostFrequentCard] += jokers
    }

	// Five-of-a-kind
	if len(cardCounts) == 1 {
		return 7
	}

	if len(cardCounts) == 2 {
		for _, count := range cardCounts {
			// Four-of-a-kind
			if count == 1 || count == 4 {
				return 6
				// Full House
			} else {
				return 5
			}
		}
	}

	if len(cardCounts) == 3 {
		for _, count := range cardCounts {
			// Three-of-a-kind
			if count == 3 {
				return 4
			}
		}
		// Two Pair
		return 3
	}

	if len(cardCounts) == 4 {
		return 2
	}

	if len(cardCounts) == 5 {
		return 1
	}

	return 0
}

func GetHands(input []string) []Hand {
	hands := []Hand{}
	for i := range input {
		split := strings.Split(input[i], " ")
		cards := split[0]
		bid, _ := strconv.Atoi(split[1])
		hands = append(hands, Hand{cards, bid})
	}
	return hands
}
