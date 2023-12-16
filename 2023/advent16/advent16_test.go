package main

import "testing"

const splitvert = `
.\..
.\.|
....`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "split_vert", args: args{input: splitvert}, want: 7},
		{name: "sample", args: args{input: sample}, want: 46},
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorLeft(t *testing.T) {
	type args struct {
		dir int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "south", args: args{dir: SOUTH}, want: WEST},
		{name: "west", args: args{dir: WEST}, want: SOUTH},
		{name: "north", args: args{dir: NORTH}, want: EAST},
		{name: "east", args: args{dir: EAST}, want: NORTH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorLeft(tt.args.dir); got != tt.want {
				t.Errorf("mirrorLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorRight(t *testing.T) {
	type args struct {
		dir int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "south", args: args{dir: SOUTH}, want: EAST},
		{name: "west", args: args{dir: WEST}, want: NORTH},
		{name: "north", args: args{dir: NORTH}, want: WEST},
		{name: "east", args: args{dir: EAST}, want: SOUTH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorRight(tt.args.dir); got != tt.want {
				t.Errorf("mirrorRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
