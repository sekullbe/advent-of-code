package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
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

	sum := 0
	patterns := parsers.SplitByEmptyNewlineToSlices(input)
	for _, pattern := range patterns {
		mp, vertical := findMirrorPoint(pattern)
		if vertical {
			sum += mp
		} else {
			sum += 100 * mp
		}

	}

	return sum
}

func run2(input string) int {

	return 0
}

// finds the mirror points(mirroring between N and N+1 returns N) and returns  true if it exists
// uses the term 'row' but could be a column
// returns # of columns to the left/above the mirror point
// there can be more than one
func findMirrorPoints(row string) []int {
	mps := []int{}
	for mp := 1; mp < len(row); mp++ { // 1 can't be the mirror point
		left := row[:mp]
		right := row[mp:]
		diff := len(left) - len(right)
		if diff > 0 {
			left = left[diff:]
		} else if diff < 0 { // right is longer, so trim it to the same length as left
			right = right[:len(right)+diff]
		}
		if left == tools.ReverseString(right) {
			//fmt.Printf("%s-%s %d\n", left, right, mp)
			mps = append(mps, mp)
		}
	}
	return mps
}

func findMirrorPoint(pattern []string) (mp int, vertical bool) {

	var found bool
	mp, found = findSingleMirrorPoint(pattern)
	if found {
		return mp, true
	}
	mp, found = findSingleMirrorPoint(rotatePattern(pattern))
	if !found {
		panic("could not find a mirror in either direction")
	}
	return mp, false

}

func findSingleMirrorPoint(pattern []string) (int, bool) {

	mps := []mapset.Set[int]{}
	for _, row := range pattern {
		mps = append(mps, mapset.NewSet[int](findMirrorPoints(row)...))
	}
	singlemp := mps[0].Clone()
	for _, mp := range mps {
		singlemp = singlemp.Intersect(mp)
	}
	if singlemp.Cardinality() == 1 {
		return singlemp.Pop()
	}
	return 0, false

}

/*
rotate 90deg CCW, so the top becomes the left
we want to calculate cols from the left as rows from the top

ab    bdfh
cd -> aceg
ef
gh
*/
// columns -> rows, starting from the right
func rotatePattern(pattern []string) []string {
	out := []string{}
	width := len(pattern[0])
	height := len(pattern)
	for x := width - 1; x >= 0; x-- {
		var row string
		for y := 0; y < height; y++ {
			row += string(pattern[y][x])
		}
		out = append(out, row)
	}
	return out
}
