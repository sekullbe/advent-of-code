package main

type die struct {
	lastRoll int
	sides    int
}

func newDie(s int) die {
	var d die
	d.sides = s
	d.lastRoll = 0
	return d
}

func (d *die) roll() int {
	d.lastRoll += 1
	if d.lastRoll > d.sides {
		d.lastRoll = 1
	}
	i := d.lastRoll
	return i
}

func (d die) read() int {
	return d.lastRoll
}

func (d *die) rolls(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += d.roll()
	}
	return total
}

/*
func newDie() die {
	var d die
	d = 0
	return d
}

func (d *die) roll() int {
	*d += 1
	if *d > 100 {
		*d = 1
	}
	i := int(*d)
	return i
}

func (d die) read() int {
	return int(d)
}

func (d *die) universes(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += d.roll()
	}
	return total
}
*/
