package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"maps"
	"math"
	"slices"
)

//go:embed input.txt
var inputText string

func main() {
	r1, r2 := run(inputText)
	fmt.Printf("Magic number: %d\n", r1)
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", r2)
}

type State struct {
	Pos geometry.Point2
	Dir geometry.Point2
}
type QI struct { // QueueItem
	State State
	Cost  int
	Path  map[geometry.Point2]bool
}

// cribbed this from https://github.com/mnml/aoc/blob/main/2024/16/1.go so I could figure it out for myself
// integrated that with my own board

func run(input string) (int, int) {

	b := grid.ParseBoardString(input)
	b.Dir = grid.EAST
	start, err := b.Find('S')
	if err != nil {
		panic("can't find start")
	}
	end, err := b.Find('E')
	if err != nil {
		panic("can't find end")
	}
	_ = end

	dist := map[State]int{}
	queue := []QI{{
		// Dir is a point instead of an int because that makes it easy to compute cost
		// if the neighbor offset == Dir then it's a straight line, else it's a turn and move
		State: State{
			Pos: start,
			Dir: grid.Pt(1, 0), // EAST
		},
		Cost: 0,
		Path: map[geometry.Point2]bool{start: true},
	}}

	// part1: min score seen, part2: all points seen on sittable tiles
	part1, part2 := math.MaxInt, make(map[geometry.Point2]bool)
	// this is Dijkstra
	for len(queue) > 0 {
		slices.SortFunc(queue, func(a, b QI) int {
			return cmp.Compare(a.Cost, b.Cost)
		})
		i := queue[0]
		queue = queue[1:]

		if c, ok := dist[i.State]; ok && c < i.Cost {
			continue
		}
		dist[i.State] = i.Cost

		if b.AtPoint(i.State.Pos).Contents == 'E' && i.Cost <= part1 {
			part1 = i.Cost
			maps.Copy(part2, i.Path)
		}

		for dir, cost := range map[geometry.Point2]int{
			i.State.Dir:                     1,
			{-i.State.Dir.Y, i.State.Dir.X}: 1001,
			{i.State.Dir.Y, -i.State.Dir.X}: 1001,
		} {
			n := State{i.State.Pos.Add(dir), dir}
			if b.AtPoint(n.Pos).Contents == '#' {
				continue
			}
			path := maps.Clone(i.Path)
			path[n.Pos] = true
			queue = append(queue, QI{n, i.Cost + cost, path})
		}
	}
	return part1, len(part2)

}

func run2(input string) int {

	return 0
}
