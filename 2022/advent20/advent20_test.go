package main

import (
	"reflect"
	"testing"
)

func Test_computeNewIndex(t *testing.T) {
	type args struct {
		s   []int
		idx int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 0}, want: 0},
		{name: "small", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 5}, want: 5},
		{name: "larger", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 10}, want: 0},
		{name: "larger still", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 30}, want: 0},
		{name: "larger than that ", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 35}, want: 5},
		{name: "negative", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: -1}, want: 9},
		{name: "-10", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: -10}, want: 0},
		{name: "-11", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: -11}, want: 9},
		{name: "-100", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: -100}, want: 0},
		{name: "negative and a half", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: -15}, want: 5},
		{name: "sample -3", args: args{s: []int{1, -3, 2, 3, -2, 0, 4}, idx: -3}, want: 4},
		{name: "sample 4 right", args: args{s: []int{1, 2, -3, 0, 3, 4, -2}, idx: 10}, want: 3},
		{name: "loop around sample 0", args: args{s: []int{1, 2, -3, 0, 3, 4, -2}, idx: 49}, want: 0},
		{name: "loop around sample 1", args: args{s: []int{1, 2, -3, 0, 3, 4, -2}, idx: 50}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeNewIndex(tt.args.s, tt.args.idx); got != tt.want {
				t.Errorf("computeNewIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateElt(t *testing.T) {
	type args struct {
		s    []int
		idx  int
		move int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "still", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 3, move: 0}, want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{name: "simple +1", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 3, move: 1}, want: []int{9, 8, 7, 5, 6, 4, 3, 2, 1, 0}},
		{name: "simple +2", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 3, move: 2}, want: []int{9, 8, 7, 5, 4, 6, 3, 2, 1, 0}},
		{name: "simple -1", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 3, move: -1}, want: []int{9, 8, 6, 7, 5, 4, 3, 2, 1, 0}},
		{name: "simple -2", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 3, move: -2}, want: []int{9, 6, 8, 7, 5, 4, 3, 2, 1, 0}},
		{name: "round the right", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 9, move: 1}, want: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{name: "left 1-1", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 1, move: -1}, want: []int{8, 9, 7, 6, 5, 4, 3, 2, 1, 0}},
		{name: "left 0-1", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 0, move: -1}, want: []int{8, 7, 6, 5, 4, 3, 2, 1, 0, 9}},
		{name: "round the left -3", args: args{s: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, idx: 1, move: -3}, want: []int{9, 7, 6, 5, 4, 3, 2, 8, 1, 0}},
		{name: "round the left -3 sample", args: args{s: []int{1, -3, 2, 3, -2, 0, 4}, idx: 1, move: -3}, want: []int{1, 2, 3, -2, -3, 0, 4}},
		{name: "round the right 4 sample", args: args{s: []int{1, 2, -3, 0, 3, 4, -2}, idx: 5, move: 4}, want: []int{1, 2, -3, 4, 0, 3, -2}},
		{name: "round the right 7 sample", args: args{s: []int{1, 2, -3, 0, 3, 4, -2}, idx: 5, move: 7}, want: []int{1, 2, -3, 0, 3, 4, -2}},
		{name: "round the right 7 sampleish", args: args{s: []int{1, 2, -3, 0, 3, 21, -2}, idx: 5, move: 7}, want: []int{1, 2, -3, 0, 3, 21, -2}},
		{name: "round the right sample last elt", args: args{s: []int{1, 2, -3, 0, 3, 4, -2}, idx: 5, move: 4}, want: []int{1, 2, -3, 4, 0, 3, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateElt(tt.args.s, tt.args.idx, tt.args.move); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateElt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mix(t *testing.T) {
	// make a sample input
	testInput := []int{1, 2, -3, 3, -2, 0, 4}
	input := []*int{}
	for _, i2 := range testInput {
		n := i2
		input = append(input, &n)
	}

	type args struct {
		input []*int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "example", args: args{input: input}, want: []int{1, 2, -3, 4, 0, 3, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mix(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	var sampleText = "1 2 -3 3 -2 0 4"

	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{sampleText}, want: 3},
		{name: "real", args: args{inputText}, want: 8372},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floormod(t *testing.T) {
	type args struct {
		n   int
		mod int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "-1%4", args: args{-1, 4}, want: 3},
		{name: "-5%4", args: args{-5, 4}, want: 3},
		{name: "4%4", args: args{4, 4}, want: 0},
		{name: "5%4", args: args{5, 4}, want: 1},
		{name: "39%4", args: args{39, 4}, want: 3},
		{name: "40%4", args: args{40, 4}, want: 0},
		{name: "41%4", args: args{41, 4}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floormod(tt.args.n, tt.args.mod); got != tt.want {
				t.Errorf("floormod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	var sampleText = "1 2 -3 3 -2 0 4"

	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{sampleText}, want: 1623178306},
		{name: "real", args: args{inputText}, want: 8372},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
