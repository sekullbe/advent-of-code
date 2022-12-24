package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"image"
	"math"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
	fmt.Println("-------------")
}

func parseElves(lines []string) *board {
	elfId := 1
	b := &board{
		grove:       make(map[image.Point]*elf),
		elves:       []*elf{},
		proposalDir: NORTH,
	}
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				loc := np(x, y)
				elf := elf{
					Point:         loc,
					proposalPoint: loc,
					board:         b,
					id:            elfId,
				}
				b.grove[loc] = &elf
				b.elves = append(b.elves, &elf)
				elfId++
			}

		}
	}
	return b
}

func (e *elf) decideProposal() image.Point {
	isNeighbor := false
	for dir := NORTH; dir <= NORTHWEST; dir++ {
		dp := neighborInDirection(e.Point, dir)
		_, ok := e.board.grove[dp]
		if ok {
			isNeighbor = true
			break
		}
	}
	if !isNeighbor {
		return e.Point // no neighbors, no proposal
	}

	pdir := e.board.proposalDir
	for scans := 0; scans < 4; scans++ {
		scanFoundElf := false
		for _, ldir := range scanlooks[pdir] {
			_, occupied := e.board.grove[neighborInDirection(e.Point, ldir)]
			if occupied { // don't propose moving this way
				scanFoundElf = true
				break
			}
		}
		if !scanFoundElf {
			return neighborInDirection(e.Point, pdir)
		}
		pdir = (pdir + 2) % 8 // because we're traversing the major points only
	}
	return e.Point
}

// returns true if actually moved
func (e *elf) move(newPoint image.Point) bool {
	if e.Point != newPoint {
		e.board.grove[newPoint] = e
		delete(e.board.grove, e.Point)
		e.Point = newPoint
		e.proposalPoint = e.Point
		return true
	}
	return false
}

func (b *board) round() int {
	proposals := make(map[image.Point][]*elf)
	for _, e := range b.elves {
		pp := e.decideProposal()
		e.proposalPoint = pp
		if e.Point != pp {
			proposals[pp] = append(proposals[pp], e)
		}
	}
	// 2nd half- look at all the proposals and if only one elf proposes to move to a point, move it
	moves := 0
	for proposedPoint, elves := range proposals {
		if len(elves) == 1 {
			moved := elves[0].move(proposedPoint)
			if moved {
				moves++
			}
		}
	}

	// rotate the proposal pointer for the next turn
	b.proposalDir = (b.proposalDir + 2) % 8
	//fmt.Printf("New scan dir: %d\n", b.proposalDir)
	return moves
}

func (b *board) score() int {
	minX, maxX, minY, maxY := b.getCorners()
	return ((maxX - minX + 1) * (maxY - minY + 1)) - len(b.elves)
}

func (b *board) getCorners() (minX, maxX, minY, maxY int) {
	maxX = math.MinInt
	maxY = math.MinInt
	minX = math.MaxInt
	minY = math.MaxInt
	for _, e := range b.elves {
		minX = tools.MinInt(minX, e.X)
		maxX = tools.MaxInt(maxX, e.X)
		minY = tools.MinInt(minY, e.Y)
		maxY = tools.MaxInt(maxY, e.Y)
	}
	return
}

func (b *board) print() {
	minX, maxX, minY, maxY := b.getCorners()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if elf, ok := b.grove[np(x, y)]; ok {
				_ = elf
				//fmt.Printf("%02d", elf.id)
				fmt.Printf("#")
			} else {
				//fmt.Print(".,")
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func run1(input string) int {
	b := parseElves(parsers.SplitByLines(input))
	//b.print()
	for round := 1; round <= 10; round++ {
		b.round()
		//fmt.Printf("After round %d\n", round)
		//b.print()
	}
	score := b.score()
	return score
}

func run2(input string) int {
	b := parseElves(parsers.SplitByLines(input))

	round := 1
	for ; b.round() > 0; round++ {
	}
	return round
}
