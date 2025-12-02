package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

type idrange struct {
	first, last int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	ranges := parseranges(input)
	invalidIds := []int{}
	for _, r := range ranges {
		for id := r.first; id <= r.last; id++ {
			if !testIdIsValidSingleRepeat(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	return tools.SumSlice(invalidIds)
}

func testIdIsValidSingleRepeat(id int) bool {
	idstr := strconv.Itoa(id)
	// by definition, an id of odd length must be valid
	if len(idstr)%2 == 1 {
		return true
	}
	// split it in half; if a=b it's invalid
	a := idstr[:len(idstr)/2]
	b := idstr[len(idstr)/2:]
	if a == b {
		//log.Printf("invalid: %s\n", idstr)
		return false
	}
	return true
}

func testIdIsValidMultipleRepeats(id int) bool {
	idstr := strconv.Itoa(id)
	if len(idstr) == 1 {
		return true
	}
	// I'd prefer to use /^(\d+)\1+$/ but go regexp engine doesn't do backreferences
	// For each substring of the id from 1 to len/2, replace that substring with Q
	// If the string is nothing but Q it's invalid

	// Simple optimization- if the substring isn't a factor of the string length, it can't be valid
	for i := 1; i <= len(idstr)/2; i++ {
		if len(idstr)%len(idstr[:i]) != 0 {
			continue
		}
		newStr := strings.ReplaceAll(idstr, idstr[:i], "Q")
		matched, err := regexp.MatchString("^Q+$", newStr)
		if err != nil {
			panic(err)
		}
		if matched {
			return false
		}

	}

	return true
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	ranges := parseranges(input)
	invalidIds := []int{}
	for _, r := range ranges {
		for id := r.first; id <= r.last; id++ {
			if !testIdIsValidMultipleRepeats(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	return tools.SumSlice(invalidIds)
}

func parseranges(input string) []idrange {
	var ranges []idrange
	for _, r := range parsers.SplitByCommasAndTrim(input) {
		var first, last int
		_, err := fmt.Sscanf(r, "%d-%d", &first, &last)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, idrange{first: first, last: last})
	}
	return ranges
}
