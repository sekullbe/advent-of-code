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

	var sum int
	for _, line := range parsers.SplitByLines(inputText) {
		sum += calcCalibrationDigitsOnly(line)
	}

	return sum
}

func run2(inputText string) int {
	var sum int
	for _, line := range parsers.SplitByLines(inputText) {
		sum += calcCalibrationDigitsOnly(replaceStringDigits(line))
	}

	return sum
}

func calcCalibrationDigitsOnly(line string) int {
	var first, last int32

	zero := '0'
	nine := '9'
	_ = zero
	_ = nine

	for _, r := range line {
		if r >= '0' && r <= '9' {
			first = r
			break
		}
	}
	for _, r := range tools.ReverseString(line) {
		if r >= '0' && r <= '9' {
			last = r
			break
		}
	}
	return int((first-'0')*10 + last - '0')
}

// This is kind of a gross hack- replace the number strings with a string that contains
// the number but can't break any concatentation that might make other numbers
// example: 'eighttwothree' needs to see the 8 first, but purely replacing the 'two' breaks the 'eight'
// that is really first.
func replaceStringDigits(line string) string {
	out := line
	out = strings.Replace(out, "zero", "ze0ro", -1)
	out = strings.Replace(out, "one", "o1ne", -1)
	out = strings.Replace(out, "two", "tw2o", -1)
	out = strings.Replace(out, "three", "thr3ee", -1)
	out = strings.Replace(out, "four", "fo4ur", -1)
	out = strings.Replace(out, "five", "fi5ve", -1)
	out = strings.Replace(out, "six", "si6x", -1)
	out = strings.Replace(out, "seven", "se7ven", -1)
	out = strings.Replace(out, "eight", "eig8ht", -1)
	out = strings.Replace(out, "nine", "ni9ne", -1)
	return out
}

/*
func findFirstDigit(line string) int {
	var digits = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|\\d)")

}
*/
