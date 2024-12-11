package main

import (
	"reflect"
	"testing"
)

const sample1 = `0 1 10 99 999`
const sample2 = `125 17`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{ name:"sample1", args: args{input: sample1}, want: 7},
		{name: "sample2", args: args{input: sample2}, want: 55312},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{name: "1", args: args{num: 1}, wantCount: 1},
		{name: "10", args: args{num: 10}, wantCount: 2},
		{name: "100", args: args{num: 100}, wantCount: 3},
		{name: "1000", args: args{num: 1000}, wantCount: 4},
		{name: "10000", args: args{num: 10000}, wantCount: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countDigits(tt.args.num); gotCount != tt.wantCount {
				t.Errorf("countDigits() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_splitDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  int
		wantRight int
	}{
		{name: "two", args: args{num: 11}, wantLeft: 1, wantRight: 1},
		{name: "four", args: args{num: 1234}, wantLeft: 12, wantRight: 34},
		{name: "one", args: args{num: 1}, wantLeft: 1, wantRight: 0},
		{name: "three", args: args{num: 123}, wantLeft: 1, wantRight: 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotRight := splitDigits(tt.args.num)
			if gotLeft != tt.wantLeft {
				t.Errorf("splitDigits() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if gotRight != tt.wantRight {
				t.Errorf("splitDigits() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}

func Test_blink(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "ex1", args: args{[]int{0, 1, 10, 99, 999}}, want: []int{1, 2024, 1, 0, 9, 9, 2021976}},
		{name: "ex2a", args: args{[]int{125, 17}}, want: []int{253000, 1, 7}},
		{name: "ex2b", args: args{[]int{253000, 1, 7}}, want: []int{253, 0, 2024, 14168}},
		{name: "ex2f", args: args{[]int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}}, want: []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blink(tt.args.stones); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreOneStone(t *testing.T) {

	stoneScores = make(map[stoneAndCount]int)
	type args struct {
		stone      int
		iterations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0-1", args: args{stone: 0, iterations: 1}, want: 1}, // 0 -> 1 -> 2024 -> 20 24 -> 2 0 2 4
		{name: "0-2", args: args{stone: 0, iterations: 2}, want: 1},
		{name: "0-3", args: args{stone: 0, iterations: 3}, want: 2},
		{name: "0-4", args: args{stone: 0, iterations: 4}, want: 4},
		{name: "0-5", args: args{stone: 0, iterations: 5}, want: 4},  // 2 0 2 4 -> 4048 1 4048 8096
		{name: "0-6", args: args{stone: 0, iterations: 6}, want: 7},  // 4048 1 4048 8096 -> 40 48 2024 40 48 80 96
		{name: "0-7", args: args{stone: 0, iterations: 7}, want: 14}, // 40 48 2024 40 48 80 96 -> 4 0 4 8 20 24 4 0 4 8 8 0 9 6
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOneStone(tt.args.stone, tt.args.iterations); got != tt.want {
				t.Errorf("scoreOneStone() = %v, want %v", got, tt.want)
			}
		})
	}
}
