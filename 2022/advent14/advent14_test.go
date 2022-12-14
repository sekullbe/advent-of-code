package main

import (
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testinput = "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9\n"

// This just prints out the sample world for me to see
func Test_sample_world(t *testing.T) {

	w := parseRockVeins(parsers.SplitByLines(testinput))
	assert.Equal(t, 9, w.lowestRock)
	assert.Equal(t, 494, w.leftRock)
	assert.Equal(t, 503, w.rightRock)
	w.printWorld()
}

func Test_sample_world_drop(t *testing.T) {
	w := parseRockVeins(parsers.SplitByLines(testinput))
	for w.dropSand(point{500, 0}) {
		w.printWorld()
	}
}

func Test_real_world(t *testing.T) {

	w := parseRockVeins(parsers.SplitByLines(inputText))
	//assert.Equal(t, 9, w.lowestRock)
	//assert.Equal(t, 494, w.leftRock)
	//assert.Equal(t, 503, w.rightRock)
	w.printWorld()
}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{inputText: testinput}, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{inputText: testinput}, want: 93},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
