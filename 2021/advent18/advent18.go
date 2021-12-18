package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

type numPair struct {
	one string
	two string
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	nums := parsers.SplitByLines(inputText)
	result := repeatAdd(nums)
	return magnitude(result)
}

func run2(inputText string) int {
	maxMagnitude := 0
	nums := parsers.SplitByLines(inputText)
	pairs := getAllNumPairs(nums)
	for _, pair := range pairs {
		mag := magnitude(repeatAdd([]string{pair.one, pair.two}))
		if mag > maxMagnitude {
			maxMagnitude = mag
		}
	}
	return maxMagnitude
}

func getAllNumPairs(nums []string) []numPair {
	var pairs []numPair
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pairs = append(pairs, numPair{one: nums[i], two: nums[j]})
			pairs = append(pairs, numPair{two: nums[i], one: nums[j]})
		}
	}
	return pairs
}

func magnitude(num string) int {
	elts := splitPair(num)
	if len(elts) == 1 {
		mag, _ := strconv.Atoi(num)
		return mag
	}
	return 3*magnitude(elts[0]) + 2*magnitude(elts[1][0:len(elts[1])])
}

// returns the pairs and strips the outer [ and ]
// If it's not a pair, returns 1-elt slice containing the number
func splitPair(num string) []string {
	lefts := -1 // so we ignore the opening [
	rights := 0
	for i, r := range num {
		switch r {
		case '[':
			lefts++
		case ']':
			rights++
		case ',':
			if lefts == rights {
				return []string{num[1:i], num[i+1 : len(num)-1]}
			}
		}
	}
	return []string{num}
}

func repeatAdd(nums []string) string {
	num := nums[0]
	for _, n := range nums[1:] {
		num = add(num, n)
	}
	return num
}

func add(num1, num2 string) string {
	return reduce(fmt.Sprintf("[%s,%s]", num1, num2))
}

func reduce(num string) string {
	for {
		// explode until there's nothing to explode
		// doing it all in the for conditions had some scoping problems
		//for num, exploded := explode(num); exploded; num, exploded = explode(num) {}
		for {
			var exploded bool
			num, exploded = explode(num)
			if !exploded {
				break
			}
		}

		// look for splits and do the first one
		var didSplit bool
		num, didSplit = doTheFirstSplit(num)
		if !didSplit {
			break
		}
	}
	return num
}

// No-op if there's nothing to split; returns original string and false
func doTheFirstSplit(num string) (string, bool) {
	re := regexp.MustCompile(`\d{2,}`)
	loc := re.FindStringIndex(num)
	if loc == nil {
		return num, false
	}
	splitNumStr := num[loc[0]:loc[1]]
	splitNum, _ := strconv.Atoi(splitNumStr)
	newLeft := splitNum / 2        // round down   6->3 7->3
	newRight := (splitNum + 1) / 2 // round up 6->7->3 7->8->4

	return fmt.Sprintf("%s[%d,%d]%s", num[:loc[0]], newLeft, newRight, num[loc[1]:]), true
}

// No-op if there's nothing to explode, returns false in that case
func explode(num string) (newNum string, exploded bool) {
	explodo, boomtoday := lookForExploders(num)
	if boomtoday {
		num = explodeAtIndex(num, explodo)
	}
	return num, boomtoday
}

// returns index of beginning pair and true if found
func lookForExploders(num string) (int, bool) {
	depth := 0
	for i, r := range num {
		if r == '[' {
			depth++
		}
		if r == ']' {
			depth--
		}
		if depth == 5 {
			return i, true
		}
	}
	return 0, false
}

func explodeAtIndex(num string, explodo int) string {
	// [[6,[5,[4,[3,2]]]],1]
	//  index -> *
	// read the next [\d+,\d+] - the nums can be 10+ because we finish exploding before splitting
	// replace [3,2] (or [32,45]) with 0
	// new string is num [0:explodo] + "0" + num[explodo+length:]
	// seek left from explodo for a \d, replace it with x+=left
	// seek right for a \d, replace it with x+right

	re := regexp.MustCompile(`\[(\d+),(\d+)\]`)
	matches := re.FindStringSubmatch(num[explodo:])
	lstr := matches[1]
	rstr := matches[2]
	explodoPairLength := len(lstr) + len(rstr) + 3 // 3 for [,]
	left, _ := strconv.Atoi(lstr)
	right, _ := strconv.Atoi(rstr)

	l := seekLeftAndReplace(num[:explodo], left)
	r := seekRightAndReplace(num[explodo+explodoPairLength:], right)

	return l + "0" + r
}

func seekLeftAndReplace(num string, newNum int) string {
	// look left for a number - any number, and it might have > 1 digit
	numIdx := strings.LastIndexAny(num, "0123456789")
	if numIdx == -1 { // handle the case where there are no numbers to the left
		return num
	}
	endNumIdx := numIdx + 1
	// seek until we get a number, which may again have more than one digit
	for ; num[numIdx-1] >= '0' && num[numIdx-1] <= '9'; numIdx-- {
	}
	oldNum, err := strconv.Atoi(num[numIdx:endNumIdx])
	if err != nil {
		log.Panicf("stop! panic time! can't parse this: %s", num[numIdx:endNumIdx])
	}
	return fmt.Sprintf("%s%d%s", num[:numIdx], oldNum+newNum, num[endNumIdx:])
}

func seekRightAndReplace(num string, newNum int) string {
	// look right for a number - any number, and it might be 10+
	numIdx := strings.IndexAny(num, "0123456789")
	if numIdx == -1 { // handle the case where there are no numbers to the right
		return num
	}
	initNumIdx := numIdx
	// seek until we get a number, which may have more than one digit
	for ; num[numIdx] >= '0' && num[numIdx] <= '9'; numIdx++ {
	}
	oldNum, err := strconv.Atoi(num[initNumIdx:numIdx])
	if err != nil {
		log.Panicf("stop! panic time! can't parse this: %s", num[initNumIdx:numIdx])
	}
	return fmt.Sprintf("%s%d%s", num[:initNumIdx], oldNum+newNum, num[numIdx:])
}
