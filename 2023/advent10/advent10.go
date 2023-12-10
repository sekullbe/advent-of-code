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

	steps := 0

	b := parseBoard(parsers.SplitByLines(input))
	prev := b.start
	curr, _ := b.pipeNeighbors(prev) // choose one arbitrarily
	for {
		steps++
		next := b.followPipe(curr, prev)
		if next == b.start {
			break
		}
		// take the step
		prev = curr
		curr = next
		//fmt.Printf("Step: %v -> %v\n", prev, curr)
	}
	// now we know how may steps in the loop less the final one, so just add 1 and divide by 2
	return (steps + 1) / 2
}

func run2(input string) int {

	// run the full circuit again, except this type keep a list of all the points on the loop, which form a polygon
	b := parseBoard(parsers.SplitByLines(input))
	prev := b.start
	curr, _ := b.pipeNeighbors(prev) // choose one arbitrarily
	points := []Point{b.start, curr}
	for {
		next := b.followPipe(curr, prev)
		points = append(points, next)
		if next == b.start {
			break
		}
		// take the step
		prev = curr
		curr = next
		//fmt.Printf("Step: %v -> %v\n", prev, curr)
	}

	// now for every point not OF the polygon, is it IN the polygon?
	pf := NewPolyfence(points)
	inside := 0
	for point, _ := range b.grid {
		if !tools.Contains(points, point) && pf.Inside(point) {
			//fmt.Printf("Inside: %v\n", point)
			inside++
		}

	}
	// 282 too high
	return inside
}
