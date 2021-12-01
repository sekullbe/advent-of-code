package advent6

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputs string

func Run1() {

	groupResponses := parsers.SplitByEmptyNewline(inputs)
	sumResponses := 0
	for _, respons := range groupResponses {
		seenResponses := make(map[rune]bool)
		for _, char := range respons {
			if char >= 'a' && char <= 'z' {
				seenResponses[char]=true
			}
		}
		distinctResponses := len(seenResponses)
		sumResponses += distinctResponses
	}

	fmt.Printf("Sum of responses: %d\n",sumResponses)
}

func Run2() {
	sumResponses := run2_doit(inputs)
	fmt.Printf("Sum of full-group responses: %d\n",sumResponses)
}

func run2_doit(inp string) int {
	groupResponses := parsers.SplitByEmptyNewline(inp)
	sumResponses := 0
	for _, respons := range groupResponses {
		matches := countMatchesInString(respons)
		sumResponses += matches
		fmt.Printf("Matches: %d Running Total: %d\n", matches, sumResponses)
	}
	//3123 is too low
	return sumResponses
}



// given a multiline string, counts how many times the same letter appears on every line
func countMatchesInString(respons string) int {
	everyoneAnsweredCount := 0
	respons = strings.TrimRight(respons, "\n") // the trailing newline on last group was making inputs off by one
	peopleInGroup := strings.Count(respons, "\n") + 1
	seenResponses := make(map[rune]int)
	for _, char := range respons {
		if char >= 'a' && char <= 'z' {
			seenResponses[char]++
			if seenResponses[char] == peopleInGroup {
				everyoneAnsweredCount++
			}
		}
	}
	return everyoneAnsweredCount
}


// break up input into groups
// count all individual letters in the group (OR them together)
// should this be a map or a set?
