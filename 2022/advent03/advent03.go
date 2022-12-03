package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	rucksacks := parsers.SplitByLines(inputText)
	var totalScore int
	for _, r := range rucksacks {
		c1, c2 := divideRucksack(r)
		item := findSharedItem(c1, c2)
		score := scoreItem(item)
		totalScore += score
	}
	return totalScore
}

func run2(inputText string) int {
	// we don't care about the divisions in the sacks
	// just load 3 elves and find the common shared item among all three
	rucksacks := parsers.SplitByLines(inputText)
	var totalScore int
	for i := 0; i < len(rucksacks); i += 3 {
		c1 := rucksacks[i]
		c2 := rucksacks[i+1]
		c3 := rucksacks[i+2]
		sharedItem := findSharedItemThree([]rune(c1), []rune(c2), []rune(c3))
		log.Printf("shared item %c", sharedItem)
		totalScore += scoreItem(sharedItem)
	}

	return totalScore
}

func divideRucksack(input string) (c1, c2 []rune) {

	len := len(input)
	for i, c := range input {
		if i < len/2 {
			c1 = append(c1, c)
		} else {
			c2 = append(c2, c)
		}
	}
	return
}

func findSharedItem(c1, c2 []rune) rune {
	// O(n^2), yeah, but I don't care
	for _, r := range c1 {
		if tools.ContainsRune(c2, r) {
			return r
		}
	}
	return 0
}

func findSharedItemThree(c1, c2, c3 []rune) rune {
	// O(n^2), yeah, but I don't care
	for _, r := range c1 {
		if tools.ContainsRune(c2, r) {
			if tools.ContainsRune(c3, r) {
				return r
			}
		}
	}
	return 0
}

func scoreItem(item rune) int {

	// a=97, A=65, so A-a = -32 but we want 27, so add 58 (in addition to the +1 to offset from zero)
	score := int(item-'a') + 1
	if score < 0 {
		score += 58
	}
	return score
}
