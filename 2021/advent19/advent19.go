package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
	"regexp"
	"strconv"
)

//go:embed input.txt
var inputText string

//go:embed test1.txt
var testText string

func main() {
	beaconCount, maxScannerDist := run1(inputText)
	fmt.Printf("Magic number: %d\n", beaconCount)
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", maxScannerDist)
}

func run1(inputText string) (int, int) {
	scanners := parseInput(inputText)
	s0 := scanners[0]
	s0.calculateDistanceFingerprint(false)
	scannersToMatch := scanners[1:]
	var locatedScanners []beacon // not really a beacon, but beacon == point
	locatedScanners = append(locatedScanners, s0.location)
	for len(scannersToMatch) > 0 {
		// get a scanner with 12+ matching beacons, and a list of the matching pairs
		matchingScanner, matches := findBestScannerMatchingFingerprint(s0.fingerprint, scannersToMatch)
		if matchingScanner == nil {
			log.Panicf("no matching beacon sets, with %d scanners left", len(scannersToMatch))
		}
		log.Printf("Scanner %d matched- lining up beacons", matchingScanner.id)
		_, x, y, z, rotatedBeacons := matchScanner(s0, matchingScanner, matches)
		// using these, we can add all of the beacons on matchingscanner to s0
		for _, rb := range rotatedBeacons {
			s0.addBeacon(translateBeacon(rb, x, y, z))
		}
		s0.calculateDistanceFingerprint(true)
		log.Printf("s0 now has %d beacons", len(s0.beacons))
		matchingScanner.location.translate(x, y, z)
		locatedScanners = append(locatedScanners, matchingScanner.location)
		scannersToMatch = removeMatchedScanner(matchingScanner, scannersToMatch)
	}
	log.Printf("done!")

	return len(s0.beacons), calcMaxManhattanDistance(locatedScanners)
}

func removeMatchedScanner(s *scanner, scanners []*scanner) (filteredScanners []*scanner) {
	for _, s2 := range scanners {
		if s2.id != s.id {
			filteredScanners = append(filteredScanners, s2)
		}
	}
	return filteredScanners
}

func matchScanner(s0 *scanner, matchingScanner *scanner, matches []twoBeaconPairs) (int, int, int, int, []beacon) {

	// key is rotation type, value is []beacon
	rotations := matchingScanner.getAllPossibleRotatedBeacons()
	// try the translations from the first match- maybe we don't have to loop them all
	// bp1.p1 = bp2.p1 and bp1.p2 = bp2.p2, or bp1.p2 = bp2.p1 and bp1.p1 = bp2.p2
	// four transforms to try ( * 24 rotations...)
	for matchToTry, _ := range matches {
		bpToTry := []beaconPair{
			{matches[matchToTry].bp1.p1, matches[matchToTry].bp2.p1},
			{matches[matchToTry].bp1.p1, matches[matchToTry].bp2.p2},
			{matches[matchToTry].bp1.p2, matches[matchToTry].bp2.p1},
			{matches[matchToTry].bp1.p2, matches[matchToTry].bp2.p2},
		}
		for rotType, rotatedBeacons := range rotations {
			for _, b := range bpToTry {
				newBp := beaconPair{b.p1, rotateByType(b.p2, rotType)}
				translatedBeacons := translateBeacons(rotatedBeacons, newBp)
				// if this is true, we know how to align s0 and matchingscanner
				match := checkBeaconMatch(s0.beacons, translatedBeacons, 12)
				if match {
					//log.Printf("rotation type %d", rotType)
					x, y, z := calculateBeaconDiff(newBp.p1, newBp.p2)
					log.Printf("Beacons aligned: translation %d,%d,%d", x, y, z)
					// So now we know the rotType and x,y,z translation to line up matchingScanner.id with s0
					return rotType, x, y, z, rotatedBeacons
				}
			}
		}
	}
	log.Panicf("matchScanner should have found something")
	return 0, 0, 0, 0, nil
}

