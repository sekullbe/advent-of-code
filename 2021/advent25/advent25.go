package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

const LOG = false

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
}

const EAST rune = '>'
const SOUTH rune = 'v'
const EMPTY rune = '.'

type point struct {
	row int
	col int
}

type square struct {
	contents rune
	next     rune
}

type board struct {
	grid map[point]square
	rows int
	cols int
}

func run1(inputText string) int {
	board := parseGrid(inputText)
	steps := 0
	if LOG {
		fmt.Printf("after %d steps\n", steps)
		fmt.Println(board.toString())
	}
	for moveCount := 1; moveCount > 0; {
		moveCount = board.step()
		steps++
		if LOG {
			fmt.Printf("after %d steps\n", steps)
			fmt.Println(board.toString())
			log.Printf("moveCount: %d", moveCount)
		}

	}
	return steps
}

func parseGrid(input string) *board {
	board := board{}
	board.grid = make(map[point]square)
	rows := parsers.SplitByLines(input)
	board.rows = len(rows)
	board.cols = len(rows[1])
	for ridx, row := range rows {
		for cidx, col := range row {
			p := point{row: ridx, col: cidx}
			board.grid[p] = square{contents: col, next: col}
		}
	}
	return &board
}

func (b *board) step() (moveCount int) {
	// need to do the scan twice, but can do the finalizing pass only once
	for pass := 1; pass <= 3; pass++ {
		for r := 0; r < b.rows; r++ {
			for c := 0; c < b.cols; c++ {
				p := point{row: r, col: c}
				currentsq := b.grid[p]
				neighborsq, neighborPoint := b.look(p) // this is the square the current cucumber is looking at, if any
				if pass == 1 {                         // East-movers move
					if currentsq.contents == EAST { // if you want to move east
						if neighborsq.contents == EMPTY { // and you can
							currentsq.next = EMPTY // then move
							neighborsq.next = currentsq.contents
							b.grid[p] = currentsq // put the modified square back into the grid- can't modify b.grid[p].contents directly
							b.grid[neighborPoint] = neighborsq
							moveCount++
						}
					}
				} else if pass == 2 { // South-movers move

					if currentsq.contents == SOUTH { // if you want to move south
						if neighborsq.contents != SOUTH && neighborsq.next == EMPTY { // no southmover in your way and no pesky eastmover has moved into that spot
							currentsq.next = EMPTY // then move
							neighborsq.next = currentsq.contents
							b.grid[p] = currentsq
							b.grid[neighborPoint] = neighborsq
							moveCount++
						}
					}
				} else { // Copy all 'nexts' to 'current'
					currentsq.contents = currentsq.next
					b.grid[p] = currentsq
				}
			}
		}
	}
	return
}

func (b board) look(p point) (square, point) {
	switch b.grid[p].contents {
	case SOUTH:
		return b.lookSouth(p)
	case EAST:
		return b.lookEast(p)
	default:
		return square{contents: EMPTY, next: EMPTY}, p
	}
}
func (b board) lookEast(p point) (square, point) {
	newCol := (p.col + 1) % b.cols
	np := point{row: p.row, col: newCol}
	return b.grid[np], np
}
func (b board) lookSouth(p point) (square, point) {
	newRow := (p.row + 1) % b.rows
	np := point{row: newRow, col: p.col}
	return b.grid[np], np
}

func (b board) toString() (s string) {
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.cols; c++ {
			p := point{row: r, col: c}
			s += string(b.grid[p].contents)
		}
		s += "\n"
	}
	return
}
