package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(50, inputText))
}

type rotation struct {
	dir      rune
	distance int
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")
	rotations := parseRotations(input)
	dial := 50
	zeroes := 0
	for _, r := range rotations {
		dial = newDial(dial, r)
		//log.Println(dial)
		if dial == 0 {
			zeroes++
		}

	}
	return zeroes
}

func parseRotations(input string) []rotation {
	var rotations []rotation
	for _, r := range parsers.SplitByLines(input) {
		var dir rune
		var distance int
		_, err := fmt.Sscanf(r, "%c%d", &dir, &distance)
		if err != nil {
			panic(err)
		}
		rotations = append(rotations, rotation{dir: dir, distance: distance})
	}
	return rotations
}

func run2(start int, input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	rotations := parseRotations(input)
	dial := start
	zeroes := 0
	for _, r := range rotations {
		tick := 1
		if r.dir == 'L' {
			tick = -1
		}
		for i := 0; i < tools.AbsInt(r.distance); i++ {
			dial += tick
			switch dial {
			case 100:
				dial = 0
			case -1:
				dial = 99
			}
			if dial == 0 {
				zeroes++
			}
		}
	}

	return zeroes
}

func newDial(d int, r rotation) int {
	if r.dir == 'R' {
		d += r.distance
	} else {
		d -= r.distance
	}
	// correct turns that ended up like -50 or 378 to -50 or 78 before continuing
	// it would work the same to compare dial%100 to 0 but this is more correct-looking
	d %= 100
	if d < 0 {
		d = 100 + d
	}
	return d
}

func countSpins(d int) int {
	return tools.AbsInt(d / 100)
}
