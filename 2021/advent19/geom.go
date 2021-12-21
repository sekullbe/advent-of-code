package main

import (
	"github.com/sekullbe/advent/tools"
	"math"
)

type beacon struct {
	x, y, z int
}

type beaconPair struct {
	p1 beacon
	p2 beacon
}

type twoBeaconPairs struct {
	bp1 beaconPair
	bp2 beaconPair
}

type fingerprint map[beaconPair]float64

type scanner struct {
	id             int
	location       beacon   // the scanner's notional position; it's a point not a 'beacon' but it made code more clear because every other point is a beacon
	beacons        []beacon // beacons relative to the scanner's location and orientation as currently stored
	rotatedBeacons [][]beacon
	fingerprint    fingerprint
}

func translateBeacon(b beacon, x, y, z int) beacon {
	return beacon{b.x + x, b.y + y, b.z + z}
}

func (p *beacon) translate(x, y, z int) {
	p.x += x
	p.y += y
	p.z += z
}

// commutativity doesn't matter because I just need to do one at a time,
// x,y,z is pitch, roll, yaw
func (p *beacon) rotate(pitch, roll, yaw int) {
	*p = rotate(*p, pitch, roll, yaw)
}

func rotate(p0 beacon, pitch, roll, yaw int) (p beacon) {
	p = beacon{p0.x, p0.y, p0.z}
	for i := 0; i < normalizeAngle(pitch); i++ {
		p.y, p.z = -p.z, p.y
	}
	for i := 0; i < normalizeAngle(roll); i++ {
		p.x, p.z = p.z, -p.x
	}
	for i := 0; i < normalizeAngle(yaw); i++ {
		p.x, p.y = -p.y, p.x
	}
	return p
}

func normalizeAngle(a int) int {
	a = a % 4
	if a < 0 {
		a = 4 + a
	}
	return a
}

func newScanner(id int) *scanner {
	return &scanner{id: id}
}

func (s *scanner) addBeaconCoords(x, y, z int) {
	s.addBeacon(beacon{x, y, z})
}

func (s *scanner) addBeacon(b beacon) {
	for _, b2 := range s.beacons {
		if b == b2 {
			return
		}
	}
	s.beacons = append(s.beacons, b)
}

func (s *scanner) translate(x, y, z int) {
	s.beacons = s.getTranslatedBeacons(x, y, z)
}

func translateBeacons(beacons []beacon, bp beaconPair) []beacon {
	ourBeacon := bp.p1
	theirBeacon := bp.p2
	x, y, z := calculateBeaconDiff(ourBeacon, theirBeacon)
	//log.Printf("translation: %d,%d,%d", x, y, z)

	translatedBeacons := []beacon{}
	for _, b := range beacons {
		translatedBeacons = append(translatedBeacons, translateBeacon(b, x, y, z))
	}
	return translatedBeacons
}

func (s scanner) translateAllBeacons(bp beaconPair) []beacon {
	ourBeacon := bp.p1
	theirBeacon := bp.p2
	x, y, z := calculateBeaconDiff(ourBeacon, theirBeacon)

	return s.getTranslatedBeacons(x, y, z)
}

func calculateBeaconDiff(b1, b2 beacon) (x, y, z int) {
	x = b1.x - b2.x
	y = b1.y - b2.y
	z = b1.z - b2.z
	return
}

// Separated these out to allow 'preview' of translated beacons
func (s scanner) getTranslatedBeacons(x, y, z int) []beacon {
	var newBeacons []beacon
	copy(s.beacons, newBeacons)
	for _, beacon := range newBeacons {
		beacon.translate(x, y, z)
	}
	return newBeacons
}

// true if the beacon sets have at least 'overlap' points in common
func checkBeaconMatch(b0, b1 []beacon, overlap int) bool {
	matchCount := 0
	for _, b0b := range b0 {
		for _, b1b := range b1 {
			//if b0b == b1b {
			if b0b.x == b1b.x && b0b.y == b1b.y && b0b.z == b1b.z {
				matchCount++
			}
		}
	}
	return matchCount >= overlap
}

