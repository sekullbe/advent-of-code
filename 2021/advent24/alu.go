package main

import (
	"fmt"
	"log"
	"math"
)

// some code would be simpler if this was a map, but meh
// this is probably going to hurt me if part 2 is "add 27 more registers"
type alu struct {
	w int
	x int
	y int
	z int
}

func newAlu() *alu {
	return &alu{0, 0, 0, 0}
}

func (a *alu) toString() string {
	return fmt.Sprintf("w: %d x: %d, y: %d, z:%d", a.get(W), a.get(X), a.get(Y), a.get(Z))
}

// Figure out the 2nd arg of any instruction- return the value in the register or the constant value
func (a *alu) evaluateValue(s step) int {
	if s.rFrom == NOREG {
		return s.value
	}
	return a.get(s.rFrom)
}

// modifies the inputs slice in place
func (a *alu) inp(s step, inputs *[]int) {
	a.set(s.rTo, (*inputs)[0])
	// got burned here a bit- thought not using pointer would work, but since this just modifies
	// the *slice* structure, which is passed by value, it didn't change the caller's slice
	// if I'd changed the underlying array, *that* would have stuck.
	*inputs = (*inputs)[1:]
}

func (a *alu) add(s step) {
	a.set(s.rTo, a.get(s.rTo)+a.evaluateValue(s))
}

func (a *alu) mul(s step) {
	a.set(s.rTo, a.get(s.rTo)*a.evaluateValue(s))
}

func (a *alu) div(s step) {
	a.set(s.rTo, a.get(s.rTo)/a.evaluateValue(s))
}

func (a *alu) mod(s step) {
	a.set(s.rTo, a.get(s.rTo)%a.evaluateValue(s))
}

func (a *alu) eql(s step) {
	if a.get(s.rTo) == a.evaluateValue(s) {
		a.set(s.rTo, 1)
	} else {
		a.set(s.rTo, 0)
	}
}

func (a *alu) get(rFrom rune) int {
	switch rFrom {
	case W:
		return a.w
	case X:
		return a.x
	case Y:
		return a.y
	case Z:
		return a.z
	}
	log.Panicf("can't read register %c", rFrom)
	return -math.MaxInt
}

func (a *alu) set(reg rune, value int) {
	switch reg {
	case W:
		a.w = value
	case X:
		a.x = value
	case Y:
		a.y = value
	case Z:
		a.z = value
	default:
		log.Panicf("don't know how to write to register %c", reg)

	}
}
