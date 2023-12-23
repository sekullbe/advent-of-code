package main

import (
	mapset "github.com/deckarep/golang-set/v2"
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
