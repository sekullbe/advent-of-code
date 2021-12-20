package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed test.txt
var testText string

// make sure this does what I think it does
func TestBuildByShifting(t *testing.T) {
	n := 0

	n <<= 1
	n |= 1
	n <<= 1
	n |= 1
	n <<= 1
	n |= 1
	assert.Equal(t, 7, n)
	n <<= 1
	n |= 0
	assert.Equal(t, 14, n)

	n = 0
	n <<= 1
	n <<= 1
	n <<= 1
	n <<= 1
	n |= 1
	n <<= 1
	n <<= 1
	n <<= 1
	n <<= 1
	n |= 1
	n <<= 1
	assert.Equal(t, 34, n)
}

func TestEnhancePixel(t *testing.T) {
	im, a := parseInput(testText)
	lit := enhancePixel(im, pixel{2, 2}, a, false)
	assert.True(t, lit)

	lit = enhancePixel(im, pixel{0, 0}, a, false)
	assert.False(t, lit)
	lit = enhancePixel(im, pixel{1, 0}, a, false)
	assert.True(t, lit)
}

func TestRun1(t *testing.T) {
	assert.Equal(t, 35, run1(testText))
}

func TestRun2(t *testing.T) {
	assert.Equal(t, 3351, run2(testText))
}
