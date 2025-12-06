package grid

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/parsers"

	"github.com/stretchr/testify/assert"
)

const sample = `
#.######
#.....v#
#.####.#
#.#>...#
#...##.#
######.#`

func Test_isSymbol(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "blank", args: args{r: '.'}, want: false},
		{name: "number", args: args{r: '1'}, want: false},
		{name: "*", args: args{r: '*'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymbol(tt.args.r); got != tt.want {
				t.Errorf("isSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestBoard_SlideTile(t *testing.T) {

	b := ParseBoardString(`X.`)
	moved := b.SlideTile(geometry.NewPoint2(0, 0), EAST)
	ot := b.At(0, 0)
	nt := b.At(1, 0)
	assert.True(t, moved)
	assert.Equal(t, EMPTY, ot.Contents)
	assert.Equal(t, 'X', nt.Contents)

	moved = b.SlideTile(geometry.NewPoint2(1, 0), EAST)
	assert.False(t, moved)
	assert.Equal(t, EMPTY, ot.Contents)
	assert.Equal(t, 'X', nt.Contents)

}

func TestDirFromTo(t *testing.T) {
	type args struct {
		from geometry.Point2
		to   geometry.Point2
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "N", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 1, Y: 0}}, want: NORTH},
		{name: "NE", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 2, Y: 0}}, want: NORTHEAST},
		{name: "E", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 2, Y: 1}}, want: EAST},
		{name: "SE", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 2, Y: 2}}, want: SOUTHEAST},
		{name: "S", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 1, Y: 2}}, want: SOUTH},
		{name: "SW", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 0, Y: 2}}, want: SOUTHWEST},
		{name: "W", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 0, Y: 1}}, want: WEST},
		{name: "NW", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 0, Y: 0}}, want: NORTHWEST},
		{name: "fail", args: args{from: geometry.Point2{X: 1, Y: 1}, to: geometry.Point2{X: 1, Y: 1}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DirFromTo(tt.args.from, tt.args.to), "DirFromTo(%v, %v)", tt.args.from, tt.args.to)
		})
	}
}

func TestPathToSteps(t *testing.T) {
	type args struct {
		path []geometry.Point2
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{name: "one step", args: args{[]geometry.Point2{{1, 1}, {1, 0}}}, want: []rune{'^'}},
		{name: "two steps", args: args{[]geometry.Point2{{1, 1}, {1, 0}, {2, 0}}}, want: []rune{'^', '>'}},
		{name: "three steps", args: args{[]geometry.Point2{{1, 1}, {1, 0}, {2, 0}, {1, 0}}}, want: []rune{'^', '>', '<'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PathToSteps(tt.args.path), "PathToSteps(%v)", tt.args.path)
		})
	}
}

func TestBoard_IsColumnEmpty(t *testing.T) {

	input := `x xx x
xx x x
x  x x`
	b := ParseBoardString(input)
	assert.False(t, b.IsColumnEmpty(0))
	assert.False(t, b.IsColumnEmpty(1))
	assert.False(t, b.IsColumnEmpty(2))
	assert.False(t, b.IsColumnEmpty(3))
	assert.True(t, b.IsColumnEmpty(4))
	assert.False(t, b.IsColumnEmpty(5))
}
