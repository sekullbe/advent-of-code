package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/oleiade/lane/v2"
	"github.com/sekullbe/advent/parsers"
	"image"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 64))
	fmt.Println("-------------")
	//fmt.Printf("Magic number: %d\n", run2(inputText, 26501365))
	fmt.Printf("Magic number: %d\n", run2(inputText, 26501365))
}

type step struct {
	pt        image.Point
	stepsLeft int //
}

func run1(input string, steps int) int {
	// smells like a BFS, and we can backtrack so all directions are valid and make sure to count the backsteps
	// time for another copy of my evolving grid library... this makes 14 copies of it across all my AoC solutions
	b := ParseBoard(parsers.SplitByLines(input))
	reachables := b.bfs(Pt(b.startX, b.startY), steps)
	for _, p := range reachables.ToSlice() {
		t := b.AtPoint(p)
		t.Contents = 'O'
		b.Grid[p] = t
	}
	//b.printBoard()

	return reachables.Cardinality()
}

func (b Board) bfs(start image.Point, maxSteps int) mapset.Set[image.Point] {

	queue := lane.NewQueue[step](step{start, maxSteps})
	seen := mapset.NewSet[image.Point]()
	reachables := mapset.NewSet[image.Point]()

	for {
		s, ok := queue.Dequeue()
		if !ok { // yesterday I did this the other way around and wrapped the code in `if ok {` but this is better
			break
		}

		// idea from Dazbo- if steps left are odd, we can't *finish* on this spot; even means we can.
		if s.stepsLeft >= 0 {
			if s.stepsLeft%2 == 0 { // we can return here
				reachables.Add(s.pt)
			}
			if s.stepsLeft > 0 { // where can we go now?
				s.stepsLeft--                    // take a step, but to where?
				ns := b.GetSquareNeighbors(s.pt) // these are Points not Tiles, but it excludes off-board tiles
				for _, n := range ns {
					if seen.Contains(n) || b.AtPoint(n).Contents == '#' {
						continue
					}
					queue.Enqueue(step{n, s.stepsLeft})
					seen.Add(n)
				}
			}
		}
	}
	return reachables
}

func run2(input string, steps int) int {

	/* from subreddit:
		Let f(n) be the number of spaces you can reach after n steps. Let X be the length of your input grid.
	   f(n), f(n+X), f(n+2X),...., is a quadratic,so you can find it by finding the first 3 values,
			then use that to interpolate the final answer.

		 26501365 = 2023 * 100 * 131 + 65

		but what are the first 3 values?
		defining gridDim as 131, half of it is 65...
		0 * gridDim + half
		1 * gridDim + half
		2 * gridDim + half
		i.e.
		gridDim /2
		gridDim * (3/2)
		gridDim * (5/2)

	    this does not work for the test data because it only works where steps = N * gridDim + gridDim/2
	*/
	b := ParseBoard(parsers.SplitByLines(input))
	gridDim := b.MaxX + 1
	values := [3]int{}

	values[0] = b.bfs2(Pt(b.startX, b.startY), gridDim/2).Cardinality()
	values[1] = b.bfs2(Pt(b.startX, b.startY), (3*gridDim)/2).Cardinality()
	values[2] = b.bfs2(Pt(b.startX, b.startY), (5*gridDim)/2).Cardinality()

	x := simplifiedLagrange(steps, gridDim, values)

	return x
}

func simplifiedLagrange(steps int, size int, values [3]int) int {
	a := values[0]/2 - values[1] + values[2]/2
	b := -3*(values[0]/2) + 2*values[1] - values[2]/2
	c := values[0]
	x := steps / size
	return a*x*x + b*x + c
}

/*
In test I handle 5000 steps in ~30 seconds, so this ain't gonna fly.
so examining the data we see that both the edge of each tile is blank, the tile is divided into 4 quadrants by channels
*/
func run2WithBFS(input string, steps int) int {
	b := ParseBoard(parsers.SplitByLines(input))
	reachables := b.bfs2(Pt(b.startX, b.startY), steps)

	return reachables.Cardinality()
}

func (b Board) bfs2(start image.Point, maxSteps int) mapset.Set[image.Point] {

	queue := lane.NewQueue[step](step{start, maxSteps})
	seen := mapset.NewSet[image.Point]()
	reachables := mapset.NewSet[image.Point]()

	for {
		s, ok := queue.Dequeue()
		if !ok { // yesterday I did this the other way around and wrapped the code in `if ok {` but this is better
			break
		}

		// idea from Dazbo- if steps left are odd, we can't *finish* on this spot; even means we can.
		if s.stepsLeft >= 0 {
			if s.stepsLeft%2 == 0 { // we can return here
				reachables.Add(s.pt)
			}
			if s.stepsLeft > 0 { // where can we go now?
				s.stepsLeft--                            // take a step, but to where?
				ns := b.GetSquareNeighborsNoChecks(s.pt) // these are Points not Tiles, but it excludes off-board tiles
				for _, n := range ns {
					//fmt.Printf("seen %d reachables %d ns %d steps %v %d\n", seen.Cardinality(), reachables.Cardinality(), len(ns), s.pt, s.stepsLeft)
					if seen.Contains(n) || b.AtPointWrapped(n).Contents == '#' {
						continue
					}
					queue.Enqueue(step{n, s.stepsLeft})
					seen.Add(n)
				}
			}
		}
	}
	return reachables
}
