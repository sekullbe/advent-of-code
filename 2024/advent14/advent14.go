package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var inputText string

type robot struct {
	pos geometry.Point
	vel geometry.Point //point makes a fine vector
}

const (
	NE = iota
	SE
	SW
	NW
)

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 100, 102))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText, 100, 102))
}

func run1(input string, maxX int, maxY int) int {
	lines := parsers.SplitByLines(input)
	robots := []robot{}
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}
	ticks := 100
	for tick := 0; tick < ticks; tick++ {
		for i, r := range robots {
			r.moveRobot(maxX, maxY)
			robots[i] = r
		}
	}

	return computeDangerLevel(robots, maxX, maxY)
}

func computeDangerLevel(robots []robot, maxX int, maxY int) int {
	quads := [4]int{0, 0, 0, 0}
	for _, r := range robots {
		q := computeQuadrant(r.pos, maxX, maxY)
		if q >= 0 {
			quads[q]++
		}

	}

	return quads[0] * quads[1] * quads[2] * quads[3]
}

func run2(input string, maxX int, maxY int) int {
	lines := parsers.SplitByLines(input)
	robots := []robot{}
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}
	display(robots, maxX, maxY)

	reader := bufio.NewReader(os.Stdin)

	maxdl := -1
	isNewMaxDl := true
	for i := 1; i < 10000000; i++ {
		dl := computeDangerLevel(robots, maxX, maxY)

		if dl < maxdl {
			maxdl = dl
			isNewMaxDl = true
		} else {
			isNewMaxDl = false
		}

		_ = isNewMaxDl
		for i, r := range robots {
			r.moveRobot(maxX, maxY)
			robots[i] = r
		}

		if i%(maxY+1)-65 == 0 || i%(maxX+1)-11 == 0 {
			display(robots, maxX, maxY)
			fmt.Printf("%d: dl=%d maxdl %v\n", i, dl, isNewMaxDl)
			reader.ReadString('\n')
		}
	}
	// hpat at 65,168 (i+103) vpat at 11,112 (i+101)
	// min and max danger level didn't help
	// at this point just enter through all the printed patterns
	return 0
}

func parseRobot(robLine string) robot {
	robot := robot{
		pos: geometry.Point{},
		vel: geometry.Point{},
	}
	_, err := fmt.Fscanf(strings.NewReader(robLine), "p=%d,%d v=%d,%d", &robot.pos.X, &robot.pos.Y, &robot.vel.X, &robot.vel.Y)
	if err != nil {
		log.Fatalf("parser faield: %s", robLine)
	}
	return robot
}

func (r *robot) moveRobot(maxX, maxY int) {
	r.pos = r.pos.MovePoint2WithWrap(r.vel.X, r.vel.Y, maxX, maxY)
}

func computeQuadrant(p geometry.Point2, maxX int, maxY int) int {
	if p.X < maxX/2 && p.Y < maxY/2 {
		return NW
	}
	if p.X > maxX/2 && p.Y > maxY/2 {
		return SE
	}
	if p.X < maxX/2 && p.Y > maxY/2 {
		return SW
	}
	if p.X > maxX/2 && p.Y < maxY/2 {
		return NE
	}
	return -1
}

func display(robots []robot, maxX int, maxY int) {
	b := grid.Board{
		Grid: make(grid.Grid),
		MaxX: maxX,
		MaxY: maxY,
	}
	for _, r := range robots {
		t := grid.NewTile(r.pos, '#')
		b.Grid[r.pos] = &t
	}

	b.PrintBoard()

}
