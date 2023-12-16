package main

import (
	"bytes"
	"github.com/sekullbe/advent/parsers"
	"strings"
	"testing"
)

const sample = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

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
