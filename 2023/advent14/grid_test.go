package main

import (
	"bytes"
	"github.com/sekullbe/advent/parsers"
	"strings"
	"testing"
)

const sample = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func Test_MoveRock(t *testing.T) {

	b := parseBoard(parsers.SplitByLines(sample))

	if b.At(0, 0).rock != ROUND {
		t.Errorf("wrong rock at point after parse")
	}
	if b.At(1, 0).rock != SPACE {
		t.Errorf("wrong space at point after parse")
	}

	if b.At(4, 3).rock != ROUND {
		t.Errorf("wrong rock at point before move")
	}

	// simple move
	err := b.moveRock(Pt(4, 3), Pt(4, 2))
	if err != nil {
		t.Errorf("moveRock wasn't supposed to error")
	}
	if b.At(4, 3).rock != SPACE {
		t.Errorf("rock didn't leave")
	}
	if b.At(4, 2).rock != ROUND {
		t.Errorf("rock didn't arrive")
	}

	// zero move
	err = b.moveRock(Pt(4, 2), Pt(4, 2))
	if err != nil {
		t.Errorf("moveRock wasn't supposed to error")
	}
	if b.At(4, 2).rock != ROUND {
		t.Errorf("rock didn't stay in place")
	}

	// blocked move
	err = b.moveRock(Pt(4, 3), Pt(0, 0))
	if err == nil {
		t.Errorf("moveRock was supposed to error")
	}
	if b.At(4, 2).rock != ROUND {
		t.Errorf("rock didn't stay in place")
	}

	// moving a #
	err = b.moveRock(Pt(4, 1), Pt(1, 0))
	if err == nil {
		t.Errorf("expected an error moving a cube")
	}

	// moving a space
	err = b.moveRock(Pt(2, 0), Pt(1, 0))
	if err == nil {
		t.Errorf("expected an error moving a space")
	}

}

func Test_printBoard(t *testing.T) {

	b := parseBoard(parsers.SplitByLines(sample))
	buffer := bytes.Buffer{}
	b.fprintBoard(&buffer)
	got := strings.TrimSpace(buffer.String())
	want := strings.TrimSpace(sample)
	want = strings.ReplaceAll(want, ".", "·")

	if got != want {
		t.Errorf("parsed and printed boards don't match")
	}

}
