package main

import (
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

const simpleMap = `0123
1234
8965
9876`

const testMap = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: testMap}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: testMap}, want: 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_followTrail_Simple(t *testing.T) {
	b := grid.ParseBoard(parsers.SplitByLines(simpleMap))
	trailhead := grid.Pt(0, 0)
	peaks := make(map[geometry.Point]int)
	followTrail(b, trailhead, peaks, true)
	assert.Equal(t, 2, len(peaks))
}

func Test_followTrail_Sampletext(t *testing.T) {
	b := grid.ParseBoard(parsers.SplitByLines(testMap))
	trailhead := grid.Pt(2, 0)
	peaks := make(map[geometry.Point]int)
	followTrail(b, trailhead, peaks, true)
	assert.Equal(t, 5, len(peaks))
}
