package main

import (
	_ "embed"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/polygon"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

var rectangleCache map[string][]geometry.Point2

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	tiles := parseTiles(input)
	maxArea := 0

	for _, t1 := range tiles {
		for _, t2 := range tiles {
			a := geometry.Area(t1, t2)
			maxArea = max(maxArea, a)
		}

	}

	return maxArea
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	// this is way too much brute force- it's going to take hours and
	// the cache is using 70+ GB
	// it's not the inside check that's eating time

	redTiles := parseTiles(input)
	pf := polygon.NewPolyfence(redTiles)
	rectangleCache = make(map[string][]geometry.Point2)

	onTheLine := mapset.NewSet[geometry.Point2]()
	for i := 0; i < len(redTiles)-1; i++ {
		onTheLine.Append(allTilesInRectangle(redTiles[i], redTiles[i+1])...)
	}
	onTheLine.Append(allTilesInRectangle(redTiles[len(redTiles)-1], redTiles[0])...)

	maxArea := 0

	for _, t1 := range redTiles {
		for _, t2 := range redTiles {
			//log.Printf("Considering %v-%v", t1, t2)
			a := geometry.Area(t1, t2)
			// skip it if it can't possibly be better
			if a < maxArea {
				//log.Printf("Bailing, a=%d", a)
				continue
			}
			//log.Printf("Testing, a=%d", a)

			// compute all redTiles in the t1-t2 rectangle
			// and check that each one is inside the polyfence
			// this is the problem, because those are huge

			// first compute the other 2 corners of the rectangle and see if _they_ are inside- if not, bail
			inside := true
			for _, c := range allCornersGivenTwo(t1, t2) {
				if !(pf.Inside(c) || onTheLine.Contains(c)) { // Inside doesn't count being _on_ the line
					inside = false
					//log.Printf("Corner not inside at pt %v", c)
					break
				}
			}
			if !inside {
				continue
			}
			//log.Printf("trying innards")
			// instead of checking every point, try checking the borders only
			// faster still would be to check for intersections but it works in 1m like this

			for _, p := range allPointsOnRectangleBorder(t1, t2) {
				//for _, p := range allTilesInRectangle(t1, t2) {
				if !(pf.Inside(p) || onTheLine.Contains(p)) { // Inside doesn't count being _on_ the line
					inside = false
					//log.Printf("Innards not inside at pt %v", p)
					break
				}
			}
			if !inside {
				continue
			}

			maxArea = max(maxArea, a)
			//log.Printf("Inside with area %d. Max area %d", a, maxArea)
		}

	}

	return maxArea
}

func parseTiles(input string) []geometry.Point2 {
	tiles := []geometry.Point2{}
	lines := parsers.SplitByLines(input)
	for _, line := range lines {
		xy := parsers.StringsWithCommasToIntSlice(line)
		tiles = append(tiles, geometry.Point2{xy[0], xy[1]})
	}
	return tiles
}

func allTilesInRectangle(t1, t2 geometry.Point2) []geometry.Point2 {
	v, ok := rectangleCache[fmt.Sprintf("%v-%v", t1, t2)]
	if ok {
		return v
	}
	minX := min(t1.X, t2.X)
	minY := min(t1.Y, t2.Y)
	maxX := max(t1.X, t2.X)
	maxY := max(t1.Y, t2.Y)
	points := []geometry.Point2{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			points = append(points, geometry.Point2{x, y})
		}
	}
	rectangleCache[fmt.Sprintf("%v-%v", t1, t2)] = points
	return points
}

func allCornersGivenTwo(t1, t2 geometry.Point2) []geometry.Point2 {
	minX := min(t1.X, t2.X)
	minY := min(t1.Y, t2.Y)
	maxX := max(t1.X, t2.X)
	maxY := max(t1.Y, t2.Y)
	return []geometry.Point2{{minX, minY}, {minX, maxY}, {maxX, minY}, {maxX, maxY}}
}

func allPointsOnRectangleBorder(t1, t2 geometry.Point2) []geometry.Point2 {
	minX := min(t1.X, t2.X)
	minY := min(t1.Y, t2.Y)
	maxX := max(t1.X, t2.X)
	maxY := max(t1.Y, t2.Y)
	points := []geometry.Point2{}
	for x := minX; x <= maxX; x++ {
		points = append(points, geometry.Point2{x, minY})
		if minY != maxY {
			points = append(points, geometry.Point2{x, maxY})
		}
	}
	for y := minY + 1; y <= maxY-1; y++ {
		points = append(points, geometry.Point2{minX, y})
		if minX != maxX {
			points = append(points, geometry.Point2{maxX, y})
		}
	}
	return points
}
