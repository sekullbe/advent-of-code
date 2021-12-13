package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_grid_fold(t *testing.T) {
	g := grid{point{2, 0}: true}
	assert.False(t, g.getPoint(point{0, 0}))
	assert.True(t, g.getPoint(point{2, 0}))

	g.fold("x", 1)

	assert.True(t, g.getPoint(point{0, 0}))
	assert.False(t, g.getPoint(point{2, 0}))

	g.addPoint(point{10, 5})
	fmt.Println(g.display())
}

func Test_grid_foldover(t *testing.T) {
	g := grid{point{2, 0}: true, point{3, 0}: true, point{6, 1}: true}
	assert.False(t, g.getPoint(point{0, 0}))
	assert.True(t, g.getPoint(point{2, 0}))

	g.fold("x", 2)

	assert.True(t, g.getPoint(point{1, 0}))
	assert.True(t, g.getPoint(point{2, 0}))
	assert.True(t, g.getPoint(point{-2, 1}))

}
