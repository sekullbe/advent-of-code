package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"image"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 2000000))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type sensor struct {
	location image.Point
	nearest  image.Point
}

type beacon struct {
	location image.Point
}

type world struct {
	// should this contain a sensor OR a beacon? combined object? or two grids
	sensors                []sensor
	beacons                []beacon
	grid                   map[image.Point]tile
	minX, maxX, minY, maxY int
}
type tile struct {
	hasSensor       bool
	hasBeacon       bool
	inNoBeaconRange bool
}

func newWorld() world {
	return world{
		grid:    make(map[image.Point]tile),
		sensors: []sensor{},
		beacons: []beacon{},
	}
}

func (w *world) parseSensors(lines []string) {
	for _, line := range lines {
		var sx, sy, bx, by int

		//                         Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		n, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil || n != 4 {
			log.Panicf("parse error, got %d/4, error:%s", n, err)
		}
		s := sensor{location: image.Point{sx, sy}, nearest: image.Point{bx, by}}
		b := beacon{location: image.Point{bx, by}}
		w.grid[s.location] = tile{hasSensor: true}
		w.grid[b.location] = tile{hasBeacon: true}
		w.sensors = append(w.sensors, s)
		w.beacons = append(w.beacons, b)
		if w.minX > s.location.X {
			w.minX = s.location.X
		}
		if w.minX > b.location.X {
			w.minX = b.location.X
		}
		if w.minY > s.location.Y {
			w.minY = s.location.Y
		}
		if w.minY > b.location.Y {
			w.minY = b.location.Y
		}
		if w.maxX < s.location.X {
			w.maxX = s.location.X
		}
		if w.maxX < b.location.X {
			w.maxX = b.location.X
		}
		if w.maxY < s.location.Y {
			w.maxY = s.location.Y
		}
		if w.maxY < b.location.Y {
			w.maxY = b.location.Y
		}
	}
}
func rectangularDistance(p1, p2 image.Point) int {
	return tools.AbsInt(p1.X-p2.X) + tools.AbsInt(p1.Y-p2.Y)
}

func (w *world) computeNoBeaconLocations() {
	for _, s := range w.sensors {
		// Measure distance from s to its closest beacon
		d := rectangularDistance(s.location, s.nearest)
		// mark every location closer than s in the grid as true
		// look at every point x| p.X-d -> p.X+D, y | p.Y-d -> p.Y+d
		// if its dist <= d, mark it true
		for y := s.location.Y - d; y <= s.location.Y+d; y++ {
			for x := s.location.X - d; x <= s.location.X+d; x++ {
				td := rectangularDistance(image.Point{X: x, Y: y}, s.location)
				if td <= d {
					t, ok := w.grid[image.Point{X: x, Y: y}]
					if ok {
						t.inNoBeaconRange = true
					} else {
						w.grid[image.Point{X: x, Y: y}] = tile{inNoBeaconRange: true}
					}
				}
			}
		}
	}
}

func run1(inputText string, yToCheck int) int {
	w := newWorld()
	w.parseSensors(parsers.SplitByLines(inputText))
	w.computeNoBeaconLocations()
	// look at all grid points where x in minx->maxX and y=yToCheck
	noBeaconCount := 0
	for x := w.minX; x <= w.maxX; x++ {
		p := image.Point{x, yToCheck}
		t, ok := w.grid[p]
		if ok && t.inNoBeaconRange && !t.hasBeacon {
			noBeaconCount++
		}

	}

	return noBeaconCount
}

func run2(inputText string) int {

	return 0
}
