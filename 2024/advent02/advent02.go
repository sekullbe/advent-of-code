package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	safe := 0
	reports := parseReports(input)
	for _, report := range reports {
		if testallOneDirection(report) && testDiff(report, 1, 3) {
			safe++
		}
	}
	return safe
}

func run2(input string) int {

	return 0
}

func parseReports(input string) [][]int {
	reports := [][]int{}
	for _, s := range parsers.SplitByLines(input) {
		reports = append(reports, parsers.StringsToIntSlice(s))
	}
	return reports
}

func testallOneDirection(s []int) bool {
	inc := true
	dec := true
	for i := 1; i < len(s); i++ {
		inc = inc && s[i] > s[i-1]
		dec = dec && s[i] < s[i-1]
	}
	return inc || dec
}

// any two entries differ by at least min and at most max (
func testDiff(s []int, min int, max int) bool {
	for i := 1; i < len(s); i++ {
		d := tools.AbsInt(s[i] - s[i-1])
		if d < min || d > max {
			return false
		}
	}
	return true
}
