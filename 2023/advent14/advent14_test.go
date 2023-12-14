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

func Test_run2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sample}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
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

func TestRollSouth(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample))
	b.rollSouth()
	b.printBoard()
}

func TestRollWest(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample))
	b.rollWest()
	b.printBoard()
}

func TestRollEast(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample))
	b.rollEast()
	b.printBoard()
}

func TestWE4(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample))

	b.rollWest()
	b.printBoard()
	b.rollEast()
	b.printBoard()

	b.rollWest()
	b.printBoard()
	b.rollEast()
	b.printBoard()
	/*
		b.rollWest()
		b.rollEast()
		b.printBoard()

		b.rollWest()
		b.rollEast()
		b.printBoard()

	*/

}

func TestCycle(t *testing.T) {
	b := parseBoard(parsers.SplitByLines(sample))
	b.cycle()
	buffer := bytes.Buffer{}
	b.fprintBoard(&buffer)
	got := strings.TrimSpace(buffer.String())
	want := strings.TrimSpace(`
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`)
	want = strings.ReplaceAll(want, ".", "·")
	if got != want {
		b.printBoard()
		t.Errorf("cycle 1 incorrect")
	}

	b.cycle()
	buffer = bytes.Buffer{}
	b.fprintBoard(&buffer)
	got = strings.TrimSpace(buffer.String())
	want = strings.TrimSpace(`
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`)
	want = strings.ReplaceAll(want, ".", "·")
	if got != want {
		b.printBoard()
		t.Errorf("cycle 2 incorrect")
	}

	b.cycle()
	buffer = bytes.Buffer{}
	b.fprintBoard(&buffer)
	got = strings.TrimSpace(buffer.String())
	want = strings.TrimSpace(`
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`)
	want = strings.ReplaceAll(want, ".", "·")
	if got != want {
		b.printBoard()
		t.Errorf("cycle 3 incorrect")
	}
}

func TestLongSouth(t *testing.T) {

	alley := `
.O.
...
...
...
...
...
...`

	b := parseBoard(parsers.SplitByLines(alley))
	b.rollSouth()
	b.resetEsts()
	b.rollNorth()
	b.resetEsts()
	b.rollSouth()
	b.printBoard()
}

func TestLongWest(t *testing.T) {

	alley := `
.........
........O
.........`

	b := parseBoard(parsers.SplitByLines(alley))
	b.rollWest()
	b.printBoard()

}

func TestLongEast(t *testing.T) {

	alley := `
.........
O........
.........`

	b := parseBoard(parsers.SplitByLines(alley))
	b.rollEast()
	b.printBoard()
	b.resetEsts()
	b.rollWest()
	b.printBoard()
	b.resetEsts()
	b.rollEast()
	b.printBoard()

}
