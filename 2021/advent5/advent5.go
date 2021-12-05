package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"math"
	"regexp"
	"sort"
	"strconv"
)

//go:embed input.txt
var inputText string

const maxDim int = 999

type chart map[int]map[int]int
type point struct {
	x int
	y int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	intersections := 0
	chart := setUpChart()
	// read coordinates
	for _, line := range parsers.SplitByLines(inputText) {
		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])

		if x1 != x2 && y1 != y2 {
			continue // not horizontal or vertical
		}
		// find all the coordinates in the line and ++ them in the chart
		points := pointsInLine(x1, y1, x2, y2)
		for _, p := range points {
			chart[p.x][p.y]++
			if chart[p.x][p.y] == 2 { // only count the first one
				intersections++
			}
		}
	}

	if maxDim < 80 {
		fmt.Println(chart.toString())
	}
	return intersections
}

func pointsInLine(x1, y1, x2, y2 int) []point {
	var points []point

	if y1 == y2 {
		// horizontal, traverse x
		exxes := []int{x1, x2}
		sort.Ints(exxes)
		for x := exxes[0]; x <= exxes[1]; x++ {
			points = append(points, point{x: x, y: y1})
		}
	} else if x1 == x2 {
		// vertical, traverse y
		whys := []int{y1, y2}
		sort.Ints(whys)
		for y := whys[0]; y <= whys[1]; y++ {
			points = append(points, point{x: x1, y: y})
		}
	} else {
		xj := 1
		yj := 1
		if x2 < x1 {
			xj = -1
		}
		if y2 < y1 {
			yj = -1
		}
		// we know |x1-x2| == |y1-y2|
		for i := 0; ; i += xj {
			points = append(points, point{x: x1 + i, y: y1 + absint(i)*yj})
			if x1+i == x2 {
				break
			}
		}
	}

	return points
}

func absint(i int) int {
	return int(math.Abs(float64(i)))
}

func setUpChart() chart {
	chart := make(chart)
	for x := 0; x < maxDim; x++ {
		chart[x] = make(map[int]int)
		for y := 0; y < maxDim; y++ {
			chart[x][y] = 0
		}
	}
	return chart
}

func run2(inputText string) int {
	intersections := 0
	chart := setUpChart()
	// read coordinates
	for _, line := range parsers.SplitByLines(inputText) {
		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		x2, _ := strconv.Atoi(matches[3])
		y2, _ := strconv.Atoi(matches[4])

		// find all the coordinates in the line and ++ them in the chart
		points := pointsInLine(x1, y1, x2, y2)
		for _, p := range points {
			chart[p.x][p.y]++
			if chart[p.x][p.y] == 2 { // only count the first one
				intersections++
			}
		}
	}

	if maxDim < 80 {
		fmt.Println(chart.toString())
	}

	// 894 too low
	return intersections
}

func (chart *chart) toString() string {
	var out string
	for y := 0; y < maxDim; y++ {
		for x := 0; x < maxDim; x++ {
			out += fmt.Sprintf("%d", (*chart)[x][y])
		}
		out += "\n"
	}
	return out
}
