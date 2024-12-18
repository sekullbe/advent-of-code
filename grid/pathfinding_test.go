package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_FindPath(t *testing.T) {
	boardString := `...#...
..#..#.
....#..
...#..#
..#..#.
.#..#..
#.#....`

	b := ParseBoardString(boardString)
	path, pathLen, found := b.FindPath(Pt(0, 0), Pt(6, 6))
	assert.Equal(t, 23, len(path)) // 22+1 because it counts the 0,0 space
	assert.Equal(t, 22, pathLen)   // 22 steps
	assert.True(t, found)

}
