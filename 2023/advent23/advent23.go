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

	return 0
}

// this is a BFS; maybe I need DFS instead
func (b Board) pathBfs(startPt image.Point, endPt image.Point) int {
	pq := lane.NewQueue[state]()
	pq.Enqueue(state{pt: startPt, steps: 0})
	seen := mapset.NewSet[image.Point]() // steps too? or just the points?
	longestPath := 0

	for pq.Size() > 0 {
		st, _ := pq.Dequeue()

		if st.pt == endPt { // hooray we're done
			longestPath = max(longestPath, st.steps)
			//fmt.Printf("winnah! steps=%d\n", st.steps)
		}

		if seen.Contains(st.pt) {
			continue
		}
		seen.Add(st.pt)
		possibleNextStates := []state{}
		// if we're on an arrow we must follow it
		if arrowDir, isArrow := directionOfArrow(b.AtPoint(st.pt).Contents); isArrow {
			ntile := b.AtPoint((NeighborInDirection(st.pt, arrowDir)))
			if ntile.Contents != FOREST {
				//fmt.Printf("Looking at %v because arrow\n", ntile.Pt)
				nst := state{
					pt:    NeighborInDirection(st.pt, arrowDir),
					steps: st.steps + 1,
				}
				possibleNextStates = append(possibleNextStates, nst)
			}
		} else {
			for _, npt := range b.GetSquareNeighbors(st.pt) {
				ntile := b.AtPoint(npt)
				// do i need to handle the special case where you can't go into a slope pointing at you?
				if _, isArrow := directionOfArrow(ntile.Contents); isArrow {

				}
				if ntile.Contents != FOREST {
					possibleNextStates = append(possibleNextStates, state{
						pt:    npt,
						steps: st.steps + 1,
					})
				}
			}
		}
		for _, nextState := range possibleNextStates {
			if b.InRange(nextState.pt) {
				//fmt.Printf("Looking at %v\n", nextState.pt)
				pq.Enqueue(nextState)
			}
		}

	}
	return longestPath
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
			for _, npt := range b.GetSquareNeighbors(st.pt) {
				ntile := b.AtPoint(npt)
				// do i need to handle the special case where you can't go into a slope pointing at you?
				//if _, isArrow := directionOfArrow(ntile.Contents); isArrow {
				//
				//}
				var nseen mapset.Set[image.Point] = (*st.seen).Clone()
				if ntile.Contents != FOREST {
					possibleNextStates = append(possibleNextStates, state{
						pt:    npt,
						steps: st.steps + 1,
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

/*
read the map into a grid
make a graph of the grid by scanning each empty space
  add each neighbor as an edge
    handling the special case where you can't go back
    case to handle there is that you can get stuck if you come across a < pointing the way you came
  then do bfs to the exit

'yourbasic/graph' has a ShortestPath but not a LongestPath
it does have a BFS though
*/
