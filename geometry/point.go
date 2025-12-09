package geometry

// int version of "gopkg.in/karalabe/cookiejar.v2/geometry"
// also doesn't use pointers

import (
	"math"

	"github.com/sekullbe/advent/tools"
)

// Two dimensional point.
type Point2 struct {
	X, Y int
}
type Point = Point2

// Three dimensional point.
type Point3 struct {
	X, Y, Z int
}

// Allocates and returns a new 2D point.
func NewPoint2(x, y int) Point2 {
	return Point2{x, y}
}

// Allocates and returns a new 3D point.
func NewPoint3(x, y, z int) Point3 {
	return Point3{x, y, z}
}

// Calculates the distance between x and y.
func (x Point2) Dist(y Point2) float64 {
	return math.Sqrt(float64(x.DistSqr(y)))
}

// Calculates the distance between x and y.
func (x Point3) Dist(y Point3) float64 {
	return math.Sqrt(float64(x.DistSqr(y)))
}

// Calculates the squared distance between x and y.
func (x Point2) DistSqr(y Point2) int {
	dx := x.X - y.X
	dy := x.Y - y.Y
	return dx*dx + dy*dy
}

// Calculates the squared distance between x and y.
func (x Point3) DistSqr(y Point3) int {
	dx := x.X - y.X
	dy := x.Y - y.Y
	dz := x.Z - y.Z
	return dx*dx + dy*dy + dz*dz
}

// Returns whether two points are equal.
func (x Point2) Equal(y Point2) bool {
	return tools.AbsInt(x.X-y.X) == 0 && tools.AbsInt(x.Y-y.Y) == 0
}

// Returns whether two points are equal.
func (x Point3) Equal(y Point3) bool {
	return tools.AbsInt(x.X-y.X) == 0 && tools.AbsInt(x.Y-y.Y) == 0 && tools.AbsInt(x.Z-y.Z) == 0
}

func (p Point2) Add(q Point2) Point2 {
	return Point2{p.X + q.X, p.Y + q.Y}
}

func (p Point2) MovePoint2(dx, dy int) Point2 {
	return NewPoint2(p.X+dx, p.Y+dy)
}

func (p Point2) MovePoint2WithWrap(dx, dy, maxX, maxY int) Point2 {
	return NewPoint2(wrapmod(p.X+dx, maxX+1), wrapmod(p.Y+dy, maxY+1))
}

func wrapmod(a, b int) int {
	return (a%b + b) % b
}

func (p Point3) MovePoint3(dx, dy, dz int) Point3 {
	return NewPoint3(p.X+dx, p.Y+dy, p.Z+dz)
}

// offset is what you have to add to A to get to B
func CalculateOffsets(a, b Point2) (x, y int) {
	return b.X - a.X, b.Y - a.Y
}

// Area of a rectangle with these points in opposite corners
func Area(a, b Point2) int {
	return (tools.AbsInt((a.X - b.X)) + 1) * (tools.AbsInt((a.Y - b.Y)) + 1)
}
