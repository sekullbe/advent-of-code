package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	countOfEasyNumbers := 0
	for _, s := range parsers.SplitByLines(inputText) {
		inputs, outputs := parseLine(s)
		_ = inputs
		for _, output := range outputs {
			if output.num > -1 {
				countOfEasyNumbers++
			}
		}
	}
	return countOfEasyNumbers
}

func run2(inputText string) int {
	total := 0
	for _, line := range parsers.SplitByLines(inputText) {
		inputs, outputs := parseLine(line)
		deductions := deduceCodes(inputs)
		num := 0
		for i, output := range outputs {
			n := deductions[output.pattern]
			num += tools.PowInt(10, 3-i) * n
		}
		total += num
	}
	return total
}

type digitPattern struct {
	pattern string
	num     int // -1 until we know it for sure
}

func newDigitPattern(p string) digitPattern {
	dp := digitPattern{pattern: SortString(p), num: -1}
	switch len(p) {
	case 2:
		dp.num = 1
	case 3:
		dp.num = 7
	case 4:
		dp.num = 4
	case 7:
		dp.num = 8
	default:
		dp.num = -1 // we can't know yet
	}
	return dp
}

// really only useful for testing, so I can feed in tex lines instead of manually creating []digitPattern{...}s
func deduceCodesFromInputLine(line string) map[string]int {
	inputs, _ := parseLine(line)
	return deduceCodes(inputs)
}

func deduceCodes(inputs []digitPattern) map[string]int {
	oneElts := []rune{}
	fourElts := []rune{}
	deductions := make(map[string]int)
	for _, input := range inputs {
		if input.num == 1 {
			oneElts = append(oneElts, rune(input.pattern[0]))
			oneElts = append(oneElts, rune(input.pattern[1]))
		}
	}
	// gratuitously inefficient but I do not care
	for _, input := range inputs {
		// need to get the four elts that are not one elts
		if input.num == 4 {
			for _, r := range input.pattern {
				if !tools.ContainsRune(oneElts, r) {
					fourElts = append(fourElts, r)
				}
			}
		}
	}
	for _, input := range inputs {
		if input.num > -1 {
			deductions[input.pattern] = input.num
			continue
		}
		eltsInOne := tools.CountElementsInString(oneElts, input.pattern)
		eltsInFour := tools.CountElementsInString(fourElts, input.pattern)
		if len(input.pattern) == 5 {
			if eltsInOne == 1 && eltsInFour == 1 {
				deductions[input.pattern] = 2
			}
			if eltsInOne == 2 && eltsInFour == 1 {
				deductions[input.pattern] = 3
			}
			if eltsInOne == 1 && eltsInFour == 2 {
				deductions[input.pattern] = 5
			}
		} else if len(input.pattern) == 6 {
			if eltsInOne == 1 && eltsInFour == 2 {
				deductions[input.pattern] = 6
			}
			if eltsInOne == 2 && eltsInFour == 2 {
				deductions[input.pattern] = 9
			}
			if eltsInOne == 2 && eltsInFour == 1 {
				deductions[input.pattern] = 0
			}
		}
	}
	return deductions
}

func parseLine(line string) (inputs []digitPattern, outputs []digitPattern) {
	inputMode := true
	for _, s := range strings.Fields(line) {
		if s == "|" {
			inputMode = false
			continue
		}
		dp := newDigitPattern(s)
		if inputMode {
			inputs = append(inputs, dp)
		} else {
			outputs = append(outputs, dp)
		}
	}
	return
}
