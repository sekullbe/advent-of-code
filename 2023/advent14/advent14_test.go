package main

import (
	"bytes"
	"github.com/sekullbe/advent/parsers"
	"strings"
	"testing"
)

const sampleAfterMove = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sample}, want: 136},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRollNorth(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample))
	b.rollNorth()

	buffer := bytes.Buffer{}
	b.fprintBoard(&buffer)
	got := strings.TrimSpace(buffer.String())
	want := strings.TrimSpace(sampleAfterMove)
	want = strings.ReplaceAll(want, ".", "·")

	if got != want {
		t.Errorf("moved board isn't correct")
		b.printBoard()
	}

}
