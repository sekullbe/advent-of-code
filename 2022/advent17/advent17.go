package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/tools"
	"image"
	"log"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 2022))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run1(inputText, 1_000_000_000_000))
}

const (
	FLAT = iota
	PLUS
	ELL
	TALL
	BLOCK
)

// widths are 0-6
const CAVERN_MAX_X = 6

type board struct {
	cave      map[image.Point]any
	maxheight int
	minheight int
	nextRock  int
}

// stateKey valid AFTER a tick has completed
type stateKey struct {
	peaks   [7]int
	windIdx int
	rockIdx int
}

type stateVal struct {
	height int
	rocks  int
}

var exists = struct{}{}

type rock []image.Point

func np(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

func pointLeft(p image.Point) image.Point {
	return np(p.X-1, p.Y)
}
func pointRight(p image.Point) image.Point {
	return np(p.X+1, p.Y)
}
func pointDown(p image.Point) image.Point {
	return np(p.X, p.Y-1)
}

func newBoard() board {
	return board{
		cave:      make(map[image.Point]any),
		maxheight: 0,
		minheight: 0,
		nextRock:  FLAT,
	}
}

func (b *board) makeNewRock() rock {
	rockbottom := b.maxheight + 3
	rockleft := 2
	var newRock rock
	switch b.nextRock {
	case FLAT:
		newRock = rock{np(rockleft, rockbottom),
			np(rockleft+1, rockbottom),
			np(rockleft+2, rockbottom),
			np(rockleft+3, rockbottom)}
	case PLUS:
		newRock = rock{np(rockleft, rockbottom+1),
			np(rockleft+1, rockbottom),
			np(rockleft+1, rockbottom+1),
			np(rockleft+1, rockbottom+2),
			np(rockleft+2, rockbottom+1)}
	case ELL:
		newRock = rock{np(rockleft, rockbottom),
			np(rockleft+1, rockbottom),
			np(rockleft+2, rockbottom),
			np(rockleft+2, rockbottom+1),
			np(rockleft+2, rockbottom+2)}
	case TALL:
		newRock = rock{np(rockleft, rockbottom),
			np(rockleft, rockbottom+1),
			np(rockleft, rockbottom+2),
			np(rockleft, rockbottom+3)}
	case BLOCK:
		newRock = rock{np(rockleft, rockbottom),
			np(rockleft+1, rockbottom),
			np(rockleft, rockbottom+1),
			np(rockleft+1, rockbottom+1)}
	}
	b.nextRock += 1
	if b.nextRock > BLOCK {
		b.nextRock = FLAT
	}
	return newRock
}

func (b board) rockLands(r rock) bool {
	// check every point of rock to see if it crosses board or bottom but NOT if it goes off the edge
	for _, point := range r {
		_, crash := b.cave[point]
		if crash || point.Y < 0 {
			return true
		}
	}
	return false
}

func rockHitsSide(r rock) bool {
	for _, point := range r {
		if point.X < 0 || point.X > CAVERN_MAX_X {
			return true
		}
	}
	return false
}

func (b board) printBoard(message string, r rock) {
	fmt.Println(message)
	for y := b.maxheight + 6; y >= b.minheight; y-- {
		fmt.Print("|")
		for x := 0; x < 7; x++ {
			p := image.Point{x, y}
			_, ok := b.cave[p]
			isRock := tools.Contains(r, p)
			if ok {
				fmt.Print("#")
			} else if isRock {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}

// moves return a new rock with different coordinates and do not check for collision
func movedDown(r rock) rock {
	var nr rock
	for _, p := range r {
		nr = append(nr, pointDown(p))
	}
	return nr
}

func movedLeft(r rock) rock {
	var nr rock
	for _, p := range r {
		nr = append(nr, pointLeft(p))
	}
	return nr
}

func movedRight(r rock) rock {
	var nr rock
	for _, p := range r {
		nr = append(nr, pointRight(p))
	}
	return nr
}

func (b *board) landRock(r rock) {
	for _, point := range r {
		b.cave[point] = exists
		if point.Y+1 > b.maxheight {
			b.maxheight = point.Y + 1
		}
	}
}

func run1(inputText string, dropRocks int) int {

	winds := strings.TrimSpace(inputText)
	b := newBoard()
	newStateKey := stateKey{
		peaks:   [7]int{0, 0, 0, 0, 0, 0, 0},
		windIdx: 0,
		rockIdx: FLAT,
	}
	needNewRock := true
	numMoves := 0
	var fallingRock rock
	states := make(map[stateKey]stateVal)
	doStateCheck := true
	heightFromCycles := 0
	for rockCount := 1; rockCount <= dropRocks; {
		if needNewRock {
			newStateKey.rockIdx = b.nextRock // add the rock we just added
			fallingRock = b.makeNewRock()
			//b.printBoard("new rock", fallingRock)
			needNewRock = false
		}
		// read a rune from the input
		windRune := winds[numMoves%len(winds)]
		newStateKey.windIdx = numMoves % len(winds)
		movedRock := fallingRock
		switch windRune {
		case '<':
			movedRock = movedLeft(fallingRock)
		case '>':
			movedRock = movedRight(fallingRock)
		default:
			log.Panicf("what the heck wind rune is %c", windRune)
		}
		// if the fallingRock hits the side or bottom, it doesn't move, else it does
		if !(rockHitsSide(movedRock) || b.rockLands(movedRock)) {
			fallingRock = movedRock
		}
		//b.printBoard("just moved "+string(windRune), fallingRock)
		// Now move it down and see if it lands
		movedRock = movedDown(fallingRock)
		if b.rockLands(movedRock) { // if it would overlap, it lands where it is
			b.landRock(fallingRock)
			needNewRock = true
			//b.printBoard("landed", fallingRock)
			if doStateCheck {
				peaks := b.computeHeightmap()
				newStateKey.peaks = peaks
				sv, ok := states[newStateKey]
				if ok {
					// we've found a cycle between now and sv(StateValue)
					log.Printf("newStateKey match after %d rocks, maxheight %d, prev height %d, delta %d", rockCount, b.maxheight, sv.height, b.maxheight-sv.height)
					cycleLength := rockCount - sv.rocks                     // rocks per cycle
					heightPerCycle := b.maxheight - sv.height               // how much does each cycle grow the stack
					rocksLeft := dropRocks - rockCount                      // how many rocks do we have left to drop
					cyclesRemaining := rocksLeft / cycleLength              // how many cycles until we're done (drop fractions)
					heightFromCycles = heightPerCycle * cyclesRemaining     // how much height will all those cycles add
					rockCount = rockCount + (cyclesRemaining * cycleLength) // how many rocks left to drop to complete the drops
					doStateCheck = false                                    // we know the cycle, stop checking
					// could reset the grid, but instead just continue as is, and add in the computed height change at the end
					// the maxheight there will be height from before cycling and height from afer cycling
					// i.e. .......[BUNCHA CYCLES].......done
				}
				states[newStateKey] = stateVal{
					height: b.maxheight,
					rocks:  rockCount,
				}
			}
			rockCount++

		} else {
			fallingRock = movedRock
		}
		numMoves++
	}
	return b.maxheight + heightFromCycles
}

func (b *board) computeHeightmap() (heights [7]int) {
	// compute heightmap
	// this is dist from maxheight to the highest y in each column
	for x := 0; x < 7; x++ {
		for y := b.maxheight; y >= 0; y-- {
			_, ok := b.cave[np(x, y)]
			if ok {
				heights[x] = b.maxheight - y
				break
			}
		}
	}
	return
}
