package main

import (
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xyz"
	"math"
	"testing"
)

const sample = `
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

// this finds *that* lines intersect but not *where* they intersect
func TestGeomDoesWhatIWant(t *testing.T) {

	l1s := geom.Coord{0, 0, 0}
	l1e := geom.Coord{10, 0, 0}
	l2s := geom.Coord{5, -5, 0}
	l2e := geom.Coord{5, 5, 0}
	dist := xyz.DistanceLineToLine(l1s, l1e, l2s, l2e)
	assert.Equal(t, 0.0, dist)

	l2s = geom.Coord{5, -5, 1}
	l2e = geom.Coord{5, 5, 1}
	dist = xyz.DistanceLineToLine(l1s, l1e, l2s, l2e)
	assert.Equal(t, 1.0, dist)

	l1s = geom.Coord{-5, -5, -5}
	l1e = geom.Coord{5, 5, 5}
	l2s = geom.Coord{-5, 5, 5}
	l2e = geom.Coord{5, -5, -5}
	dist = xyz.DistanceLineToLine(l1s, l1e, l2s, l2e)
	assert.Equal(t, 0.0, dist)

}

func TestHowBigIsMaxint(t *testing.T) {
	assert.True(t, math.MaxInt > 200000000000000)
	assert.True(t, math.MaxInt > 400000000000000)
}

func TestSmallestCoordinate(t *testing.T) {
	maxX := float64(math.MaxInt)
	maxY := float64(math.MaxInt)
	maxZ := float64(math.MaxInt)
	for _, s := range parsers.SplitByLines(inputText) {
		h := parseHailstone3D(s, 0)
		maxX = min(maxX, h.pos[0])
		maxY = min(maxY, h.pos[1])
		maxZ = min(maxZ, h.pos[2])
	}
	fmt.Printf("X: %d\nY: %d\nZ: %d\n", int(maxX), int(maxY), int(maxZ))
	fmt.Println("L: 200000000000000")
	fmt.Println("H: 400000000000000")
	// so the smallest coordinates are much smaller than the range we're looking for
}

func Test_run1(t *testing.T) {
	type args struct {
		input    string
		minCoord int
		maxCoord int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				input:    sample,
				minCoord: 7,
				maxCoord: 27,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run1(tt.args.input, tt.args.minCoord, tt.args.maxCoord), "run1(%v, %v, %v)", tt.args.input, tt.args.minCoord, tt.args.maxCoord)
		})
	}
}
