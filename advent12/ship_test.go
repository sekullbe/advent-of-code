package advent12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ship_rotate(t *testing.T) {
	s := ship{x:0, y:0, facing: N}

	s.rotate(true,0)
	assert.True(t, s.facing == N)

	s.rotate(true,90)
	assert.True(t, s.facing == E)

	s.rotate(false, 90)
	assert.True(t, s.facing == N)

	s.rotate(true,360)
	assert.True(t, s.facing == N)

	s.rotate(false,90)
	assert.Equal(t, W, s.facing)

	s.rotate(true,180)
	assert.True(t, s.facing == E)

	s.rotate(false,270)
	assert.Equal(t, S, s.facing)

	s.rotate(false,270)
	assert.Equal(t, W, s.facing)
}

func Test_ship_move(t *testing.T) {
	s := ship{x:0, y:0, facing: N}
	a := assert.New(t)

	s.move(N, 0)
	a.Equal(0, s.x)
	a.Equal(0, s.y)

	s.move(N, 10)
	a.Equal(0, s.x)
	a.Equal(10, s.y)

	s.move(E, 10)
	a.Equal(10, s.x)
	a.Equal(10, s.y)

	s.move(W, 20)
	a.Equal(-10, s.x)
	a.Equal(10, s.y)

	s.move(S, 30)
	a.Equal(-10, s.x)
	a.Equal(-20, s.y)
}

func Test_ship_moveForward(t *testing.T) {
	s := ship{x: 0, y: 0, facing: N}
	a := assert.New(t)

	s.moveForward(0)
	a.Equal(0, s.x)
	a.Equal(0, s.y)

	s.moveForward(100)
	a.Equal(0, s.x)
	a.Equal(100, s.y)

	s.rotate(true,90)
	s.moveForward(100)
	a.Equal(100, s.x)
	a.Equal(100, s.y)

	s.rotate(true, 270)
	s.moveForward(100)
	a.Equal(100, s.x)
	a.Equal(200, s.y)

	s.rotate(false, 180)
	s.moveForward(200)
	a.Equal(100, s.x)
	a.Equal(0, s.y)

}

