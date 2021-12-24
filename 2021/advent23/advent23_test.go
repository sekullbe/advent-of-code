package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestStep(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{B, A}},
			room{0, []int{A, B}},
			room{0, []int{C, C}},
			room{0, []int{D, D}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	assert.Equal(t, 46, c)
}

func TestStep2(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{B, A}},
			room{0, []int{A, B}},
			room{0, []int{D, C}},
			room{0, []int{C, D}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	assert.Equal(t, 4600, c)
}
func TestStep3(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{B, B}},
			room{0, []int{A, A}},
			room{0, []int{C, C}},
			room{0, []int{D, D}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	assert.Equal(t, 114, c)
}

func TestExample(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{B, A}},
			room{0, []int{C, D}},
			room{0, []int{B, C}},
			room{0, []int{D, A}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	for k, s := range stateCache {
		if strings.Contains(k, "-1 -1 -1 -1 -1 -1 -1") {
			log.Printf("%s:%d", k, s)
		}
	}
	assert.Equal(t, 12521, c)
}

// I'm tired and can't be arsed to write parsing code, so if people can do the whole
//puzzle by hand I can enter the inputsn by hand.
func Test_ForReal1(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{D, C}},
			room{0, []int{D, C}},
			room{0, []int{A, B}},
			room{0, []int{A, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	for k, s := range stateCache {
		if strings.Contains(k, "-1 -1 -1 -1 -1 -1 -1") {
			log.Printf("%s:%d", k, s)
		}
	}
	assert.Equal(t, 16489, c)
}

func Test_Example2(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{B, D, D, A}},
			room{0, []int{C, C, B, D}},
			room{0, []int{B, B, A, C}},
			room{0, []int{D, A, C, A}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	for k, s := range stateCache {
		if strings.Contains(k, "-1 -1 -1 -1 -1 -1 -1") {
			log.Printf("%s:%d", k, s)
		}
	}
	assert.Equal(t, 16489, c)
}

func Test_ForReal2(t *testing.T) {
	initialize()
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{D, C}},
			room{0, []int{D, C}},
			room{0, []int{A, B}},
			room{0, []int{A, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	for k, s := range stateCache {
		if strings.Contains(k, "-1 -1 -1 -1 -1 -1 -1") {
			log.Printf("%s:%d", k, s)
		}
	}
	assert.Equal(t, 16489, c)
}
