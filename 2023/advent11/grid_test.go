package main

import (
	"github.com/sekullbe/advent/parsers"
	"testing"
)

const sample1 = `.....
.S-7.
.|.|.
.L-J.
.....`

const sample1b = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const sample2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const sample2b = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

// isn't really a unit test... could be one by injecting a different writer into printBoard and using Fprintf
func Test_printBoards(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample1))
	b.printBoard()
	b = parseBoard(parsers.SplitByLines(sample1b))
	b.printBoard()
	b = parseBoard(parsers.SplitByLines(sample2))
	b.printBoard()
	b = parseBoard(parsers.SplitByLines(sample2b))
	b.printBoard()
	b = parseBoard(parsers.SplitByLines(inputText))
	b.printBoard()
}
