package main

import (
	"bytes"
	"github.com/sekullbe/advent/parsers"
	"strings"
	"testing"
)

const sample = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func Test_printBoard(t *testing.T) {

	b := ParseBoard(parsers.SplitByLines(sample))
	buffer := bytes.Buffer{}
	b.FprintBoard(&buffer)
	got := strings.TrimSpace(buffer.String())
	want := strings.TrimSpace(sample)
	want = strings.ReplaceAll(want, ".", "·")

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
