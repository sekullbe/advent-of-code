package main

import "log"

type point struct {
	x int
	y int
}

type tree struct {
	height                 int
	visN, visE, visW, visS bool // visible from each direction
	visible                bool
	sW, sE, sN, sS         int //scenic scores for each direction
	scenicScore            int
}

type forest map[point]tree

type world struct {
	forest     forest
	maxX, maxY int
}

func newWorld(maxX, maxY int) world {
	w := world{maxX: maxX, maxY: maxY, forest: newforest()}
	return w
}

func newforest() forest {
	return make(forest)
}

func (w *world) addTree(x, y, h int) point {
	f := w.forest
	p := point{x: x, y: y}
	t := tree{height: h}
	if x == 0 {
		t.visible = true
		t.visW = true
	}
	if y == 0 {
		t.visible = true
		t.visN = true
	}
	if x == w.maxX {
		t.visible = true
		t.visE = true
	}
	if y == w.maxY {
		t.visible = true
		t.visS = true
	}
	f[p] = t
	return p
}

func (f forest) getTreeAt(p point) tree {
	return f[p]
}

// this was the most annoying part of the whole thing, to convert the map to a forest without writing over the edge and panicing
func (w *world) display() string {
	f := w.forest
	var out string
	for y := 0; y <= w.maxY; y++ {
		for x := 0; x <= w.maxX; x++ {
			//out += strconv.Itoa(f.getTreeAt(point{x, y}).height)
			if f.getTreeAt(point{x, y}).visible {
				out += "*"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

func (w *world) computeVisibility() int {
	for x := 0; x <= w.maxX; x++ {
		w.computeVisibilityFromNorthForCol(x)
		w.computeVisibilityFromSouthForCol(x)
	}
	for y := 0; y <= w.maxY; y++ {
		w.computeVisibilityFromEastForRow(y)
		w.computeVisibilityFromWestForRow(y)
	}
	maxScenicScore := 0
	for p, t := range w.forest {
		scenicScore := t.sE * t.sN * t.sW * t.sS
		if scenicScore > maxScenicScore {
			maxScenicScore = scenicScore
			log.Printf("Tree at %d,%d has score E%d*W%d*N%d*S%d= %d\n", p.x, p.y, t.sE, t.sW, t.sN, t.sS, scenicScore)
		}
	}
	return maxScenicScore
}

// keep a map of the last coordinate a tree of a given height was seen
// then for this tree, it is height H; look in the map for any height H->9
// subtract our X from that X, that's the distance

type lastSeenTreeOfEachHeight map[int]int

// Finds the last seen stored tree of supplied height or taller
func (lst lastSeenTreeOfEachHeight) findNearestTallTreeFromNorthOrWest(height int) int {
	lastSeenTallTreeLoc := 0
	for ht := height; ht <= 9; ht++ {
		l := lst[ht]
		if l >= lastSeenTallTreeLoc {
			lastSeenTallTreeLoc = l
		}
	}
	return lastSeenTallTreeLoc
}

func (lst lastSeenTreeOfEachHeight) findNearestTallTreeFromSouthOrEast(height int, max int) int {
	lastSeenTallTreeLoc := max
	for ht := height; ht <= 9; ht++ {
		l := lst[ht]
		if l == 0 {
			l = max
		}
		if l <= lastSeenTallTreeLoc {
			lastSeenTallTreeLoc = l
		}
	}
	return lastSeenTallTreeLoc
}

func (w *world) computeVisibilityFromWestForRow(y int) {
	// we already know about x == 0
	tallestSoFar := w.forest.getTreeAt(point{0, y}).height
	lst := make(lastSeenTreeOfEachHeight)
	for x := 1; x < w.maxX; x++ {
		p := point{x, y}
		t := w.forest.getTreeAt(p)
		if t.height > tallestSoFar {
			t.visW = true
			t.visible = true
			tallestSoFar = t.height
		} else {
			t.visW = false
		}
		t.sW = x - lst.findNearestTallTreeFromNorthOrWest(t.height)
		lst[t.height] = x
		w.forest[p] = t
	}
}

func (w *world) computeVisibilityFromEastForRow(y int) {
	// we already know about x == 0
	tallestSoFar := w.forest.getTreeAt(point{w.maxX, y}).height
	lst := make(lastSeenTreeOfEachHeight)
	for x := w.maxX - 1; x >= 1; x-- {
		p := point{x, y}
		t := w.forest.getTreeAt(p)
		if t.height > tallestSoFar {
			t.visE = true
			t.visible = true
			tallestSoFar = t.height
		} else {
			t.visE = false
		}
		ntt := lst.findNearestTallTreeFromSouthOrEast(t.height, w.maxX)
		if ntt == 0 {
			ntt = w.maxX
		}
		t.sE = ntt - x
		lst[t.height] = x
		w.forest[p] = t
	}
}

func (w *world) computeVisibilityFromNorthForCol(x int) {
	// we already know about y == 0
	tallestSoFar := w.forest.getTreeAt(point{x, 0}).height
	lst := make(lastSeenTreeOfEachHeight)
	for y := 1; y < w.maxY; y++ {
		p := point{x, y}
		t := w.forest.getTreeAt(p)
		if t.height > tallestSoFar {
			t.visN = true
			t.visible = true
			tallestSoFar = t.height
		} else {
			t.visN = false
		}
		t.sN = y - lst.findNearestTallTreeFromNorthOrWest(t.height)
		lst[t.height] = y
		w.forest[p] = t
	}
}
func (w *world) computeVisibilityFromSouthForCol(x int) {
	// we already know about y == 0
	tallestSoFar := w.forest.getTreeAt(point{x, w.maxY}).height
	lst := make(lastSeenTreeOfEachHeight)
	for y := w.maxY - 1; y >= 1; y-- {
		p := point{x, y}
		t := w.forest.getTreeAt(p)
		if t.height > tallestSoFar {
			t.visS = true
			t.visible = true
			tallestSoFar = t.height
		} else {
			t.visS = false
		}
		ntt := lst.findNearestTallTreeFromSouthOrEast(t.height, w.maxY)
		if ntt == 0 {
			ntt = w.maxY
		}
		t.sS = ntt - y
		lst[t.height] = y
		w.forest[p] = t
	}
}

func (f forest) countVisibleTrees() int {
	count := 0
	for _, t := range f {
		if t.visible {
			count++
		}
	}
	return count
}
