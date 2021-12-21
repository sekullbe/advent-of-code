package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDie(t *testing.T) {
	d100 := newDie(100)
	r := d100.roll()
	assert.Equal(t, 1, r)
	assert.Equal(t, 1, d100.read())
	assert.Equal(t, 1, d100.read())

	r = d100.roll()
	assert.Equal(t, 2, r)
	assert.Equal(t, 2, d100.read())
}

func TestRolls(t *testing.T) {
	d100 := newDie(100)
	assert.Equal(t, 1+2+3+4+5+6, d100.rolls(6))
	assert.Equal(t, 6, d100.read())
}

func TestDieRollover(t *testing.T) {
	d100 := newDie(100)
	d100.lastRoll = 99

	r := d100.roll()
	assert.Equal(t, 100, r)
	assert.Equal(t, 100, d100.read())

	r = d100.roll()
	assert.Equal(t, 1, r)
	assert.Equal(t, 1, d100.read())
}
