package main

import (
	"reflect"
	"testing"

	"github.com/sekullbe/advent/geometry"
)

const sampleText = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

// per reddit, if this is 84 code looks outside polygon; 180 is "really" outside polygon
const hpolygon = `1,0
3,0
3,5
16,5
16,0
18,0
18,9
13,9
13,7
6,7
6,9
1,9`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{sampleText}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allTilesInRectangle(t *testing.T) {
	type args struct {
		t1 geometry.Point2
		t2 geometry.Point2
	}
	tests := []struct {
		name string
		args args
		want []geometry.Point2
	}{
		{name: "basic", args: args{t1: geometry.Point2{1, 1}, t2: geometry.Point2{2, 2}}, want: []geometry.Point2{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{name: "backwards", args: args{t1: geometry.Point2{10, 10}, t2: geometry.Point2{9, 9}},
			want: []geometry.Point2{{9, 9}, {9, 10}, {10, 9}, {10, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allTilesInRectangle(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allTilesInRectangle() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{sampleText}, want: 24},
		{name: "h", args: args{hpolygon}, want: 33},
		//{name: "real", args: args{inputText}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allPointsOnRectangleBorder(t *testing.T) {
	type args struct {
		t1 geometry.Point2
		t2 geometry.Point2
	}
	tests := []struct {
		name string
		args args
		want []geometry.Point2
	}{
		{name: "basic", args: args{t1: geometry.Point2{1, 1}, t2: geometry.Point2{2, 2}}, want: []geometry.Point2{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{name: "backwards", args: args{t1: geometry.Point2{10, 10}, t2: geometry.Point2{9, 9}},
			want: []geometry.Point2{{9, 9}, {9, 10}, {10, 9}, {10, 10}}},
		{name: "line", args: args{
			t1: geometry.Point2{0, 0},
			t2: geometry.Point2{0, 3}},
			want: []geometry.Point2{{0, 0}, {0, 3}, {0, 1}, {0, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPointsOnRectangleBorder(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPointsOnRectangleBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
