package main

import (
	_ "embed"
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"regexp"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	score := 0
	for _, card := range parsers.SplitByLines(inputText) {
		if card != "" {
			score += evaluateOneCard(card)
		}
	}

	return score
}

func run2(inputText string) int {
	totalCards := 0
	copies := make(map[int]int)
	cardLines := parsers.SplitByLines(inputText)

	for _, cardLine := range cardLines {
		cardNum, winners, numbers, err := parseCard(cardLine)
		if err != nil {
			continue
		}
		copies[cardNum]++ // initial card
		winCount := countWinningNumbersInCard(winners, numbers)
		for i := cardNum + 1; i <= cardNum+winCount; i++ {
			copies[i] += copies[cardNum]
		}
	}
	for _, numCopies := range copies {
		totalCards += numCopies
	}

	return totalCards
}

func parseCard(card string) (gameNum int, winners []int, numbers []int, err error) {
	//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	re := regexp.MustCompile(`Card\s+(\d+):\s+(.+)\s+\|\s+(.*)`)
	matches := re.FindStringSubmatch(card)
	if len(matches) < 4 {
		return 0, nil, nil, errors.New("can't parse string: " + card)
	}
	gameNum = tools.Atoi(matches[1])
	winners = parsers.StringsToIntSlice(matches[2])
	numbers = parsers.StringsToIntSlice(matches[3])
	return
}

func countWinningNumbersInCard(winners, numbers []int) int {
	wset := mapset.NewSet[int](winners...)
	nset := mapset.NewSet[int](numbers...)
	winnersOnly := wset.Intersect(nset)
	return winnersOnly.Cardinality()
}

func evaluateOneCard(card string) (score int) {
	_, winningNums, gameNums, err := parseCard(card)
	if err != nil {
		return 0
	}
	// parse the winning numbers as keys into a map 'winners' with val -1
	// for each card number, if it's in the map, winners[num]++
	// then iterate the keys, if val > 0, score += 2^(winners[num])
	winners := make(map[int]int)
	for _, winner := range winningNums {
		winners[winner] = 0
	}
	for _, num := range gameNums {
		if _, ok := winners[num]; ok {
			winners[num]++
		}
	}
	for _, wins := range winners {
		if wins > 0 {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}
