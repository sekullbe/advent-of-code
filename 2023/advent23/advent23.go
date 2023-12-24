package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/oleiade/lane/v2"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"image"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	//fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type state struct {
	pt    image.Point
	steps int
	seen  *mapset.Set[image.Point]
}

func run1(input string) int {
	b := ParseBoard(parsers.SplitByLines(input))
	// parse to graph, or just try it as is...
	cost := b.pathDfs(Pt(1, 0), Pt(b.MaxX-1, b.MaxY))

	return cost
}

func run2(input string) int {
	input = strings.ReplaceAll(input, "v", ".")
	input = strings.ReplaceAll(input, "^", ".")
	input = strings.ReplaceAll(input, "<", ".")
	input = strings.ReplaceAll(input, ">", ".")
	b := ParseBoard(parsers.SplitByLines(input))
	// this is correct from the tests but takes too long, 3+ hours
	//cost := b.pathDfs(Pt(1, 0), Pt(b.MaxX-1, b.MaxY))
	// revisit 2021 day 15 solution where I turn the board into a graph and DFS that
	// can I reuse a canned DFS or do I need to make a new one?
	// this version stores visited in the stack state instead of globally so I think I need a new one

	// interesting, every possible junction is surrounded by ramps, and ramps exist nowhere else.
	// is that actually useful? len(neighbors) > 2 is just as good

	// reworked https://github.com/pemoreau/advent-of-code/blob/main/go/2023/23/day23.go
	// so I could understand it with my data structures
	neighbors := b.buildGraph(Pt(1, 0))
	visited := make(map[image.Point]bool)
	target := Pt(b.MaxX-1, b.MaxY)
	path := 0

	cost := explore(neighbors, Pt(1, 0), target, visited, path, 0)

	return cost
}

func explore(neighbors Graph, start, goal image.Point, visited map[image.Point]bool, steps int, maxSteps int) int {
	if start == goal {
		if steps > maxSteps {
			maxSteps = steps
		}
		return maxSteps
	}

	visited[start] = true
	for _, pc := range neighbors[start] {
		if !visited[pc.pt] {
			maxSteps = explore(neighbors, pc.pt, goal, visited, steps+pc.steps, maxSteps)
		}
	}
	visited[start] = false
	return maxSteps
}

type Cost struct {
	pt    image.Point
	steps int
}

type Graph map[image.Point][]Cost

func (b Board) buildGraph(start image.Point) Graph {
	var costs = make(Graph)

	var todo = []image.Point{}
	todo = append(todo, start)

	for len(todo) > 0 {
		pt := todo[0]
		todo = todo[1:]
		// this grid is a slice, while mine's a map
		t := b.AtPoint(pt)
		if t.Contents == FOREST {
			continue
		}
		for _, n := range b.GetClearNeighbors(pt) {
			//pc, ok := exploreSinglePath(grid, p, n, 1, true)
			cost, ok := b.followPath(pt, n, 1)
			if ok && !tools.Contains(costs[pt], cost) {
				costs[pt] = append(costs[pt], cost)
				todo = append(todo, cost.pt)
			}
		}
	}
	return costs
}

func (b Board) followPath(prev image.Point, curr image.Point, steps int) (Cost, bool) {
	t := b.AtPoint(curr)
	if t.Contents != FOREST {
		if len(b.GetClearNeighbors(curr)) > 2 {
			return Cost{pt: curr, steps: steps}, true
		}
	}
	for _, n := range b.GetClearNeighbors(curr) {
		if n != prev {
			//return exploreSinglePath(grid, current, n, cost+1, part2)
			return b.followPath(curr, n, steps+1)
		}
	}

	return Cost{pt: curr, steps: steps}, true

}

func (b Board) pathDfs(startPt image.Point, endPt image.Point) int {
	stack := lane.NewStack[state]()
	nset := mapset.NewSet[image.Point]()
	stack.Push(state{pt: startPt, steps: 0, seen: &nset})
	longestPath := 0

	for stack.Size() > 0 {
		st, _ := stack.Pop()

		if st.pt == endPt { // hooray we're done
			longestPath = max(longestPath, st.steps)
			fmt.Printf("winnah! steps=%d\n", st.steps)
			continue
		}

		if (*st.seen).Contains(st.pt) {
			continue
		}
		(*st.seen).Add(st.pt)
		possibleNextStates := []state{}
		// if we're on an arrow we must follow it
		if arrowDir, isArrow := directionOfArrow(b.AtPoint(st.pt).Contents); isArrow {
			// don't need to look - arrows always point safely
			var nseen mapset.Set[image.Point] = (*st.seen).Clone()
			nst := state{
				pt:    NeighborInDirection(st.pt, arrowDir),
				steps: st.steps + 1,
				seen:  &nseen,
			}
			possibleNextStates = append(possibleNextStates, nst)
		} else {
			pt := st.pt
			// try to be smart and if there's only one way out,
			// go to that point and add 1 to steps, skipping all the interim checks
			substeps := 1
			neigbors := b.GetClearNeighbors(st.pt)
			ncount := len(neigbors)
			for ncount == 2 {
				onlyNeighbor := neigbors[0]
				if (*st.seen).Contains(neigbors[0]) {
					onlyNeighbor = neigbors[1]
				}
				if (*st.seen).Contains(neigbors[1]) {
					break
				}
				if onlyNeighbor == endPt {
					break
				}
				//fmt.Printf("Optimize: %v\n", onlyNeighbor)
				substeps += 1
				pt = onlyNeighbor
				(*st.seen).Add(onlyNeighbor)
				neigbors = b.GetClearNeighbors(pt)
				ncount = len(neigbors)
			}
			for _, npt := range neigbors {
				ntile := b.AtPoint(npt)
				// do i need to handle the special case where you can't go into a slope pointing at you?
				//if _, isArrow := directionOfArrow(ntile.Contents); isArrow {
				//
				//}
				var nseen mapset.Set[image.Point] = (*st.seen).Clone()
				if !(*st.seen).Contains(ntile.Pt) {
					possibleNextStates = append(possibleNextStates, state{
						pt:    npt,
						steps: st.steps + substeps,
						seen:  &nseen,
					})
				}
			}
		}
		for _, nextState := range possibleNextStates {
			if b.InRange(nextState.pt) {
				//fmt.Printf("Looking at %v\n", nextState.pt)
				stack.Push(nextState)
			}
		}

	}
	return longestPath
}

func (b Board) GetClearNeighbors(p image.Point) []image.Point {
	ns := []image.Point{}
	for d := NORTH; d <= WEST; d += 2 {
		np := NeighborInDirection(p, d)
		if b.InRange(np) && b.AtPoint(np).Contents != FOREST {
			ns = append(ns, np)
		}
	}
	return ns
}

func directionOfArrow(arrow rune) (int, bool) {
	switch arrow {
	case '<':
		return WEST, true
	case '>':
		return EAST, true
	case '^':
		return NORTH, true
	case 'V':
		fallthrough
	case 'v':
		return SOUTH, true
	default:
		return 0, false
	}
}
