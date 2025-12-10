package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleText = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func Test_parseFinalState(t *testing.T) {
	type args struct {
		state string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "000", args: args{state: "[...]"}, want: 0},
		{name: "001", args: args{state: "[..#]"}, want: 1},
		{name: "010", args: args{state: "[.#.]"}, want: 2},
		{name: "011", args: args{state: "[.##]"}, want: 3},
		{name: "0110", args: args{state: "[.##.]"}, want: 6},
		{name: "11111111", args: args{state: "[########]"}, want: 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFinalState(tt.args.state); got != tt.want {
				t.Errorf("parseFinalState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{sampleText}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pushAButtonBitmap(t *testing.T) {
	type args struct {
		state     int64
		numLights int
		button    []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "simple", args: args{6, 4, []int{0}}, want: 14},
		{name: "simple", args: args{6, 4, []int{1}}, want: 2},
		{name: "simple", args: args{6, 4, []int{2}}, want: 4},
		{name: "simple", args: args{6, 4, []int{3}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushAButtonBitmap(tt.args.state, tt.args.numLights, tt.args.button); got != tt.want {
				t.Errorf("pushAButtonBitmap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RepeatedPressesActuallyWork(t *testing.T) {

	var state int64 = 0
	state = pushAButtonBitmap(state, 4, []int{0, 2})
	state = pushAButtonBitmap(state, 4, []int{0, 1})
	assert.Equal(t, state, int64(6))

}
