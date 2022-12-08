package main

type point struct {
	x int
	y int
}

type tree struct {
	height                 int
	visN, visE, visW, visS bool // visible from each direction
	visible                bool
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

func (w *world) computeVisibility() {
	for x := 0; x <= w.maxX; x++ {
		w.computeVisibilityFromNorthForCol(x)
		w.computeVisibilityFromSouthForCol(x)
	}
	for y := 0; y <= w.maxY; y++ {
		w.computeVisibilityFromEastForRow(y)
		w.computeVisibilityFromWestForRow(y)
	}
}

func (w *world) computeVisibilityFromWestForRow(y int) {
	// we already know about x == 0
	tallestSoFar := w.forest.getTreeAt(point{0, y}).height
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
		w.forest[p] = t
	}
}
func (w *world) computeVisibilityFromEastForRow(y int) {
	// we already know about x == 0
	tallestSoFar := w.forest.getTreeAt(point{w.maxX, y}).height
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
		w.forest[p] = t
	}
}

func (w *world) computeVisibilityFromNorthForCol(x int) {
	// we already know about y == 0
	tallestSoFar := w.forest.getTreeAt(point{x, 0}).height
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
		w.forest[p] = t
	}
}
func (w *world) computeVisibilityFromSouthForCol(x int) {
	// we already know about y == 0
	tallestSoFar := w.forest.getTreeAt(point{x, w.maxY}).height
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
