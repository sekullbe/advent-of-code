package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"sort"
)

//go:embed input.txt
var inputText string

type caveChart [][]point

//type caveChartRow map[int]point
type visitedChart = [][]bool

type world struct {
	chart   caveChart
	visited visitedChart
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	chart := parseChart(inputText)
	riskSum := 0
	for ir, row := range chart {
		for ic, p := range row {
			if isLowPoint(chart, ir, ic) {
				risk := p.height + 1
				riskSum += risk
			}
		}
	}
	return riskSum
}

type point struct {
	row    int
	col    int
	height int
	basin  int
}

func run2(inputText string) int {
	chart := parseChart(inputText)
	lowPoints := []point{}
	for ir, row := range chart {
		for ic, _ := range row {
			if isLowPoint(chart, ir, ic) {
				lowPoints = append(lowPoints, chart[ir][ic])
			}
		}
	}

	world := world{chart: chart}
	world.setVisitedMatrixToFalse()

	// now floodfill from each of those points
	for i, lowPoint := range lowPoints {
		world.floodFill(lowPoint.row, lowPoint.col, i+1)
	}
	//fmt.Println(world.chart.printBasins())

	// and now measure the size of each basin
	sizeOfBasins := world.countSizeOfBasins(len(lowPoints))
	sort.Ints(sizeOfBasins)
	biggestThree := sizeOfBasins[len(sizeOfBasins)-3:]

	magic := 1
	for _, bs := range biggestThree {
		magic *= bs
	}

	return magic
}

func (w *world) countSizeOfBasins(howManyBasins int) []int {
	basinSizes := make([]int, howManyBasins+1)
	for _, row := range w.chart {
		for _, p := range row {
			if p.basin > 0 {
				basinSizes[p.basin]++
			}
		}
	}
	return basinSizes
}

func (chart caveChart) printBasins() string {
	var out string
	for r := 0; r < len(chart); r++ {
		for c := 0; c < len(chart[r]); c++ {
			if chart[r][c].basin == -1 {
				out += "?"
			} else if chart[r][c].basin == 0 {
				out += "*"
			} else {
				out += fmt.Sprintf("%d", chart[r][c].basin)
			}
		}
		out += "\n"
	}
	return out
}

// Floodfill algorithm modified from:
// https://nathanleclaire.com/blog/2014/04/05/implementing-a-concurrent-floodfill-with-golang/

func (w *world) floodFill(x, y int, basin int) {
	// If unbuffered, this channel will block when we go to send the
	// initial nodes to visit (at most 4).  Not cool man.
	toVisit := make(chan point, 4)
	visitDone := make(chan bool)

	originalbasin := w.chart[x][y].basin

	w.setVisitedMatrixToFalse()

	go w.doFloodFill(x, y, basin, originalbasin, toVisit, visitDone)
	remainingVisits := 1

	for {
		select {
		case nextVisit := <-toVisit:
			if !w.visited[nextVisit.row][nextVisit.col] {
				w.visited[nextVisit.row][nextVisit.col] = true
				remainingVisits++
				go w.doFloodFill(nextVisit.row, nextVisit.col, basin, originalbasin, toVisit, visitDone)
			}
		case <-visitDone:
			remainingVisits--
		default:
			if remainingVisits == 0 {
				return
			}
		}
	}
}

func (w *world) doFloodFill(x int, y int, basin int, originalbasin int, toVisit chan point, visitDone chan bool) {
	w.chart[x][y].basin = basin
	neighbors := w.getAndMarkNeighbors(x, y, basin)
	for _, neighbor := range neighbors {
		if neighbor.basin == originalbasin {
			toVisit <- neighbor
		}
	}
	visitDone <- true
}

func (w *world) setVisitedMatrixToFalse() {
	width := len(w.chart)
	height := len(w.chart[0])
	w.visited = make([][]bool, width)
	for i := 0; i < width; i++ {
		w.visited[i] = make([]bool, height)
		for j := 0; j < height; j++ {
			w.visited[i][j] = false
		}
	}
}

func (w *world) getAndMarkNeighbors(row, col int, basin int) []point {
	return neighborPoints(w.chart, row, col)
}

func neighborPoints(chart caveChart, row, col int) []point {
	neighbors := []point{}
	var b int
	if row-1 >= 0 {
		b = chart[row-1][col].basin
		neighbors = append(neighbors, point{row: row - 1, col: col, height: chart[row-1][col].height, basin: b})
	}
	if row+1 < len(chart) {
		b = chart[row+1][col].basin
		neighbors = append(neighbors, point{row: row + 1, col: col, height: chart[row+1][col].height, basin: b})
	}
	if col-1 >= 0 {
		b = chart[row][col-1].basin
		neighbors = append(neighbors, point{row: row, col: col - 1, height: chart[row][col-1].height, basin: b})
	}
	if col+1 < len(chart[row]) {
		b = chart[row][col+1].basin
		neighbors = append(neighbors, point{row: row, col: col + 1, height: chart[row][col+1].height, basin: b})
	}
	return neighbors
}

func neighborHeights(chart caveChart, row, col int) []int {
	neighbors := neighborPoints(chart, row, col)
	nd := []int{}
	for _, neighbor := range neighbors {
		nd = append(nd, neighbor.height)
	}
	return nd
}

func isLowPoint(chart caveChart, row, col int) bool {
	height := chart[row][col].height
	for _, nh := range neighborHeights(chart, row, col) {
		if nh <= height {
			return false
		}
	}
	return true
}

func parseChart(inputText string) caveChart {
	lines := parsers.SplitByLines(inputText)
	chart := make(caveChart, len(lines))
	for row, line := range lines {
		if len(line) == 0 {
			continue
		}
		chart[row] = make([]point, len(line))
		for col, r := range line {
			p := point{row: row, col: col, height: int(r - '0'), basin: -1}
			if p.height == 9 {
				p.basin = 0
			}
			chart[row][col] = p
		}
	}
	return chart
}
