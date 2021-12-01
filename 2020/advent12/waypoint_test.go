package advent12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_waypoint_rotate(t *testing.T) {
	// R (x,y) -> (y,-x) -> (-x,-y) -> (-y,x)
	// L (x,y) -> (-y,x) -> (-x,-y) -> (y,-x)

	w := waypoint{x:0, y:0}
	w.rotate(true, 90)
	assert.Equal(t, 0, w.x)
	assert.Equal(t, 0, w.y)

	w = waypoint{x:10, y:4}
	w.rotate(true, 0)
	assert.Equal(t, 10, w.x)
	assert.Equal(t, 4, w.y)

	w = waypoint{x:10, y:4}
	w.rotate(true, 90)
	assert.Equal(t, 4, w.x)
	assert.Equal(t, -10, w.y)

	w = waypoint{x:2, y:3}
	w.rotate(true, 180)
	assert.Equal(t, -2, w.x)
	assert.Equal(t, -3, w.y)

	w = waypoint{x:2, y:3}
	w.rotate(true, 270)
	assert.Equal(t, -3, w.x)
	assert.Equal(t, 2, w.y)

	w = waypoint{x:2, y:3}
	w.rotate(false, 0)
	assert.Equal(t, 2, w.x)
	assert.Equal(t, 3, w.y)

	w = waypoint{x:2, y:3}
	w.rotate(false, 90)
	assert.Equal(t, -3, w.x)
	assert.Equal(t, 2, w.y)

	w = waypoint{x:2, y:3}
	w.rotate(false, 180)
	assert.Equal(t, -2, w.x)
	assert.Equal(t, -3, w.y)

	w = waypoint{x:2, y:3}
	w.rotate(false, 270)
	assert.Equal(t, 3, w.x)
	assert.Equal(t,-2, w.y)
}

func Test_waypoint_move(t *testing.T) {
	w := waypoint{x: 0, y:0}
	a := assert.New(t)

	w.move(N, 0)
	a.Equal(0, w.x)
	a.Equal(0, w.y)

	w.move(N, 10)
	a.Equal(0, w.x)
	a.Equal(10, w.y)

	w.move(E, 10)
	a.Equal(10, w.x)
	a.Equal(10, w.y)

	w.move(W, 20)
	a.Equal(-10, w.x)
	a.Equal(10, w.y)

	w.move(S, 30)
	a.Equal(-10, w.x)
	a.Equal(-20, w.y)

}