// get all the possible rotated beacons for a scanner
// a scanner has 24 possible orientations - 6 directions (+/- x,y,z) and 4 rotations (flat, left, right, inverted)
// so for each beacon relative to the scanner, there are 24
// returns 24 slices of rotated points - thanks to u/PityUpvote for python original
func (s *scanner) getAllPossibleRotatedBeacons() (beacons [][]beacon) {
	if len(s.rotatedBeacons) > 0 {
		return s.rotatedBeacons
	}
	for i := 0; i < 24; i++ {
		beacons = append(beacons, []beacon{})
	}
	for _, b := range s.beacons {
		// positive x
		for i := 0; i < 24; i++ {
			beacons[i] = append(beacons[i], rotateByType(b, i))
		}
	}
	s.rotatedBeacons = beacons
	return
}

func rotateByType(b beacon, rotType int) beacon {
	switch rotType {
	case 0:
		return beacon{+b.x, +b.y, +b.z}
	case 1:
		return beacon{+b.x, -b.z, +b.y}
	case 2:
		return beacon{+b.x, -b.y, -b.z}
	case 3:
		return beacon{+b.x, +b.z, -b.y}
	// negative x
	case 4:
		return beacon{-b.x, -b.y, +b.z}
	case 5:
		return beacon{-b.x, +b.z, +b.y}
	case 6:
		return beacon{-b.x, +b.y, -b.z}
	case 7:
		return beacon{-b.x, -b.z, -b.y}
	// positive y
	case 8:
		return beacon{+b.y, +b.z, +b.x}
	case 9:
		return beacon{+b.y, -b.x, +b.z}
	case 10:
		return beacon{+b.y, -b.z, -b.x}
	case 11:
		return beacon{+b.y, +b.x, -b.z}
	// negative y
	case 12:
		return beacon{-b.y, -b.z, +b.x}
	case 13:
		return beacon{-b.y, +b.x, +b.z}
	case 14:
		return beacon{-b.y, +b.z, -b.x}
	case 15:
		return beacon{-b.y, -b.x, -b.z}
	// positive z
	case 16:
		return beacon{+b.z, +b.x, +b.y}
	case 17:
		return beacon{+b.z, -b.y, +b.x}
	case 18:
		return beacon{+b.z, -b.x, -b.y}
	case 19:
		return beacon{+b.z, +b.y, -b.x}
	// negative z
	case 20:
		return beacon{-b.z, -b.x, +b.y}
	case 21:
		return beacon{-b.z, +b.y, +b.x}
	case 22:
		return beacon{-b.z, +b.x, -b.y}
	case 23:
		return beacon{-b.z, -b.y, -b.x}
	}
	return b
}

func dist(p1 beacon, p2 beacon) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2))
}

func setFingerprints(scanners []*scanner) {
	for _, s := range scanners {
		s.calculateDistanceFingerprint(false)
	}
}

// a fingerprint is a map beacon -> float
// for each beacon we store the sum of its distances to all other beacons in s
// so if we look at some other scanner and see the same numbers, we know there is overlap
func (s *scanner) calculateDistanceFingerprint(recalculate bool) fingerprint {
	if len(s.fingerprint) > 0 && !recalculate {
		return s.fingerprint
	}
	fingerprint := make(fingerprint)
	for _, b1 := range s.beacons {
		for _, b2 := range s.beacons {
			if b1 == b2 {
				continue
			}
			// Remove duplicates
			_, fpb1b2 := fingerprint[beaconPair{b1, b2}]
			_, fpb2b1 := fingerprint[beaconPair{b2, b1}]
			if fpb1b2 || fpb2b1 {
				continue
			}
			d := dist(b1, b2)
			fingerprint[beaconPair{b1, b2}] = d
		}
	}
	s.fingerprint = fingerprint
	return fingerprint
}

func calcMaxManhattanDistance(points []beacon) int {
	max := 0
	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}
			d := manhattanDistance(p1, p2)
			if d > max {
				max = d
			}
		}
	}

	return max
}

func manhattanDistance(p1, p2 beacon) int {
	return tools.AbsInt(p1.x-p2.x) + tools.AbsInt(p1.y-p2.y) + tools.AbsInt(p1.z-p2.z)
}
