package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"time"
)

//go:embed input.txt
var inputText string

type lock = [5]int
type key = [5]int

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")
	locks, keys := parseLocksAndKeys(input)

	fits := 0

	for _, l := range locks {
		for _, k := range keys {
			couldfit := true
			for col := 0; col < 5; col++ {
				if l[col]+k[col] > 5 {
					couldfit = false
				}
			}
			if couldfit {
				fits++
			}

		}

	}

	return fits
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	return 0
}

func parseLocksAndKeys(input string) ([]lock, []key) {
	locks := []lock{}
	keys := []key{}

	blocks := parsers.SplitByEmptyNewlineToSlices(input)
	for _, block := range blocks {
		if block[0][0] == '#' {
			locks = append(locks, parseLock(block))
		} else {
			keys = append(keys, parseKey(block))
		}

	}
	return locks, keys
}

func parseLock(block []string) lock {
	l := lock{-1, -1, -1, -1, -1}
	for i := 1; i < len(block); i++ {
		for col := 0; col < 5; col++ {
			if block[i][col] == '.' && l[col] == -1 {
				l[col] = i - 1
			}
		}
	}
	return l
}

func parseKey(block []string) key {
	k := key{-1, -1, -1, -1, -1}
	for i := len(block) - 2; i >= 0; i-- {
		for col := 0; col < 5; col++ {
			if block[i][col] == '.' && k[col] == -1 {
				k[col] = 5 - i
			}
		}
	}
	return k
}
