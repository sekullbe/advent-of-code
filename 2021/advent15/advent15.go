package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/yourbasic/graph"
)

//go:embed input.txt
var inputText string

type point struct {
	row int
	col int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 100))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText, 100, 5))
}

// graphSize so I don't have to keep figuring it out or passing it around
func run1(inputText string, graphSize int) int {
	caveMap, rows, cols := parseToMap(inputText, 1)
	caveGraph, endnode := parseToGraph(caveMap, rows, cols)
	path, dist := graph.ShortestPath(caveGraph, 0, endnode)
	for _, node := range path {
		r, c := pointIdToCoordinates(node, graphSize)
		fmt.Printf("%d,%d\n", r, c)
	}
	return int(dist)
}

func run2(inputText string, graphSize int, tiles int) int {
	caveMap, rows, cols := parseToMap(inputText, tiles)
	caveGraph, endnode := parseToGraph(caveMap, rows, cols)
	path, dist := graph.ShortestPath(caveGraph, 0, endnode)
	for _, node := range path {
		r, c := pointIdToCoordinates(node, graphSize*tiles)
		fmt.Printf("%d,%d\n", r, c)
	}
	return int(dist)

}

func (p point) neighbors(maxRow, maxCol int) []point {
	neighbors := []point{}
	if p.row-1 > 0 {
		neighbors = append(neighbors, point{p.row - 1, p.col})
	}
	if p.row+1 <= maxRow {
		neighbors = append(neighbors, point{p.row + 1, p.col})
	}
	if p.col-1 > 0 {
		neighbors = append(neighbors, point{p.row, p.col - 1})
	}
	if p.col+1 <= maxCol {
		neighbors = append(neighbors, point{p.row, p.col + 1})
	}
	return neighbors
}

func (p point) uniquePointId(rows int) int {
	return p.row*rows + p.col
}

func pointIdToCoordinates(id, rows int) (row, col int) {
	row = id / rows
	col = id % rows
	return
}

func parseToMap(inputText string, tiles int) (map[point]int, int, int) {
	caveMap := make(map[point]int)
	lines := parsers.SplitByLines(inputText)
	rows := len(lines) // r/c of the original tile
	cols := len(lines[0])
	for row, line := range lines {
		//line := copyLineForTiles(line, tiles)
		for col, riskRune := range line {
			for tileR := 0; tileR < tiles; tileR++ {
				for tileC := 0; tileC < tiles; tileC++ {
					p := point{row + (rows * tileR), col + (cols * tileC)}
					risk := computeRollingRisk(int(riskRune-'0'), tileR+tileC)
					//fmt.Printf("%d,%d=%d\n", p.col, p.row, risk)
					caveMap[p] = risk
				}
			}
		}
	}
	return caveMap, rows * tiles, cols * tiles
}

// returns graph and ID of end node ( so we don't have to compute it later)
func parseToGraph(caveMap map[point]int, rows, cols int) (*graph.Mutable, int) {
	caveGraph := graph.New(len(caveMap))
	for p, _ := range caveMap {
		for _, n := range p.neighbors(rows-1, cols-1) {
			caveGraph.AddCost(p.uniquePointId(rows), n.uniquePointId(rows), int64(caveMap[n]))
		}
	}
	return caveGraph, point{rows - 1, cols - 1}.uniquePointId(rows)
}

func copyLineForTiles(line string, tiles int) string {
	if tiles < 1 {
		panic("can't have zero tiles, and negative tiles is just weird")
	}
	var newLine string
	for i := 0; i < tiles; i++ {
		for _, riskRune := range line {
			risk := computeRollingRisk(int(riskRune-'0'), i)
			newLine += fmt.Sprintf("%d", risk)
		}
	}
	return newLine
}

func computeRollingRisk(risk int, roll int) int {
	risk += roll
	if risk > 9 {
		risk %= 9
	}
	return risk
}
