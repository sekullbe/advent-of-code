package main

import (
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

const sampleText = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sampletext", args: args{input: sampleText}, want: 14},
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
		{name: "sampletext", args: args{input: sampleText}, want: 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAntinodes(t *testing.T) {
	type args struct {
		a geometry.Point
		b geometry.Point
	}
	tests := []struct {
		name  string
		args  args
		want  geometry.Point
		want1 geometry.Point
	}{
		{name: "1",
			args: args{a: geometry.Point{4, 3}, b: geometry.Point{5, 5}},
			want: geometry.Point{3, 1}, want1: geometry.Point{6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateAntinodes(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateAntinodes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("calculateAntinodes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_calculateAntinodesAggressively(t *testing.T) {
	type args struct {
		a     geometry.Point
		b     geometry.Point
		board *grid.Board
	}
	tests := []struct {
		name string
		args args
		want []geometry.Point
	}{
		{name: "1",
			args: args{a: geometry.Point{4, 3}, b: geometry.Point{5, 5}, board: grid.ParseBoard(parsers.SplitByLines(sampleText))},
			want: []geometry.Point{{5, 5}, {4, 3}, {3, 1}, {4, 3}, {3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAntinodesAggressively(tt.args.a, tt.args.b, tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateAntinodesAggressively() = %v, want %v", got, tt.want)
			}
		})
	}
}
