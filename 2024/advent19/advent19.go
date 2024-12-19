package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"regexp"
	"strings"
	"time"
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
	patterns []string
	designs  []string
}

func run1(input string) int {
	defer tools.Track(time.Now(), "1")
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
	defer tools.Track(time.Now(), "2")
	onsen := parseOnsen(input)
	cache := make(map[string]int)
	count := 0
	for _, design := range onsen.designs {
		ways := countWays(design, onsen.patterns, cache)
		count += ways
	}
	return count
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

func countWays(design string, patterns []string, cache map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	if count, ok := cache[design]; ok {
		return count
	}

	// could probably find a way to not bother with patterns known not to work, but this is fast enough

	total := 0
	for _, pattern := range patterns {
		if len(pattern) > len(design) || design[:len(pattern)] != pattern {
			continue
		}
		total += countWays(design[len(pattern):], patterns, cache)
	}
	cache[design] = total
	return total
}
