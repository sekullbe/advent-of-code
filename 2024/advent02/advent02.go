package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"slices"
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
		if testReport(report) {
			safe++
		}
	}
	return safe
}

func run2(input string) int {
	safe := 0
	reports := parseReports(input)
	for _, report := range reports {
		if testReport(report) {
			safe++
		} else {
			drs := generateDampenedReports(report)
			// if any of dr pass then it's safe
			for _, dr := range drs {
				if testReport(dr) {
					safe++
					break
				}
			}
		}
	}
	return safe
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

func generateDampenedReports(r []int) [][]int {
	dampenedLevels := [][]int{}
	for i := 0; i < len(r); i++ {
		dl := slices.Clone(r)
		dampenedLevels = append(dampenedLevels, tools.RemoveElt(dl, i))
	}
	return dampenedLevels
}

func testReport(report []int) bool {
	return testallOneDirection(report) && testDiff(report, 1, 3)
}
