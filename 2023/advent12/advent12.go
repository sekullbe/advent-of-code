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

type nomogramRow struct {
	spring string // maybe this can just be a string
	groups []int
}

func run1(input string) int {
	sum := 0
	for _, line := range parsers.SplitByLines(input) {
		n := parseRow(line)
		sum += countPossibleArrangements(n)
		fmt.Print(".")
	}
	fmt.Println()
	return sum
}

func run2(input string) int {

	return 0
}

func parseRow(line string) nomogramRow {
	n := nomogramRow{}
	fs := strings.Fields(line)
	n.spring = fs[0]
	n.groups = parsers.StringsWithCommasToIntSlice(fs[1])
	return n
}

func validateRow(n nomogramRow) bool {
	// turn each group into a regexp \s+#{n}
	// pad the row with spaces front & back
	// and match
	row := "." + n.spring + "."
	var restr string = `^\.*`
	for _, group := range n.groups {
		restr += fmt.Sprintf(`\.+#{%d}`, group)
	}
	restr += `\.*$`
	re := regexp.MustCompile(restr)
	return re.MatchString(row)
}

func validateRowWithNumberDigits(n nomogramRow) bool {
	// turn each group into a regexp \s+#{n}
	// pad the row with spaces front & back
	// and match
	row := "0" + n.spring + "0"
	var restr string = `^0*`
	for _, group := range n.groups {
		restr += fmt.Sprintf(`0+1{%d}`, group)
	}
	restr += `0*$`
	re := regexp.MustCompile(restr)
	return re.MatchString(row)
}

// brutest of brute force
// take the number of ? in the spring and calculate 2^N
// for i = 0 to 2^n
// sub each digit into each ?- 0 is . and 1 is #
// validate the row
func countPossibleArrangements(n nomogramRow) int {
	possibles := 0
	bits := strings.Count(n.spring, "?")
	max := tools.PowInt64(2, bits)
	baseSpring := strings.ReplaceAll(n.spring, ".", "0")
	baseSpring = strings.ReplaceAll(baseSpring, "#", "1")
	for i := int64(0); i < max; i++ {
		sb := fmt.Sprintf("%0"+fmt.Sprint(bits)+"b", i)
		// now sub each bit in for a ? in the spring
		spring := baseSpring
		for j := 0; j < bits; j++ {
			spring = strings.Replace(spring, "?", string(sb[j]), 1)
		}
		if validateRowWithNumberDigits(nomogramRow{spring: spring, groups: n.groups}) {
			possibles++
		}
	}
	return possibles
}
