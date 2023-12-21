package main

import (
	"bytes"
	"github.com/sekullbe/advent/parsers"
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
