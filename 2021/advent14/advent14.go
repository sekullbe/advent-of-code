package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"math"
	"regexp"
	"strings"
)

//go:embed input.txt
var inputText string

type elementPair struct {
	e1 string
	e2 string
}
type rules map[elementPair]string
type pairCount map[elementPair]int

func main() {
	fmt.Printf("Magic number: %d\n", run(inputText, 10)) // 2892 too high
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run(inputText, 40))
}

// don't need a run1 and run2; run2 is the same thing with more steps
func run(inputText string, steps int) int {
	template, theRules := parseInput(inputText)
	pairCounts := polymerize(template, theRules, steps)

	// count the elts in the pairs, and get the most and least common
	eltCounts := countElementsInPairs(pairCounts, string(template[0]), string(template[len(template)-1]))
	maxCount := 0
	minCount := math.MaxInt
	for _, c := range eltCounts {
		if c > maxCount {
			maxCount = c
		}
		if c < minCount {
			minCount = c
		}
	}
	return maxCount - minCount
}

func countElementsInPairs(pc pairCount, first, last string) map[string]int {
	ec := make(map[string]int)
	for pair, c := range pc {
		ec[pair.e1] += c
		ec[pair.e2] += c
	}
	// adjust for first and last because they were each only counted once
	ec[first]++
	ec[last]++

	// because they're pairs, we've seen them all twice
	for e, c := range ec {
		ec[e] = c / 2
	}
	return ec
}

/*
break the input down into pairs
i.e. NN NC CB
then when you add an element, remove the old pairs and add new ones
eg
NN -> NCN = NC CN
NC -> NBC = NC BC
CB -> CHB = CH HB

from example
NCNBCHB
NC CN NB BC CH HB
they are the same
*/

func initializePairs(template string) pairCount {
	count := make(pairCount)
	// read each elementPair of runes in the string and add it to the paircount
	reader := strings.NewReader(template)
	e1, _, _ := reader.ReadRune()
	for {
		e2, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		count[elementPair{string(e1), string(e2)}]++
		e1 = e2
	}
	return count
}

func polymerize(template string, rules rules, steps int) pairCount {
	count := initializePairs(template)
	for i := 0; i < steps; i++ {
		newCount := copyCount(count)
		for pair, newElt := range rules {
			// eg NN -> NBN - there are no more NNs but there are now an equal number of NB and BN pairs
			p := count[pair]
			if p > 0 { // if we have any matching pairs, add the split pairs and remove the originals, i.e. AB -> AC CB
				newCount[pair] -= p
				newCount[elementPair{pair.e1, newElt}] += p
				newCount[elementPair{newElt, pair.e2}] += p
			}
		}
		// clean up any zero pairs- makes the tests a little easier to write
		for pair, c := range newCount {
			if c == 0 {
				delete(newCount, pair)
			}
		}
		count = newCount
	}
	return count
}

func copyCount(originalMap pairCount) pairCount {
	newMap := make(pairCount)
	for k, v := range originalMap {
		newMap[k] = v
	}
	return newMap
}

func parseInput(inputText string) (template string, theRules rules) {
	// input is template, blank line, rules
	lines := parsers.SplitByLines(inputText)
	template = lines[0]
	theRules = parseRules(lines[2:])
	return template, theRules
}

func parseRules(lines []string) rules {
	theRules := make(rules)
	re := regexp.MustCompile(`([A-Z])([A-Z]) -> ([A-Z])`)
	for i := 0; i < len(lines); i++ {
		matches := re.FindStringSubmatch(lines[i])
		theRules[elementPair{matches[1], matches[2]}] = matches[3]
	}
	return theRules
}
