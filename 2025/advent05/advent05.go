package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

type idrange struct {
	start int
	end   int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")
	freshCount := 0
	freshRanges, ingredientIds := parseInput(input)
	// it's not reasonable to put all valid IDs into a map, but they are sorted which is a small optimization
	for _, id := range ingredientIds {
		for _, rng := range freshRanges {
			if freshcheck(id, rng) {
				freshCount++
				break
			}
			// stop if it's no longer possible for the ranges to match
			if rng.start > id {
				break
			}

		}
	}
	return freshCount
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	freshRanges, _ := parseInput(input)
	mergedRanges := []idrange{freshRanges[0]}
	// for each range, see if it overlaps any range in mergedranges
	// if it does, extend that range
	// if not, add it
	// then count and sum the length of each range since they do not overlap
	for _, fr := range freshRanges {
		didOverlap := false
		for i, mr := range mergedRanges {
			if rangesOverlap(fr, mr) {
				mergedRanges[i] = mergeRanges(fr, mr)
				didOverlap = true
			}
		}
		if !didOverlap {
			mergedRanges = append(mergedRanges, fr)
		}
	}
	sumRanges := 0
	lo.ForEach(mergedRanges, func(r idrange, _ int) {
		sumRanges += (r.end - r.start + 1)
	})

	return sumRanges
}

func parseInput(input string) ([]idrange, []int) {
	inputParts := parsers.SplitByEmptyNewlineToSlices(input)

	idranges := lo.Map(inputParts[0], func(rng string, _ int) idrange {
		rngParts := strings.Split(rng, "-")
		return idrange{tools.Atoi(rngParts[0]), tools.Atoi(rngParts[1])}
	})
	// now sort the slice by range start
	slices.SortFunc(idranges, func(a, b idrange) int {
		return cmp.Compare(a.start, b.start)
	})

	ingredientIds := lo.Map(inputParts[1], func(id string, _ int) int {
		return tools.Atoi(id)
	})
	return idranges, ingredientIds
}

func freshcheck(id int, rng idrange) bool {
	return id >= rng.start && id <= rng.end
}

func rangesOverlap(a, b idrange) bool {
	return a.start <= b.end && b.start <= a.end
}

// this will fail if a & b do not overlap, so don't do that
func mergeRanges(a, b idrange) idrange {
	return idrange{tools.MinInt(a.start, b.start), tools.MaxInt(a.end, b.end)}
}
