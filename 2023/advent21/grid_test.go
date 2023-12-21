package main

import (
	"bytes"
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const sample = `
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func Test_printBoard(t *testing.T) {

	b := ParseBoard(parsers.SplitByLines(sample))
	buffer := bytes.Buffer{}
	b.FprintBoard(&buffer)
	//b.printBoard()
	got := strings.TrimSpace(buffer.String())
	want := strings.TrimSpace(sample)
	want = strings.ReplaceAll(want, ".", "·")
	// because we parse out the S and replace it with empty
	want = strings.ReplaceAll(want, "S", "·")

	if got != want {
		t.Errorf("parsed and printed boards don't match")
	}

}

func TestClockwise(t *testing.T) {
	type args struct {
		dir   int
		ticks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "n to ne", args: args{dir: NORTH, ticks: 1}, want: NORTHEAST},
		{name: "nw to n", args: args{dir: NORTHWEST, ticks: 1}, want: NORTH},
		{name: "nw to ne", args: args{dir: NORTHWEST, ticks: 2}, want: NORTHEAST},
		{name: "full loop", args: args{dir: NORTHWEST, ticks: 8}, want: NORTHWEST},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clockwise(tt.args.dir, tt.args.ticks); got != tt.want {
				t.Errorf("Clockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounterClockwise(t *testing.T) {
	type args struct {
		dir   int
		ticks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "n to ne", args: args{dir: NORTH, ticks: 7}, want: NORTHEAST},
		{name: "nw to n", args: args{dir: NORTHWEST, ticks: 7}, want: NORTH},
		{name: "nw to ne", args: args{dir: NORTHWEST, ticks: 6}, want: NORTHEAST},
		{name: "full loop", args: args{dir: NORTHWEST, ticks: 8}, want: NORTHWEST},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CounterClockwise(tt.args.dir, tt.args.ticks); got != tt.want {
				t.Errorf("CounterClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_AtPointWrapped(t *testing.T) {
	b := ParseBoard(parsers.SplitByLines(sample))
	//b.printBoard()
	assert.Equal(t, '#', b.AtPointWrapped(Pt(5, 12)).Contents)
	assert.Equal(t, '#', b.AtPointWrapped(Pt(5, -2)).Contents)
	assert.Equal(t, '#', b.AtPointWrapped(Pt(5, -13)).Contents)
	assert.Equal(t, '#', b.AtPointWrapped(Pt(5, 9)).Contents)
	assert.Equal(t, '#', b.AtPoint(Pt(5, 9)).Contents)
	assert.Equal(t, '.', b.AtPoint(Pt(5, 5)).Contents)
	assert.Equal(t, '#', b.AtPoint(Pt(5, 6)).Contents)
	assert.Equal(t, '.', b.AtPoint(Pt(5, 8)).Contents)
	assert.Equal(t, '.', b.AtPointWrapped(Pt(5, 8)).Contents)
	assert.Equal(t, '.', b.AtPointWrapped(Pt(5, -3)).Contents)

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			assert.Equal(t, b.AtPoint(Pt(x%11, y%11)), b.AtPointWrapped(Pt(x, y)))
		}

	}
}

func TestBoard_GetSquareNeighborsNoChecks(t *testing.T) {
	b := ParseBoard(parsers.SplitByLines(sample))
	ns := b.GetSquareNeighborsNoChecks(Pt(10000, -10000))
	assert.Equal(t, 4, len(ns))

}

func Test_wrapmod(t *testing.T) {

	assert.Equal(t, 63, wrapmod(-1, 64))
	assert.Equal(t, 1, wrapmod(-63, 64))
	assert.Equal(t, 0, wrapmod(-64, 64))
	assert.Equal(t, 63, wrapmod(-65, 64))
	assert.Equal(t, 0, wrapmod(-128, 64))
	assert.Equal(t, 63, wrapmod(-129, 64))
	assert.Equal(t, 0, wrapmod(0, 64))
	assert.Equal(t, 1, wrapmod(1, 64))
	assert.Equal(t, 1, wrapmod(65, 64))
	assert.Equal(t, 1, wrapmod(129, 64))

}
