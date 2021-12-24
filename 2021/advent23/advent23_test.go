package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestStep(t *testing.T) {
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

func TestBuriedMover(t *testing.T) {
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{A, B}},
			room{0, []int{B, A}},
			room{0, []int{C, C}},
			room{0, []int{D, D}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	assert.Equal(t, 112, c)
}

func TestScore(t *testing.T) {
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{A, A}},
			room{0, []int{B}},
			room{0, []int{C, C}},
			room{0, []int{D, D}}},
		corridor: corridor{-1, -1, B, -1, -1, -1, -1},
	}
	c := begin(initialState)
	assert.Equal(t, 20, c)
}

func TestStep2(t *testing.T) {
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
	assert.Equal(t, 4646, c)
}
func TestStep3(t *testing.T) {
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

// I'm tired and can't be bothered to write parsing code, so if people can do the whole
// puzzle by hand I can enter the inputs by hand.
func Test_ForReal1(t *testing.T) {
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
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{B, D, D, A}},
			room{0, []int{C, C, B, D}},
			room{0, []int{B, B, A, C}},
			room{0, []int{D, A, C, A}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := beginWithDepth(initialState, 4)
	for k, s := range stateCache {
		if strings.Contains(k, "-1 -1 -1 -1 -1 -1 -1") {
			log.Printf("%s:%d", k, s)
		}
	}
	assert.Equal(t, 44169, c)
}

func Test_ForReal2(t *testing.T) {
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{D, D, D, C}},
			room{0, []int{D, C, B, C}},
			room{0, []int{A, B, A, B}},
			room{0, []int{A, A, C, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := beginWithDepth(initialState, 4)
	for k, c := range stateCache {
		if strings.Contains(k, "-1 -1 -1 -1 -1 -1 -1") {
			log.Printf("%s:%d", k, c)
		}
	}
	assert.Equal(t, 43413, c)
}

func Test_ExampleFromReddit2(t *testing.T) {
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{A, C}},
			room{0, []int{D, C}},
			room{0, []int{A, D}},
			room{0, []int{B, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := begin(initialState)
	assert.Equal(t, 13495, c)
}
func Test_ExampleFromReddit4(t *testing.T) {
	initialState := &state{
		cost: 0,
		rooms: []room{
			room{0, []int{A, D, D, C}},
			room{0, []int{D, C, B, C}},
			room{0, []int{A, B, A, D}},
			room{0, []int{B, A, C, B}}},
		corridor: corridor{-1, -1, -1, -1, -1, -1, -1},
	}
	c := beginWithDepth(initialState, 4)
	assert.Equal(t, 53767, c)
}
