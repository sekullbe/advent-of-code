package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const sample = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

const sample2 = `
#.######
#.....v#
#.####.#
#.#>...#
#...##.#
######.#`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				input: sample,
			},
			want: 94,
		},
		{
			name: "sample2",
			args: args{
				input: sample2,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run1(tt.args.input), "run1(%v)", tt.args.input)
		})
	}
}

func TestSetComparable(t *testing.T) {

	a := mapset.NewSet[int](1, 2, 3)
	b := mapset.NewSet[int](1, 2, 3)
	assert.Equal(t, a, b)
	type f struct {
		set mapset.Set[int]
	}
	sa := f{a}
	sb := f{b}
	assert.Equal(t, sa, sb)

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
		{
			name: "sample",
			args: args{
				input: sample,
			},
			want: 154,
		},
		{
			name: "sample2",
			args: args{
				input: sample2,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run2(tt.args.input), "run2(%v)", tt.args.input)
		})
	}
}

func TestConjunctionCount(t *testing.T) {
	lines := parsers.SplitByLines(inputText)
	b := ParseBoard(lines)
	for point, tile := range b.Grid {
		ns := b.GetClearNeighbors(point)
		if len(ns) > 2 && tile.Contents == EMPTY {
			for i := 0; i < len(ns); i++ {
				if b.AtPoint(ns[i]).Contents == EMPTY {
					t.Errorf("crap, point %v not surrounded by ramps\n", point)
				}
			}
		}
	}
}
