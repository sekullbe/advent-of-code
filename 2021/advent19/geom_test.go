package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// not sure I need this
func Test_point_rotate(t *testing.T) {
	type args struct {
		pitch int
		roll  int
		yaw   int
	}
	tests := []struct {
		name  string
		point beacon
		args  args
		want  beacon
	}{
		{
			name:  "origin sanity check",
			point: beacon{0, 0, 0},
			args:  args{1, 2, 3},
			want:  beacon{0, 0, 0},
		},
		{
			name:  "simple pitch",
			point: beacon{0, 1, 0},
			args:  args{1, 0, 0},
			want:  beacon{0, 0, 1},
		},
		{
			name:  "simple pitch down",
			point: beacon{0, 1, 0},
			args:  args{-1, 0, 0},
			want:  beacon{0, 0, -1},
		},
		{
			name:  "simple roll",
			point: beacon{0, 0, 0},
			args:  args{0, 1, 1},
			want:  beacon{0, 0, 1},
		},
		{
			name:  "simple yaw",
			point: beacon{0, 0, 0},
			args:  args{0, 0, 0},
			want:  beacon{0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.point
			p.rotate(tt.args.pitch, tt.args.roll, tt.args.yaw)
			assert.Equal(t, tt.want, p)
		})
	}
}

func Test_manhattanDistance(t *testing.T) {
	type args struct {
		p1 beacon
		p2 beacon
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{p1: beacon{1105, -1205, 1229}, p2: beacon{-92, -2380, -20}},
			want: 3621,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, manhattanDistance(tt.args.p1, tt.args.p2), "manhattanDistance(%v, %v)", tt.args.p1, tt.args.p2)
		})
	}
}
