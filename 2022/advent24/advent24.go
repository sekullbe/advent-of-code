package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"image"
	"math"
	"math/bits"
)

//go:embed input.txt
var inputText string

/* I was totally out of clues here so I took the solution from
https://github.com/nklaassen/advent-of-code/tree/main/2022/24
and rewrote it in my own style to try to understand it.
Main change is converting the [][] grid to my usual map[image.Point]
and in a couple of places explicitly using rune instead of byte.
*/

const (
	maxTime = 1024
)

func np(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

func main() {

	lines := parsers.SplitByLines(inputText)
	state := newState()
	//state.valley[0] = make([][]BlizzardSet, len(lines))
	state.vm[0] = make(map[image.Point]BlizzardSet)
	for y, line := range lines {
		for x, r := range line {
			state.vm[0][np(x, y)] = BlizzardSet(parseBlizzard(r))
		}
	}
	// state contains a grid of walls & blizzards
	for i := 0; i < maxTime; i++ {
		state.computeNext(i)
	}
	start := SpaceTime{t: 0, Point: np(0, 1)}
	goal := np(120, 26)
	t1 := state.solve(start, goal)
	fmt.Println("part 1:", t1)

	start, goal = SpaceTime{t: t1, Point: goal}, start.Point
	t2 := state.solve(start, goal)
	//fmt.Println("back to start:", t2)
	start, goal = SpaceTime{t: t2, Point: np(goal.X, goal.Y)}, np(start.X, start.Y)
	t3 := state.solve(start, goal)
	fmt.Println("part 2:", t3)
}

type State struct {
	vm []map[image.Point]BlizzardSet
}

func newState() *State {
	return &State{
		vm: make([]map[image.Point]BlizzardSet, maxTime+1),
	}
}

type SpaceTime struct {
	image.Point
	t int
}

func (state *State) solve(start SpaceTime, goal image.Point) int {
	// bfs in spacetime
	seen := map[SpaceTime]bool{
		start: true,
	}
	q := []SpaceTime{start}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, offset := range []image.Point{np(0, -1), np(0, 1), np(-1, 0), np(1, 0), np(0, 0)} {
			next := SpaceTime{t: curr.t + 1, Point: np(curr.X+offset.X, curr.Y+offset.Y)}
			if next.Y < 0 || next.Y >= 27 {
				continue
			}
			if next.X == goal.X && next.Y == goal.Y {
				return next.t
			}
			//if state.valley[next.t][next.p.Y][next.p.X].empty() && !seen[next] {
			if state.vm[next.t][np(next.X, next.Y)].empty() && !seen[next] {
				seen[next] = true
				q = append(q, next)
			}
		}
	}
	return math.MaxInt
}

func (state *State) computeNext(t int) {
	state.vm[t+1] = make(map[image.Point]BlizzardSet)

	// Here I have to iterate over all points by coordinate, not all the key/value pairs in the
	// map. Go returns those in random order. I'm not sure if the problem was
	// iterating out of order or if there were missing points (even though I tried to
	// keep the map grid not sparse). Which kind of removes much of the point of a map-based grid...
	for y := 0; y < 27; y++ {
		for x := 0; x < 122; x++ {
			p := np(x, y)
			for _, b := range state.vm[t][p].blizzards() {
				//for _, b := range bset.blizzards() { // get the blizzards in that spot
				pn := state.moveBlizzard(b, p)
				state.vm[t+1][pn] = state.vm[t+1][pn] | BlizzardSet(b)
				// this won't work, because while you can call a pointer method
				// on an element of a slice, you cannot do the same on a member of a map.
				// see https://goplay.tools/snippet/sh7NXbylrTk
				//state.vm[t+1][pn].set(b)
			}
		}
	}
}

func (state *State) printAt(t int) {
	for y := 0; y < 27; y++ {
		for x := 0; x < 122; x++ {
			fmt.Print(state.vm[t][np(x, y)])
		}
		fmt.Println()
	}
}

func (state *State) moveBlizzard(b Blizzard, p image.Point) (newp image.Point) {
	switch b {
	case 1: // >
		newp = np(p.X+1, p.Y)
	case 2: // v
		newp = np(p.X, p.Y+1)
	case 4: // <
		newp = np(p.X-1, p.Y)
	case 8: // ^
		newp = np(p.X, p.Y-1)
	case 16: // #
		return np(p.X, p.Y)
	default:
		panic(b)
	}
	numRows := 27
	numCols := 122
	switch {
	case newp.Y == 0: // off the top
		newp.Y = numRows - 2
	case newp.Y == numRows-1: // off the bottom
		newp.Y = 1
	case newp.X == 0: // off the left
		newp.X = numCols - 2
	case newp.X == numCols-1: // off the right
		newp.X = 1
	}
	return newp
}

type Blizzard uint8

func (b Blizzard) String() string {
	switch b {
	case 1:
		return ">"
	case 2:
		return "v"
	case 4:
		return "<"
	case 8:
		return "^"
	case 16:
		return "#"
	}
	return " "
}

func parseBlizzard(b rune) Blizzard {
	switch b {
	case '.':
		return 0
	case '>':
		return 1
	case 'v':
		return 2
	case '<':
		return 4
	case '^':
		return 8
	case '#':
		return 16
	}
	panic(b)
}

type BlizzardSet uint8

func (b BlizzardSet) String() string {
	if pop := bits.OnesCount8(uint8(b)); pop <= 1 {
		return Blizzard(b).String()
	} else {
		return fmt.Sprint(pop)
	}
}

func (b BlizzardSet) blizzards() (blizzards []Blizzard) { // unwrap the bits in the uint8(BlizzardSet) into multiple uint8s(Blizzards) - I think I'd use a struct of bools?
	for bit := uint8(1); bit <= 16; bit <<= 1 {
		if uint8(b)&bit != 0 {
			blizzards = append(blizzards, Blizzard(bit))
		}
	}
	return
}

func (b *BlizzardSet) set(blizzard Blizzard) {
	*b |= BlizzardSet(blizzard)
}

func (b BlizzardSet) empty() bool {
	return uint8(b) == 0
}
