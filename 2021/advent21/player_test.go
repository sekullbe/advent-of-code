package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_player_move(t *testing.T) {
	p := newPlayer(1, 1)
	p.move(1)
	assert.Equal(t, 2, p.pos())
	assert.Equal(t, 2, p.score)

}
