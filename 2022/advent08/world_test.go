package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_world_computeVisibilityFromWestForRow(t *testing.T) {

	// 6x1 world  123352
	w := newWorld(5, 0)
	w.addTree(0, 0, 1)
	w.addTree(1, 0, 2)
	w.addTree(2, 0, 3)
	w.addTree(3, 0, 3)
	w.addTree(4, 0, 5)
	w.addTree(5, 0, 2)

	w.computeVisibilityFromWestForRow(0)
	assert.True(t, w.forest.getTreeAt(point{0, 0}).visW)
	assert.True(t, w.forest.getTreeAt(point{1, 0}).visW)
	assert.True(t, w.forest.getTreeAt(point{2, 0}).visW)
	assert.False(t, w.forest.getTreeAt(point{3, 0}).visW)
	assert.True(t, w.forest.getTreeAt(point{4, 0}).visW)
	assert.False(t, w.forest.getTreeAt(point{5, 0}).visW)
}

func Test_world_computeVisibilityFromEastForRow(t *testing.T) {

	// 6x1 world  163442
	w := newWorld(5, 0)
	w.addTree(0, 0, 1)
	w.addTree(1, 0, 6)
	w.addTree(2, 0, 3)
	w.addTree(3, 0, 4)
	w.addTree(4, 0, 4)
	w.addTree(5, 0, 2)

	w.computeVisibilityFromEastForRow(0)
	assert.False(t, w.forest.getTreeAt(point{0, 0}).visE)
	assert.True(t, w.forest.getTreeAt(point{1, 0}).visE)
	assert.False(t, w.forest.getTreeAt(point{2, 0}).visE)
	assert.False(t, w.forest.getTreeAt(point{3, 0}).visE)
	assert.True(t, w.forest.getTreeAt(point{4, 0}).visE)
	assert.True(t, w.forest.getTreeAt(point{5, 0}).visE)
}

func Test_world_computeVisibilityFromNorthForCol(t *testing.T) {

	// 1x6 world  163742
	w := newWorld(0, 5)
	w.addTree(0, 0, 1)
	w.addTree(0, 1, 6)
	w.addTree(0, 2, 3)
	w.addTree(0, 3, 7)
	w.addTree(0, 4, 4)
	w.addTree(0, 5, 2)

	w.computeVisibilityFromNorthForCol(0)
	assert.True(t, w.forest.getTreeAt(point{0, 0}).visN)
	assert.True(t, w.forest.getTreeAt(point{0, 1}).visN)
	assert.False(t, w.forest.getTreeAt(point{0, 2}).visN)
	assert.True(t, w.forest.getTreeAt(point{0, 3}).visN)
	assert.False(t, w.forest.getTreeAt(point{0, 4}).visN)
	assert.False(t, w.forest.getTreeAt(point{0, 5}).visN)
}

func Test_world_computeVisibilityFromSouthForCol(t *testing.T) {

	// 1x6 world  163742
	w := newWorld(0, 5)
	w.addTree(0, 0, 1)
	w.addTree(0, 1, 6)
	w.addTree(0, 2, 3)
	w.addTree(0, 3, 7)
	w.addTree(0, 4, 4)
	w.addTree(0, 5, 2)

	w.computeVisibilityFromSouthForCol(0)
	assert.False(t, w.forest.getTreeAt(point{0, 0}).visS)
	assert.False(t, w.forest.getTreeAt(point{0, 1}).visS)
	assert.False(t, w.forest.getTreeAt(point{0, 2}).visS)
	assert.True(t, w.forest.getTreeAt(point{0, 3}).visS)
	assert.True(t, w.forest.getTreeAt(point{0, 4}).visS)
	assert.True(t, w.forest.getTreeAt(point{0, 5}).visS)
}
