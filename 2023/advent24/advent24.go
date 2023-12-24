package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/xy/lineintersection"
	"github.com/twpayne/go-geom/xy/lineintersector"
	"math"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 200000000000000, 400000000000000))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string, minCoord, maxCoord int) int {
	hailstones := []Hailstone{}
	for i, line := range parsers.SplitByLines(input) {
		hailstones = append(hailstones, parseHailstone2D(line, i))
	}
	// for each pair of hailstones (ugh, O(N^2))
	// extend the line segment and see if they intersect and the intersection is in the area
	count := 0
	for _, h1 := range hailstones {
		for _, h2 := range hailstones {
			if h1.id <= h2.id {
				continue
			}
			t := h1.howLongToProject(minCoord, maxCoord)
			ip, ok := intersects2D(h1, h2, t)
			if !ok {
				continue // no intersection
			}
			fmt.Printf("%d intersects %d at (%f,%f)", h1.id, h2.id, ip.X(), ip.Y())
			if ip.X() >= float64(minCoord) && ip.Y() >= float64(minCoord) && ip.X() <= float64(maxCoord) && ip.Y() <= float64(maxCoord) {
				fmt.Println(" inside")
				count++
			} else {
				fmt.Println(" outside")
			}
		}
	}
	return count
}

func run2(input string) int {
	// so there's some line that at some T intersects every single hailstone
	// it'll have to be identical to the loc of some stone at that time,

	return 0
}

type Hailstone struct {
	id  int
	pos geom.Coord
	vel geom.Coord
}

func parseHailstone3D(line string, id int) Hailstone {
	// 19, 13, 30 @ -2,  1, -2
	lineParts := strings.Split(line, "@")
	coords := parsers.StringsWithCommasToIntSlice(lineParts[0])
	velocity := parsers.StringsWithCommasToIntSlice(lineParts[1])
	return Hailstone{
		id:  id,
		pos: toCoord3(coords[0], coords[1], coords[2]),
		vel: toCoord3(velocity[0], velocity[1], velocity[2])}
}

func parseHailstone2D(line string, id int) Hailstone {
	// 19, 13, 30 @ -2,  1, -2
	lineParts := strings.Split(line, "@")
	coords := parsers.StringsWithCommasToIntSlice(lineParts[0])
	velocity := parsers.StringsWithCommasToIntSlice(lineParts[1])
	return Hailstone{
		id:  id,
		pos: toCoord2(coords[0], coords[1]),
		vel: toCoord2(velocity[0], velocity[1])}
}

// github.com/twpayne/go-geom library uses float64
// under the hood a Coord is just a []float64
func toCoord2(x, y int) geom.Coord {
	return geom.Coord{float64(x), float64(y)}
}
func toCoord3(x, y, z int) geom.Coord {
	return geom.Coord{float64(x), float64(y), float64(z)}
}

func addCoords2D(a, b geom.Coord) geom.Coord {
	return geom.Coord{a.X() + b.X(), a.Y() + b.Y()}
}
func scaleCoordinate2D(a geom.Coord, factor int) geom.Coord {
	return geom.Coord{a.X() * float64(factor), a.Y() * float64(factor)}
}

func addCoords3D(a, b geom.Coord) geom.Coord {
	return geom.Coord{a.X() + b.X(), a.Y() + b.Y(), a[2] + b[2]}
}
func scaleCoordinate3D(a geom.Coord, factor int) geom.Coord {
	return geom.Coord{a.X() * float64(factor), a.Y() * float64(factor), a[2] * float64(factor)}
}

// where will this hailstone be at time t given its starting position and velocity?
func (h Hailstone) project2D(time int) geom.Coord {
	return addCoords2D(h.pos, scaleCoordinate2D(h.vel, time))
}

// compute how long we need to project2D a stone such that its x and y coords go through the target zone
// there don't appear to be any with 0 X or Y velocity
func (h Hailstone) howLongToProject(minCoord, maxCoord int) int {

	maxC := float64(maxCoord)
	minC := float64(minCoord)
	xt := max(math.Abs(minC-h.pos.X()), math.Abs(maxC-h.pos.X()))
	yt := max(math.Abs(minC-h.pos.Y()), math.Abs(maxC-h.pos.Y()))
	return int(max(xt, yt))

}

func intersects2D(h1, h2 Hailstone, t int) (intersectionPoint geom.Coord, ok bool) {
	res := lineintersector.LineIntersectsLine(lineintersector.NonRobustLineIntersector{}, h1.pos, h1.project2D(t), h2.pos, h2.project2D(t))
	if res.Type() != lineintersection.PointIntersection {
		return nil, false
	}
	return res.Intersection()[0], true
}
