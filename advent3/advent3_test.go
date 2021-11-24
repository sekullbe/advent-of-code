package advent3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func Test_test(t *testing.T) {
	foo := "foo"
	assert.Equal(t, foo, "foo")
	assert := assert.New(t)
	assert.Equal(foo, "foo")
}

func Test_parseLandscape(t *testing.T) {

	inputText := ".#.#.#\n......\n######"
	ls, colCount, rowCount := parseLandscape(inputText)

	assert.Equal(t, 6, colCount)
	assert.Equal(t, 3, rowCount)

	assert.NotNil(t, ls)
	assert.False(t,ls[0][0])
	assert.True(t,ls[0][1])
	assert.True(t,ls[0][3])
	assert.True(t,ls[0][5])

	assert.False(t, ls[1][0])
	assert.False(t, ls[1][5])

	assert.True(t, ls[2][0])
	assert.True(t, ls[2][1])
	assert.True(t, ls[2][2])
	assert.True(t, ls[2][3])
	assert.True(t, ls[2][4])
	assert.True(t, ls[2][5])

}

func Test_move(t *testing.T) {
	type args struct {
		x    int
		y    int
		maxX int
		dx   int
		dy   int
	}
	tests := []struct {
		name     string
		args     args
		wantNewX int
		wantNewY int
	}{
		{
			name: "sameRow1",
			args: args{x:0, y:0, maxX:10, dx: 1, dy:1},
			wantNewX: 1,
			wantNewY: 1,
		},
		{
			name: "sameRow3",
			args: args{x:0, y:0, maxX:10, dx: 3, dy:1},
			wantNewX: 3,
			wantNewY: 1,
		},
		{
			name: "endRow",
			args: args{x:3, y:0, maxX:5, dx: 2, dy:1},
			wantNewX: 5,
			wantNewY: 1,
		},
		{
			name: "wrapRow0",
			args: args{x:5, y:0, maxX:5, dx: 1, dy:1},
			wantNewX: 0,
			wantNewY: 1,
		},
		{
			name: "wrapRow1",
			args: args{x:5, y:0, maxX:5, dx: 2, dy:1},
			wantNewX: 1,
			wantNewY: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewX, gotNewY := move(tt.args.x, tt.args.y, tt.args.maxX, tt.args.dx, tt.args.dy)
			if gotNewX != tt.wantNewX {
				t.Errorf("move() gotNewX = %v, want %v", gotNewX, tt.wantNewX)
			}
			if gotNewY != tt.wantNewY {
				t.Errorf("move() gotNewY = %v, want %v", gotNewY, tt.wantNewY)
			}
		})
	}
}
