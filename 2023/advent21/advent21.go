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
	fmt.Printf("Magic number: %d\n", run2(inputText))
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

func run2(input string) int {

	return 0
}
