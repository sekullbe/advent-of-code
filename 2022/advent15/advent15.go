package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"image"
	"log"
	"math"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1_countingOnly(inputText, 2000000))
	fmt.Printf("Magic number: %d\n", run1_withGrid(inputText, 2000000))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type sensor struct {
	location image.Point
	nearest  image.Point
	distance int // distance to the nearest beacon
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
		maxX:    math.MinInt,
		maxY:    math.MinInt,
		minY:    math.MaxInt,
		minX:    math.MaxInt,
	}
}

func (w *world) parseSensors(lines []string) {
	for _, line := range lines {
		var sx, sy, bx, by int
		n, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil || n != 4 {
			log.Panicf("parse error, got %d/4, error:%s", n, err)
		}
		s := sensor{location: image.Point{sx, sy}, nearest: image.Point{bx, by}}
		b := beacon{location: image.Point{bx, by}} // may not need to track those, but keep for now
		s.distance = rectangularDistance(s.location, b.location)
		w.minX = tools.MinInt(w.minX, s.location.X-s.distance+1)
		w.maxX = tools.MaxInt(w.maxX, s.location.X+s.distance+1)
		w.sensors = append(w.sensors, s)
		w.beacons = append(w.beacons, b)
	}
}

func (w *world) parseSensorsToGrid(lines []string) {
	for _, line := range lines {
		var sx, sy, bx, by int
		n, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if err != nil || n != 4 {
			log.Panicf("parse error, got %d/4, error:%s", n, err)
		}
		s := sensor{location: image.Point{sx, sy}, nearest: image.Point{bx, by}}
		b := beacon{location: image.Point{bx, by}} // may not need to track those, but keep for now
		s.distance = rectangularDistance(s.location, b.location)
		w.minX = tools.MinInt(w.minX, s.location.X-s.distance+1)
		w.maxX = tools.MaxInt(w.maxX, s.location.X+s.distance+1)
		w.sensors = append(w.sensors, s)
		w.grid[s.location] = tile{hasSensor: true}
		w.grid[b.location] = tile{hasBeacon: true}
	}
}

func rectangularDistance(p1, p2 image.Point) int {
	return tools.AbsInt(p1.X-p2.X) + tools.AbsInt(p1.Y-p2.Y)
}

func (w *world) computeNoBeaconLocations(yToCheck int) {
	for _, s := range w.sensors {
		// Measure distance from s to its closest beacon
		d := rectangularDistance(s.location, s.nearest)
		// mark every location closer than s in the grid as true
		// look at every point x| p.X-d -> p.X+D, y | p.Y-d -> p.Y+d
		// if its dist <= d, mark it true
		sY := s.location.Y
		if sY-d <= yToCheck && sY+d >= yToCheck {
			//for y := yToCheck; y <= yToCheck; y++ {
			for x := s.location.X - d; x <= s.location.X+d; x++ {
				td := rectangularDistance(image.Point{X: x, Y: yToCheck}, s.location)
				if td <= d {
					t, ok := w.grid[image.Point{X: x, Y: yToCheck}]
					if ok {
						t.inNoBeaconRange = true
					} else {
						w.grid[image.Point{X: x, Y: yToCheck}] = tile{inNoBeaconRange: true}
					}
				}
			}
		}
	}
}

// turns out after the fact that this works too; the problem was min/max X being wrong
func run1_withGrid(inputText string, yToCheck int) int {
	w := newWorld()
	w.parseSensorsToGrid(parsers.SplitByLines(inputText))
	w.computeNoBeaconLocations(yToCheck)
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

// this is much much faster
func run1_countingOnly(inputText string, yToCheck int) (noBeaconCount int) {
	w := newWorld()
	w.parseSensors(parsers.SplitByLines(inputText))

	for x := w.minX; x < w.maxX; x++ {
		// for each sensor, see if this loc is in its no-beacon distance
		// but also that it's not actually a beacon
		foundone := false
		for _, s := range w.sensors {
			p := image.Point{x, yToCheck}
			d := rectangularDistance(p, s.location)
			if p != s.nearest && d <= s.distance {
				noBeaconCount++
				foundone = true
				break
			}
		}
		if foundone {
			continue // maybe i should use a label, but this works
		}
	}

	return noBeaconCount
}

func run2(inputText string) int {

	return 0
}
