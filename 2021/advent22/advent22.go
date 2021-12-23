package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"regexp"
)

//go:embed input.txt
var inputText string

type cube struct {
	x, y, z int
}
type reactor map[cube]bool

// originally called 'steps' which made more sense for part 1 than part 2,
// which is why 'executeStep' is called that
type block struct {
	on     bool
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

const minCoord int = -50
const maxCoord int = 50

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	steps := parseInput(inputText)
	r := make(reactor)
	for _, s := range steps {
		executeStep(r, s)
	}
	return countActiveCubes(r)
}

func run2(inputText string) int {
	blocks := parseInput(inputText)
	adjustedBlocks := []block{} // original blocks plus any adjustment blocks from overlaps
	for _, b := range blocks {
		newAdjustmentBlocks := []block{} // new in this pass
		// Look back at every existing block including previously-added adjustment block and check for overlaps
		for _, ab := range adjustedBlocks {
			ob, isOverlap := findOverlap(b, ab)
			if !isOverlap {
				continue
			}

			// the overlap block will become an adjustment block to add or subtract any extras that the overlap generated
			if b.on && ob.on {
				// eg 3-cube ON block + 3-cube ON block with 1 overlap; you want 3+3-1 = 5
				// so in addition to the pair of 3-cube ONs, add a 1-cube OFF adjustment block
				ob.on = false
			} else if !b.on && !ob.on {
				// since ob is off, it must be an adjustment block already created by overlapping an OFF block on an ON block,
				// so we are already subtracting the overlapping OFF block. Need to add an ON block to balance the
				// second OFF overlap block we are adding
				// eg assume a large ON block (20 cubes), with a 3-cube OFF block completely overlapping it: 17 cubes
				// The new block is a 3-cube OFF block that is also completely in the ON block, and overlaps the OFF block by one
				// you want to subtract only 2 more, not 3, so add a 1-cube ON adjustment block to compensate
				ob.on = true
			} else {
				// Match the pre-existing block that is being overlapped
				// eg: have a 3-cube ON block and are adding a 3-cube OFF block, with 1 overlap
				//   that turns off 1 of the 3 ON; the result you want is 2, so you need to add a 1-cube OFF adjustment
				// and the trickiest case:
				// eg: have a 3-cube OFF block already overlapping a 8-cube ON block (it must already be an overlap of an ON block or it wouldn't be here)
				//    Now you are adding a 3-cube ON block that overlaps one cube from the OFF block, turning it back on.
				//      The result you want is 8:  8 turned on, -3 turned off ,+1 turned back on, +2 newly on
				//    This new block will overlap by 1 with BOTH the original 8-cube ON block and the 3-cube OFF block, and create two adjustment blocks:
				//      first from overlapping the 8-ON, a 1-OFF based on the first rule above (ON+ON means add an OFF to compensate) ; -1 OFF, total is 7
				//      then from overlapping the 3-OFF, you're adding a 1-ON by turning one of those back on- +1 ON, total is 8
				// Since this is layering adjustment blocks, never deleting them, you have to compensate instead of going back
				// to "edit" what was already added
				ob.on = b.on
			}
			newAdjustmentBlocks = append(newAdjustmentBlocks, ob)
		}
		// we've added adjustment blocks for any overlaps, so if this block is OFF ignore it; just add ON blocks
		// eg 3-cube ON block, adding 3-cube OFF overlapping by one
		// previously we added a 1-cube OFF adjustment block to newAdjustmentBlocks
		// so here we add the full original 3-cube block
		if b.on {
			adjustedBlocks = append(adjustedBlocks, b)
		}
		// add any adjustment blocks from this cycle
		adjustedBlocks = append(adjustedBlocks, newAdjustmentBlocks...)
	}
	onCubes := 0
	for _, b := range adjustedBlocks {
		if b.on {
			onCubes += b.size()
		} else {
			onCubes -= b.size()
		}
	}
	return onCubes
}

func parseInput(input string) (steps []block) {
	lines := parsers.SplitByLines(input)
	re := regexp.MustCompile(`(on|off) x=([-0-9]+)\.\.([-0-9]+),y=([-0-9]+)\.\.([-0-9]+),z=([-0-9]+)\.\.([-0-9]+)`)
	for _, line := range lines {
		s := block{}
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			log.Panicf("stop! panic time! can't parse this: %s", line)
		}
		s.on = matches[1] == "on"
		s.x1 = tools.Atoi(matches[2])
		s.x2 = tools.Atoi(matches[3])
		s.y1 = tools.Atoi(matches[4])
		s.y2 = tools.Atoi(matches[5])
		s.z1 = tools.Atoi(matches[6])
		s.z2 = tools.Atoi(matches[7])
		steps = append(steps, s)
	}
	return steps
}

func executeStep(r reactor, s block) {
	toggleCube(r, s.on, s.x1, s.x2, s.y1, s.y2, s.z1, s.z2)
}

func toggleCube(r reactor, on bool, x1, x2, y1, y2, z1, z2 int) {
	for x := restrictCoordinate(x1, false); x <= restrictCoordinate(x2, true); x++ {
		for y := restrictCoordinate(y1, false); y <= restrictCoordinate(y2, true); y++ {
			for z := restrictCoordinate(z1, false); z <= restrictCoordinate(z2, true); z++ {
				r[cube{x, y, z}] = on
			}
		}
	}
}

func restrictCoordinate(c int, max bool) int {
	if max && c > maxCoord {
		return maxCoord
	} else if !max && c < minCoord {
		return minCoord
	}
	return c
}

func countActiveCubes(r reactor) int {
	count := 0
	for cube, on := range r {
		if on && cube.x >= minCoord && cube.x <= maxCoord && cube.y >= minCoord && cube.y <= maxCoord && cube.z >= minCoord && cube.z <= maxCoord {
			count++
		}
	}
	return count
}

func (b block) size() int {
	return (b.x2 - b.x1 + 1) * (b.y2 - b.y1 + 1) * (b.z2 - b.z1 + 1)
}

// b1 is the first block, b2 is the one being added
// returns the overlapping block with b2's 'on' status
// use "comma,ok" instead of nil-checking 'overlap'
func findOverlap(b1, b2 block) (overlap block, isOverlap bool) {
	// overlap is the max starting coordinate and min ending coordinate
	// i.e. 3-7 and 6-10: overlap is 6-7
	// and if there is no overlap in ANY axis, there is no overlap at all
	// that looks like an inverted overlap region when calculated like this (i.e. x2 > x1)
	newx1 := tools.MaxInt(b1.x1, b2.x1)
	newx2 := tools.MinInt(b1.x2, b2.x2)
	if newx1 > newx2 {
		return
	}

	newy1 := tools.MaxInt(b1.y1, b2.y1)
	newy2 := tools.MinInt(b1.y2, b2.y2)
	if newy1 > newy2 {
		return
	}

	newz1 := tools.MaxInt(b1.z1, b2.z1)
	newz2 := tools.MinInt(b1.z2, b2.z2)
	if newz1 > newz2 {
		return
	}

	// overlap region has the same on/off status as the block it's carved from
	overlap = block{on: b2.on, x1: newx1, x2: newx2, y1: newy1, y2: newy2, z1: newz1, z2: newz2}
	isOverlap = true
	return
}
