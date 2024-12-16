package main

import (
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const sample1 = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

const sample2 = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const sample3 = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: sample1}, want: 2028},
		{name: "sample2", args: args{input: sample2}, want: 10092},
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
		{name: "sample3", args: args{input: sample3}, want: 618},
		{name: "sample2", args: args{input: sample2}, want: 9021},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveRobotWithWideBoxes_horizontally(t *testing.T) {

	b := grid.ParseBoardString(`#@[]...#`)
	b.PrintBoard()
	nr := moveRobotWithWideBoxes(b, geometry.NewPoint2(1, 0), grid.EAST)
	assert.Equal(t, geometry.NewPoint2(2, 0), nr)
	b.PrintBoard()

	nr = moveRobotWithWideBoxes(b, geometry.NewPoint2(2, 0), grid.EAST)
	assert.Equal(t, geometry.NewPoint2(3, 0), nr)
	b.PrintBoard()

	nr = moveRobotWithWideBoxes(b, geometry.NewPoint2(3, 0), grid.EAST)
	assert.Equal(t, geometry.NewPoint2(4, 0), nr)
	b.PrintBoard()

	// can't move any more
	nr = moveRobotWithWideBoxes(b, geometry.NewPoint2(4, 0), grid.EAST)
	assert.Equal(t, geometry.NewPoint2(4, 0), nr)
	assert.Equal(t, '[', b.At(5, 0).Contents)
	assert.Equal(t, ']', b.At(6, 0).Contents)

}

func Test_recursivelyMoveBoxVertically_noCommits(t *testing.T) {

	b := grid.ParseBoardString("....\n.[].\n....")
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 1), grid.NORTH, false))

	b = grid.ParseBoardString("####\n.[].\n....")
	assert.False(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 1), grid.NORTH, false))

	b = grid.ParseBoardString("####\n....\n.[].\n.[].\n....")
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 3), grid.NORTH, false))

	b = grid.ParseBoardString("####\n....\n[][]\n.[].\n....")
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 3), grid.NORTH, false))

	b = grid.ParseBoardString("####\n....\n[][]\n.[].\n....")
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(2, 3), grid.NORTH, false))

	b = grid.ParseBoardString("####\n[]..\n[][]\n.[].\n....")
	assert.False(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 3), grid.NORTH, false))

}

func Test_recursivelyMoveBoxVertically_withCommits(t *testing.T) {
	b := grid.ParseBoardString("####\n....\n.[].\n....")
	b.PrintBoard()
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 2), grid.NORTH, true))
	b.PrintBoard()

	fmt.Println("----------")

	b = grid.ParseBoardString("####\n....\n.[].\n.[].\n....")
	b.PrintBoard()
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 3), grid.NORTH, true))
	b.PrintBoard()

	b = grid.ParseBoardString("####\n....\n[][]\n.[].\n....")
	b.PrintBoard()
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(1, 3), grid.NORTH, true))
	b.PrintBoard()

	b = grid.ParseBoardString("####\n....\n[][]\n.[].\n....")
	b.PrintBoard()
	assert.True(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(2, 3), grid.NORTH, true))
	b.PrintBoard()

	// even though commit is true, nothing moves
	b = grid.ParseBoardString("####\n#...\n[][]\n.[].\n....")
	b.PrintBoard()
	assert.False(t, recursivelyMoveBoxVertically(b, geometry.NewPoint2(2, 3), grid.NORTH, true))
	b.PrintBoard()

}
