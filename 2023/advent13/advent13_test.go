package main

import (
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

const pattern1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

const pattern2 = `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func Test_findMirrorPoints(t *testing.T) {

	p1 := parsers.SplitByLines(pattern1)
	p2 := parsers.SplitByLines(pattern2)
	p2r := rotatePattern(p2)
	fmt.Println(p2r)

	type args struct {
		row  string
		maxr int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "sample1", args: args{row: p1[0]}, want: []int{5, 7}},
		{name: "sample1 bare", args: args{row: "#.##..##."}, want: []int{5, 7}},
		{name: "sample2", args: args{row: p1[1]}, want: []int{1, 5}},
		{name: "sample3", args: args{row: p1[2]}, want: []int{1, 5}},
		{name: "sample4", args: args{row: p1[3]}, want: []int{1, 5}},
		{name: "sample5", args: args{row: p1[4]}, want: []int{1, 5}},
		{name: "sample6", args: args{row: p1[5]}, want: []int{1, 3, 5, 7}},
		{name: "sample7", args: args{row: p1[6]}, want: []int{5}},
		{name: "sample21r", args: args{row: rotatePattern(p2)[0]}, want: []int{1, 4, 6}},
		{name: "sample21", args: args{row: p2r[0]}, want: []int{1, 4, 6}},
		{name: "sample21 bare", args: args{row: "###..##"}, want: []int{1, 4, 6}},
		{name: "sample22", args: args{row: p2r[1]}, want: []int{1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMirrorPoints(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMirrorPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotatePattern(t *testing.T) {
	type args struct {
		pattern []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "2x2", args: args{pattern: []string{"ab", "cd"}}, want: []string{"bd", "ac"}},
		{name: "2x4", args: args{pattern: []string{"ab", "cd", "ef", "gh"}}, want: []string{"bdfh", "aceg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotatePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotatePattern() = %v, want %v", got, tt.want)
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
		{name: "sample1", args: args{input: pattern1 + "\n\n" + pattern2}, want: 405},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMirrorPoint(t *testing.T) {
	type args struct {
		pattern []string
	}
	tests := []struct {
		name         string
		args         args
		wantMp       int
		wantVertical bool
	}{
		{name: "sample1", args: args{pattern: parsers.SplitByLines(pattern1)}, wantMp: 5, wantVertical: true},
		{name: "sample2", args: args{pattern: parsers.SplitByLines(pattern2)}, wantMp: 4, wantVertical: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMp, gotVertical := findMirrorPoint(tt.args.pattern)
			if gotMp != tt.wantMp {
				t.Errorf("findMirrorPoint() gotMp = %v, want %v", gotMp, tt.wantMp)
			}
			if gotVertical != tt.wantVertical {
				t.Errorf("findMirrorPoint() gotVertical = %v, want %v", gotVertical, tt.wantVertical)
			}
		})
	}
}
