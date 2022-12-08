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
	assert.Equal(t, 0, w.forest.getTreeAt(point{0, 0}).sW) // we don't calculate this for edges
	assert.True(t, w.forest.getTreeAt(point{1, 0}).visW)
	assert.Equal(t, 1, w.forest.getTreeAt(point{1, 0}).sW)
	assert.True(t, w.forest.getTreeAt(point{2, 0}).visW)
	assert.Equal(t, 2, w.forest.getTreeAt(point{2, 0}).sW)
	assert.False(t, w.forest.getTreeAt(point{3, 0}).visW)
	assert.Equal(t, 1, w.forest.getTreeAt(point{3, 0}).sW)
	assert.True(t, w.forest.getTreeAt(point{4, 0}).visW)
	assert.Equal(t, 4, w.forest.getTreeAt(point{4, 0}).sW)
	assert.False(t, w.forest.getTreeAt(point{5, 0}).visW)
	assert.Equal(t, 0, w.forest.getTreeAt(point{5, 0}).sW) // we don't calculate this for edges
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
	assert.Equal(t, 1, w.forest.getTreeAt(point{2, 0}).sE)

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
	assert.Equal(t, 2, w.forest.getTreeAt(point{0, 1}).sS)
	assert.False(t, w.forest.getTreeAt(point{0, 2}).visS)
	assert.True(t, w.forest.getTreeAt(point{0, 3}).visS)
	assert.True(t, w.forest.getTreeAt(point{0, 4}).visS)
	assert.True(t, w.forest.getTreeAt(point{0, 5}).visS)
}

func Test_lastSeenTreeOfEachHeight_findNearestTallTree(t *testing.T) {
	type args struct {
		height int
	}
	tests := []struct {
		name string
		lst  lastSeenTreeOfEachHeight
		args args
		want int
	}{
		// 1635642, we're at the 6, x=4
		{name: "1", args: args{6}, lst: lastSeenTreeOfEachHeight{1: 0, 6: 1, 3: 2, 5: 3}, want: 1}, // want 1, to get distance
		// 344121134, we're at the last 4,
		{name: "2", args: args{4}, lst: lastSeenTreeOfEachHeight{3: 7, 4: 2, 1: 6, 2: 4}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.lst.findNearestTallTreeFromNorthOrWest(tt.args.height), "findNearestTallTree(%v)", tt.args.height)
		})
	}
}
