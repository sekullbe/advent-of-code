package main

import "image"

// cribbed and modified from
//https://github.com/GeertJohan/go.geofence
// it seems to automatically exclude the 'outside' points not in the loop

// original comment- this site is no longer available, but see
// https://web.archive.org/web/20130126163405/http://geomalgorithms.com/a03-_inclusion.html
// The Polyfence is based on information and explanation by softSurfer (Dan Sunday)
// http://geomalgorithms.com/a03-_inclusion.html/

type Polygon []image.Point

// Polyfence is a geofence defined by a set of points (Polygon).
type Polyfence struct {
	p Polygon
}

// copy returns a copy of the vertex
func (p Polygon) copy() Polygon {
	// create a new vertex with the cap to copy existing vertex into it
	newPolygon := make(Polygon, len(p))
	// use builtin copy
	copy(newPolygon, p)
	// return new vertex
	return newPolygon
}

// NewPolyfence returns a new Polyfence for the given slice of points
func NewPolyfence(p Polygon) *Polyfence {
	// create new polyfence with copy of given vertex
	pf := &Polyfence{
		p: p.copy(),
	}

	// complete the vertex (if necessary) by adding the last point ie given A->B->C->D, add D->A
	if len(pf.p) > 0 && pf.p[0] != pf.p[len(pf.p)-1] {
		pf.p = append(pf.p, pf.p[0])
	}
	return pf
}

// Inside returns whether the given point lies inside the Polyfence
// The used algorithm is Winding Number
func (pf *Polyfence) Inside(pt image.Point) bool {
	// given point lies outside the polygon when winding number equals 0
	if calculateWindingNumber(pf.p, pt) == 0 {
		return false
	}

	// all other values: inside
	return true
}

// perform a winding number calculation on given vertex (polygon) and point
func calculateWindingNumber(p Polygon, pt image.Point) int {
	// quick return for non-poly vertex
	if len(p) < 3 {
		return 0
	}

	// the winding number counter
	wn := 0

	// amount of edges to check
	n := len(p) - 2

	// loop through all edges of the polygon
	for i := 0; i <= n; i++ {
		// start y <= pt.Y
		if p[i].Y <= pt.Y {
			// an upward crossing
			if p[i+1].Y > pt.Y { // if p[i+1].Y is also <= pt there is no crossing
				// P left of  edge
				if isLeft(p[i], p[i+1], pt) > 0 {
					// have  a valid up intersect
					wn++
				}
			}
		} else { // p[i].Y > pt.Y
			// a downward crossing
			if p[i+1].Y <= pt.Y { // if p[i+1].Y is also > pt there is no crossing
				// P right of edge (from the downward edge's POV)
				if isLeft(p[i], p[i+1], pt) < 0 {
					// have  a valid down intersect
					wn--
				}
			}
		}
	}

	// all done
	return wn
}

func isLeft(lineA image.Point, lineB image.Point, pt image.Point) int {
	return (lineB.X-lineA.X)*(pt.Y-lineA.Y) - (pt.X-lineA.X)*(lineB.Y-lineA.Y)
}
