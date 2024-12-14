package main

import (
	"github.com/sekullbe/advent/geometry"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const sampleText = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func Test_run1(t *testing.T) {
	type args struct {
		input string
		maxX  int
		maxY  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: sampleText, maxX: 10, maxY: 6}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input, tt.args.maxX, tt.args.maxY); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRobot(t *testing.T) {
	type args struct {
		robLine string
	}
	tests := []struct {
		name string
		args args
		want robot
	}{
		{name: "simple", args: args{robLine: "p=0,4 v=3,-3"}, want: robot{pos: geometry.NewPoint2(0, 4), vel: geometry.NewPoint2(3, -3)}},
		{name: "simple2", args: args{robLine: "p=10,4 v=3,-3"}, want: robot{pos: geometry.NewPoint2(10, 4), vel: geometry.NewPoint2(3, -3)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRobot(tt.args.robLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRobot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveRobot(t *testing.T) {
	r := robot{
		pos: geometry.Point2{X: 2, Y: 4},
		vel: geometry.Point2{X: 2, Y: -3},
	}
	assert.Equal(t, r.pos, geometry.NewPoint2(2, 4))

	r.moveRobot(10, 6)
	assert.Equal(t, r.pos, geometry.NewPoint2(4, 1))

	r.moveRobot(10, 6)
	assert.Equal(t, r.pos, geometry.NewPoint2(6, 5))

	r.moveRobot(10, 6)
	assert.Equal(t, r.pos, geometry.NewPoint2(8, 2))

	r.moveRobot(10, 6)
	assert.Equal(t, r.pos, geometry.NewPoint2(1, 3))
}

func Test_computeQuadrant(t *testing.T) {
	type args struct {
		p    geometry.Point2
		maxX int
		maxY int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "nw", args: args{p: geometry.NewPoint2(1, 1), maxX: 10, maxY: 6}, want: NW},
		{name: "nw2", args: args{p: geometry.NewPoint2(0, 2), maxX: 10, maxY: 6}, want: NW},
		{name: "ne", args: args{p: geometry.NewPoint2(6, 1), maxX: 10, maxY: 6}, want: NE},
		{name: "se", args: args{p: geometry.NewPoint2(7, 5), maxX: 10, maxY: 6}, want: SE},
		{name: "sw", args: args{p: geometry.NewPoint2(2, 5), maxX: 10, maxY: 6}, want: SW},
		{name: "none", args: args{p: geometry.NewPoint2(1, 3), maxX: 10, maxY: 6}, want: -1},
		{name: "none2", args: args{p: geometry.NewPoint2(5, 1), maxX: 10, maxY: 6}, want: -1},
		{name: "none3", args: args{p: geometry.NewPoint2(5, 3), maxX: 10, maxY: 6}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, computeQuadrant(tt.args.p, tt.args.maxX, tt.args.maxY), "computeQuadrant(%v, %v, %v)", tt.args.p, tt.args.maxX, tt.args.maxY)
		})
	}
}