func parseInput(inputText string) (scanners []*scanner) {
	scannerre := regexp.MustCompile(`--- scanner (\d+) ---`)
	beaconre := regexp.MustCompile(`(-?\d+),(-?\d+),(-?\d+)`)
	var s *scanner
	for _, line := range parsers.SplitByLines(inputText) {
		if scannerre.MatchString(line) {
			sid, _ := strconv.Atoi(scannerre.FindStringSubmatch(line)[1])
			s = newScanner(sid)
			scanners = append(scanners, s)
		} else {
			// parse beacons until empty line
			bcoords := beaconre.FindStringSubmatch(line)
			if bcoords == nil {
				continue // probably the empty line
			}
			bx, _ := strconv.Atoi(bcoords[1])
			by, _ := strconv.Atoi(bcoords[2])
			bz, _ := strconv.Atoi(bcoords[3])
			s.addBeaconCoords(bx, by, bz)
		}
	}
	/*
		// For testability, don't do this here, so parseInput test isn't a mess
		for _, s := range scanners {
			s.calculateDistanceFingerprint()
		}
	*/

	return
}

// returns the first scanner in the list with matching fingerprints, and the set of matching beacon pairs
func findBestScannerMatchingFingerprint(f0 fingerprint, scanners []*scanner) (matchingScanner *scanner, matches []twoBeaconPairs) {
	var match bool
	type mcount struct {
		matches int
		s       *scanner
		m       []twoBeaconPairs
	}
	matchers := make(map[int]mcount)
	for _, s := range scanners {
		setFingerprints(scanners)
		match, matches = matchFingerprints(f0, s.fingerprint, 12)
		if match {
			matchers[s.id] = mcount{len(matches), s, matches}
			matchingScanner = s
		}
	}
	var bestms *scanner
	var bestmatches []twoBeaconPairs
	max := 0
	for _, mc := range matchers {
		if mc.matches > max {
			max = mc.matches
			bestms = mc.s
			bestmatches = mc.m
		}
	}
	return bestms, bestmatches
	//	return matchingScanner, matches
}

func matchFingerprints(f0, f1 fingerprint, overlap int) (bool, []twoBeaconPairs) {
	matchCount := 0
	var matches []twoBeaconPairs
	// a fingerprint is a map of about b^2 pairs
	// it's a match if there are 12 distances in f0 that are also in f1
	for f0p, f0f := range f0 {
		for f1p, f1f := range f1 {
			if f0f == f1f {
				matchCount++
				// a,b in f0p  and c,d in f1p - could be a=c and b=d, or a=c and b=d
				matches = append(matches, twoBeaconPairs{f0p, f1p})
			}
		}
	}
	//log.Printf("Fingerprint matches: %d", matchCount)
	return matchCount >= overlap, matches
}

/*
start at s0
look at each scanner s1...
check for overlaps
do this by matching fingerprints between s0 and sn
if you have a match, you know s0 and sn overlap, and you have a pair of matching beacons

rotate and translate sn until yu get the matches, then put all of the beacons into s0
lather rinse repeat


*/

/*
so... how to do this?
start with scanner 0- it is at the origin and defines the coordinate system
get its lists of rotated points?

take a new scanner
take its lists of rotated points
There are 30 scanners which means the absolutely max coordinate is +/- 60000
 -- if the beacons and scanners were in a perfect line each scanner would be 2000 away from the next; add 1 for paranoia
fuck, i don't want to do a loop over 2E14 possibilities, how to constrain that?
obviously we'll cut short once 12 beacons are found

use the distance between beacons in a set as a fingerprint

calculate distances between each pair of beacons in the scanner - do you need to store xyz or can you store the float?
now compare each pair in s0 with s1
if you find 11 overlaps,  you know that those two beacons can merge the scanners

my fingerprints are bad
for every beacon in S0 you need to store the distances to every other beacon in S0, not the sum to ALL in S0
so a fingerprint is a map from beaconpair to float



-----
Given a group of beacons belonging to a scanner, calculate the set of absolute distances to all other beacons in the group
(throwing away position and sign information). So for the first two beacons of the first group, the set of distances would
be 81, 1, 163, though the order does no longer matter. I stored these in a frozenset in Python. Now, for each beacon, you
have a collection of distances to other beacons in the same group.

Now compare two scanner groups, and for each pair of beacons, check the number of distance-sets that overlap. If that is
at least 11, you know at least 12 beacons overlap, and you can use those two beacons as reference points to merge the
scanner groups. Theoretically you could have some false positives, as you throw away sign information, but in practice
this won't lead to problems.

Now given the reference points in the previous step, find another beacon that is present in both scanner groups. It
should not have the same x, y or z as the reference beacon (distance-set should not contain 0) and should be unique
within the set. You can use this to determine the orientation of the second scanner group relative to the first.

To merge both sets, for every beacon in the second group, rotate it using the previously determined orientation, then
translate it by the difference in scanner positions. If the new beacon position is not present in the first group, add
it to the first group. Finally, remove the second group and recalculate the distance-sets.


*/
