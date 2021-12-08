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
			if tools.ContainsInt(output.possibleNums, 1) {
				countOfEasyNumbers++
			}
			if tools.ContainsInt(output.possibleNums, 4) {
				countOfEasyNumbers++
			}
			if tools.ContainsInt(output.possibleNums, 7) {
				countOfEasyNumbers++
			}
			if tools.ContainsInt(output.possibleNums, 8) {
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
	pattern      string
	possibleNums []int
	num          int // -1 until we know it for sure
}

func newDigitPattern(p string) digitPattern {
	dp := digitPattern{pattern: SortString(p), num: -1}
	switch len(p) {
	case 2:
		dp.possibleNums = []int{1}
		dp.num = 1
	case 3:
		dp.possibleNums = []int{7}
		dp.num = 7
	case 4:
		dp.possibleNums = []int{4}
		dp.num = 4
	case 5:
		dp.possibleNums = []int{2, 3, 5}
	case 6:
		dp.possibleNums = []int{6, 9, 0}
	case 7:
		dp.possibleNums = []int{8}
		dp.num = 8
	default:
		panic("impossible digit pattern: " + p)
	}

	return dp
}

// really only useful for testing
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
	// gratuitiously inefficient but I do not care
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
			continue
		}
		if len(input.pattern) == 6 {
			if eltsInOne == 1 && eltsInFour == 2 {
				deductions[input.pattern] = 6
			}
			if eltsInOne == 2 && eltsInFour == 2 {
				deductions[input.pattern] = 9
			}
			if eltsInOne == 2 && eltsInFour == 1 {
				deductions[input.pattern] = 0
			}
			continue
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
