package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type hand struct {
	cardstr string
	bid     int
	score   int
}

var cardValues = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

const (
	HIGH = iota
	ONE_PAIR
	TWO_PAIR
	THREE
	FULL
	FOUR
	FIVE
)

func run1(input string) int {
	score := 0
	hands := parseHands(parsers.SplitByLines(input), false)
	// sort in inverse order
	sort.Slice(hands, func(i, j int) bool { return hands[i].score < hands[j].score })
	for i, h := range hands {
		score += (i + 1) * h.bid
	}
	return score
}

func run2(input string) int {
	score := 0
	hands := parseHands(parsers.SplitByLines(input), true)
	// sort in inverse order
	sort.Slice(hands, func(i, j int) bool { return hands[i].score < hands[j].score })
	for i, h := range hands {
		score += (i + 1) * h.bid
	}
	return score
}

func parseHands(lines []string, withJokers bool) []hand {

	hands := []hand{}
	// a line looks like "Q4QKK 465"
	for _, line := range lines {
		cardstr, bidstr, found := strings.Cut(line, " ")
		if !found {
			continue
		}
		h := hand{cardstr: cardstr, bid: tools.Atoi(bidstr)}
		if withJokers {
			h.score = scoreHand(h, handTypeWithJoker)
		} else {
			h.score = scoreHand(h, handType)
		}
		hands = append(hands, h)
	}

	return hands
}

func scoreHand(h hand, f func(h hand) int) int {
	// 6 digit hex number
	// first is type, 0-6
	// second is card1 2-E, third card2 2-E, etc
	// could do this with shifting but strings seem easier to visualize and debug

	ht := f(h)

	hexstr := fmt.Sprintf("%X%X%X%X%X%X", ht,
		cardValues[rune(h.cardstr[0])],
		cardValues[rune(h.cardstr[1])],
		cardValues[rune(h.cardstr[2])],
		cardValues[rune(h.cardstr[3])],
		cardValues[rune(h.cardstr[4])])

	score := tools.Must(strconv.ParseInt(hexstr, 16, 0))
	return int(score)
}

// returns one of the card type constants
func handType(h hand) int {

	bestSoFar := HIGH
	for cv, _ := range cardValues {
		count := strings.Count(h.cardstr, string(cv))
		if count == 5 {
			return FIVE
		} else if count == 4 {
			return FOUR
		} else if count == 3 {
			if bestSoFar == ONE_PAIR {
				return FULL
			}
			bestSoFar = THREE
		} else if count == 2 {
			if bestSoFar == THREE {
				return FULL
			}
			if bestSoFar == ONE_PAIR {
				return TWO_PAIR
			}
			bestSoFar = ONE_PAIR
		}
	}

	return bestSoFar
}

func handTypeWithJoker(h hand) int {
	h2 := hand{cardstr: strings.ReplaceAll(h.cardstr, "J", "")}
	ht := handType(h2)
	jokers := strings.Count(h.cardstr, "J")
	if jokers == 0 {
		return ht
	}
	if jokers == 5 {
		return FIVE
	}
	switch ht {
	case FIVE:
		return FIVE
	case FOUR:
		return FIVE
	case THREE:
		if jokers > 1 {
			return FIVE
		}
		return FOUR
	case FULL:
		return FULL
	case TWO_PAIR:
		return FULL
	case ONE_PAIR:
		if jokers >= 3 {
			return FIVE
		}
		if jokers == 2 {
			return FOUR
		}
		return THREE
	case HIGH:
		if jokers == 4 {
			return FIVE
		}
		if jokers == 3 {
			return FOUR
		}
		if jokers == 2 {
			return THREE
		}
		return ONE_PAIR
	default:
		return ht
	}

}
