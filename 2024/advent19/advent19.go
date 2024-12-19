package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"regexp"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

// input is one line of comma-sep strings, then blank line, then n lines of single strings

type onsen struct {
	patterns []string // or maybe [][]rune?
	designs  []string // or maybe [][]rune?
}

func run1(input string) int {

	// loop over all patterns
	// see if pattern matches the beginning of desired pattern
	// or... make a BIG-ASS REGEXP!!! HELL YEAH!!!
	//example: ^(?:r|wr|b|g|bwu|rb|gb|br)+$
	matches := 0
	onsen := parseOnsen(input)
	re := buildRegex(onsen.patterns)
	for _, design := range onsen.designs {
		if re.MatchString(design) {
			matches++
		}
	}

	return matches
}

func run2(input string) int {

	return 0
}

func parseOnsen(input string) onsen {
	segments := parsers.SplitByEmptyNewlineToSlices(input)

	onsen := onsen{}
	onsen.patterns = parsers.SplitByCommasAndTrim(segments[0][0])
	onsen.designs = segments[1]

	return onsen
}

func buildRegex(patterns []string) *regexp.Regexp {
	return regexp.MustCompile("^(?:" + strings.Join(patterns, "|") + ")+$")
}
