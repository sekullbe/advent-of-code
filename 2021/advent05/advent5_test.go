package main

import (
	"reflect"
	"testing"
)

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{inputText: "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2"},
			want: 5,
		},
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
		{
			name: "Example",
			args: args{inputText: "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2\n"},
			//args: args{inputText: "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2\n1,1 -> 3,3\n9,7 -> 7,9\n"},
			want: 12,
		},
		{
			name: "8s",
			args: args{inputText: "0,9 -> 5,9\n8,0 -> 0,8\n"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pointsInLine(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			name: "diagonal 1",
			args: args{x1: 0, y1: 0, x2: 3, y2: 3},
			want: []point{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
		},
		{
			name: "diagonal 2",
			args: args{x1: 3, y1: 3, x2: 0, y2: 0},
			want: []point{{3, 3}, {2, 2}, {1, 1}, {0, 0}},
		},
		{
			name: "crazy 8s",
			args: args{x1: 8, y1: 0, x2: 0, y2: 8},
			want: []point{{8, 0}, {7, 1}, {6, 2}, {5, 3}, {4, 4}, {3, 5}, {2, 6}, {1, 7}, {0, 8}},
		},
		{
			name: "twos",
			args: args{x1: 2, y1: 2, x2: 2, y2: 1},
			want: []point{{2, 1}, {2, 2}},
		},
		{
			name: "6,4 -> 2,0",
			args: args{x1: 6, y1: 4, x2: 2, y2: 0},
			want: []point{{6, 4}, {5, 3}, {4, 2}, {3, 1}, {2, 0}},
		},
		{
			name: "5,5 -> 8,2",
			args: args{x1: 5, y1: 5, x2: 8, y2: 2},
			want: []point{{5, 5}, {6, 4}, {7, 3}, {8, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointsInLine(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pointsInLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
