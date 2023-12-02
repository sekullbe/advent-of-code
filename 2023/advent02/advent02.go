package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
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

type game struct {
	num, red, green, blue int
}

func run1(inputText string) int {

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	possibles := 0

	// keep a running max of the r/g/b we've seen in each game
	lines := parsers.SplitByLines(inputText)
	for gamenum, line := range lines {
		g := game{num: gamenum, red: 0, green: 0, blue: 0}
		// could micro-optimize and just see if any pull is > max and exit, but meh
		for _, pull := range strings.Split(line, ";") {
			g.red = max(g.red, extractnum("red", pull))
			g.blue = max(g.blue, extractnum("blue", pull))
			g.green = max(g.green, extractnum("green", pull))
		}
		if g.red > maxRed || g.blue > maxBlue || g.green > maxGreen {
			continue
		}
		possibles += gamenum + 1 // +1 because we're starting with zero

	}

	return possibles
}

func extractnum(color string, contentString string) int {
	re := regexp.MustCompile(`(\d+) ` + color)
	matches := re.FindStringSubmatch(contentString)
	if len(matches) == 0 {
		return 0

	}
	return tools.Atoi(matches[1])
}

func run2(inputText string) int {
	// keep a running max of the r/g/b we've seen in each game
	lines := parsers.SplitByLines(inputText)
	powersum := 0
	for gamenum, line := range lines {
		g := game{num: gamenum, red: 0, green: 0, blue: 0}
		for _, pull := range strings.Split(line, ";") {
			g.red = max(g.red, extractnum("red", pull))
			g.blue = max(g.blue, extractnum("blue", pull))
			g.green = max(g.green, extractnum("green", pull))
		}
		power := g.red * g.blue * g.green
		powersum += power
	}
	return powersum
}
