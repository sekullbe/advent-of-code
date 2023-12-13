package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"regexp"
	"slices"
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
	cache  map[[3]int]int
}

func run1(input string) int {
	sum := 0
	for _, line := range parsers.SplitByLines(input) {
		n := parseRow(line)
		n.cache = make(map[[3]int]int)
		//sum += countPossibleArrangements(n)
		sum += n.countPossibleArrangementsBetter(0, 0, 0)
	}
	return sum
}

func run2(input string) int {
	sum := 0
	for _, line := range parsers.SplitByLines(input) {
		n := parseRowWithFolding(line, 5)
		n.cache = make(map[[3]int]int)
		sum += n.countPossibleArrangementsBetter(0, 0, 0)
	}
	return sum
}

func run2StupidSlow(input string) int {
	sum := 0
	for _, line := range parsers.SplitByLines(input) {
		n := parseRowWithFolding(line, 5)
		sum += countPossibleArrangements(n)
		fmt.Print(".")
	}
	fmt.Println()
	return sum
}

func parseRow(line string) nomogramRow {
	n := nomogramRow{}
	fs := strings.Fields(line)
	n.spring = fs[0]
	n.groups = parsers.StringsWithCommasToIntSlice(fs[1])
	return n
}

func parseRowWithFolding(line string, foldFactor int) nomogramRow {
	n := nomogramRow{}
	fs := strings.Fields(line)
	n.spring = fs[0]
	n.groups = parsers.StringsWithCommasToIntSlice(fs[1])
	return unfold(n, foldFactor)
}
func unfold(n nomogramRow, foldFactor int) nomogramRow {
	s2 := n.spring
	g2 := n.groups
	for i := 1; i < foldFactor; i++ {
		s2 += "?" + n.spring
		gc := slices.Clone(n.groups)
		g2 = append(g2, gc...)
	}

	return nomogramRow{spring: s2, groups: g2}
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

	numDamaged := tools.SumSlice(n.groups)
	for i := int64(0); i < max; i++ {
		sb := fmt.Sprintf("%0"+fmt.Sprint(bits)+"b", i)
		// now sub each bit in for a ? in the spring
		spring := baseSpring
		for j := 0; j < bits; j++ {
			spring = strings.Replace(spring, "?", string(sb[j]), 1)
		}
		// optimization- if a state doesn't have the correct number of #, skip it
		// 5x speedup but it's not good enough
		numBits := strings.Count(spring, "1")
		if numBits == numDamaged && validateRowWithNumberDigits(nomogramRow{spring: spring, groups: n.groups}) {
			possibles++
		}
	}
	return possibles
}

// need a CPA that doesn't take O(2^N) time
// go implementation of https://github.com/derailed-dash/Advent-of-Code/blob/master/src/AoC_2023/Dazbo's_Advent_of_Code_2023.ipynb
func (n *nomogramRow) countPossibleArrangementsBetter(charIdx int, currGroupIdx int, currGroupLen int) int {
	if r, ok := n.cache[[3]int{charIdx, currGroupIdx, currGroupLen}]; ok {
		return r
	}
	count := 0
	if charIdx == len(n.spring) { // end of string
		if currGroupIdx == len(n.groups) && currGroupLen == 0 { // finished the last group
			return 1
		}
		//            elif curr_group_idx == len(self.damaged_groups) - 1 and self.damaged_groups[curr_group_idx] == curr_group_len:
		if currGroupIdx == len(n.groups)-1 && n.groups[currGroupIdx] == currGroupLen { // we're on the last char of the last group, and the group is complete
			return 1
		}
		return 0 // we have not completed all groups, or current group length is too long
	}

	// Process the current char in the record by recursion
	// Determine valid states for recursion
	//for char in [".", "#"]:
	for _, c := range []string{".", "#"} {
		if string(n.spring[charIdx]) == c || string(n.spring[charIdx]) == "?" { //  We can subst char for itself (no change), or for ?
			if c == "." {
				//  we're extending the operational section or ending the damaged group
				if currGroupLen == 0 {
					// not in a group so must be extending
					count += n.countPossibleArrangementsBetter(charIdx+1, currGroupIdx, 0)
				} else if currGroupIdx < len(n.groups) && currGroupLen == n.groups[currGroupIdx] {
					// we're adding a . after a #, so the group is now complete; move on to next group
					count += n.countPossibleArrangementsBetter(charIdx+1, currGroupIdx+1, 0)
				}
			} else {
				//we're adding a #; extend the current group (which might be empty at this point)
				count += n.countPossibleArrangementsBetter(charIdx+1, currGroupIdx, currGroupLen+1)
			}
		}
	}
	n.cache[[3]int{charIdx, currGroupIdx, currGroupLen}] = count
	return count
}
