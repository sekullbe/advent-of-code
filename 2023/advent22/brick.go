package main

import (
	"github.com/sekullbe/advent/geometry"
)

type Brick struct {
	End1   geometry.Point3
	End2   geometry.Point3
	zapped bool
}

const (
	X_AXIS = iota
	Y_AXIS
	Z_AXIS
)

// MaxInAxis returns the highest value in specified axis
func (b Brick) MaxInAxis(axis int) int {
	switch axis {
	case X_AXIS:
		return max(b.End1.X, b.End2.X)
	case Y_AXIS:
		return max(b.End1.Y, b.End2.Y)
	case Z_AXIS:
		return max(b.End1.Z, b.End2.Z)
	}
	return 0 // shouldn't happen
}

// MinInAxis returns the lowest value in specified axis
func (b Brick) MinInAxis(axis int) int {
	switch axis {
	case X_AXIS:
		return min(b.End1.X, b.End2.X)
	case Y_AXIS:
		return min(b.End1.Y, b.End2.Y)
	case Z_AXIS:
		return min(b.End1.Z, b.End2.Z)
	}
	return 0 // shouldn't happen
}

// Fall falls until its lowest point is on the new Z level
func (b *Brick) Fall(targetZ int) {
	distToFall := b.MinInAxis(Z_AXIS) - targetZ
	b.End1.Z -= distToFall
	b.End2.Z -= distToFall
}

func (b *Brick) FallBy(levels int) {
	b.Fall(b.MinInAxis(Z_AXIS) - levels)
}

// AxisLength returns min  &max endpoints in any axis
func (b Brick) AxisLength(axis int) (int, int) {
	return b.MinInAxis(axis), b.MaxInAxis(axis)
}

// returns the area of this broick
func (b Brick) PlaneArea() (geometry.Point2, geometry.Point2) {
	return geometry.NewPoint2(b.MinInAxis(X_AXIS), b.MinInAxis(Y_AXIS)), geometry.NewPoint2(b.MaxInAxis(X_AXIS), b.MaxInAxis(Y_AXIS))
}

/*
   def xy_area(self):
       """ Return the x,y area occupied by this brick, as a tuple of ((corner),(corner)) """
       return ((self.min_for_axis(0), (self.min_for_axis(1))),
               (self.max_for_axis(0), (self.max_for_axis(1))))
*/
