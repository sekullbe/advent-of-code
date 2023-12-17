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
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type state struct {
	//cost int
	pt     image.Point
	facing int // direction we came into this tile from
	steps  int
}

func run1(input string) int {

	// we still want our old friend the Grid
	// those playing along at home will notice that it's shifted a bit more towards a real module
	b := ParseBoard(parsers.SplitByLines(input))
	cost := b.path(state{pt: Pt(0, 0), facing: -1, steps: 0})
	return cost
}

func (b *Board) path(startState state) int {
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(startState, 0)
	endPt := Pt(b.MaxX, b.MaxY)

	seen := mapset.NewSet[state]()

	for !pq.Empty() {
		st, pri, _ := pq.Pop() // pri == cost

		if st.pt == endPt { // hooray we're done
			return pri
		}
		// have we been here before?
		if seen.Contains(st) {
			continue
		}

		seen.Add(st)
		possibleNextStates := []state{}
		if !ValidDir(st.facing) { //not limited by previous dir, which means we're just starting
			for dir := NORTH; dir <= WEST; dir += 2 {
				possibleNextStates = append(possibleNextStates, state{
					pt:     st.pt.Add(NeigborInDirection(st.pt, dir)),
					facing: dir,
					steps:  1,
				})
			}
		} else {
			// turn left
			possibleNextStates = append(possibleNextStates, state{
				pt:     NeigborInDirection(st.pt, CounterClockwise(st.facing, 2)),
				facing: CounterClockwise(st.facing, 2),
				steps:  1,
			})
			// turn right
			possibleNextStates = append(possibleNextStates, state{
				pt:     NeigborInDirection(st.pt, Clockwise(st.facing, 2)),
				facing: Clockwise(st.facing, 2),
				steps:  1,
			})
			// go straight if we can
			if st.steps < 3 {
				possibleNextStates = append(possibleNextStates, state{
					pt:     NeigborInDirection(st.pt, st.facing),
					facing: st.facing,
					steps:  st.steps + 1,
				})
			}
		}
		for _, nextState := range possibleNextStates {
			if b.InRange(nextState.pt) {
				pq.Push(nextState, pri+b.AtPoint(nextState.pt).Contents)
			}
		}

	}
	panic("no solution")
}

func run2(input string) int {

	return 0
}
