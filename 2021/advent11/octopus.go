package main

type octopus struct {
	energy    int
	row       int
	col       int
	flashed   bool
	neighbors []*octopus
}

// this is really "flash if energized enough"
// returns number of flashes
func (o *octopus) flash() int {
	flashes := 0
	if o.energy > 9 && o.flashed == false {
		o.flashed = true
		flashes = 1
		for _, neighbor := range o.neighbors {
			neighbor.energize()
			flashes += neighbor.flash()
		}
	}
	return flashes
}

func newOctopus(energy int, row int, col int) octopus {
	return octopus{energy: energy, row: row, col: col, flashed: false}
}

func (o *octopus) reset() {
	if o.flashed {
		o.energy = 0
	}
	o.flashed = false
}

func (o *octopus) energize() {
	o.energy += 1
}

func (me *octopus) attachNeighbors(farm farm) {
	myrow, mycol := me.row, me.col
	for _, o := range farm {
		if o.row == myrow && o.col == mycol {
			continue
		}
		if o.row < myrow-1 || o.row > myrow+1 || o.col < mycol-1 || o.col > mycol+1 {
			continue
		}
		me.neighbors = append(me.neighbors, o)
		// should i do the reverse? no, because we're gonna hook up the neighbor later and it would complicate data structure
	}
}
