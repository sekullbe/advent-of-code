package main

import (
	"reflect"
	"testing"
)

const sampleText = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  []idrange
		want1 []int
	}{
		{name: "sample", args: args{sampleText}, want: []idrange{{3, 5}, {10, 14}, {16, 20}, {12, 18}}, want1: []int{1, 5, 8, 11, 17, 32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInput(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
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
		{name: "sample", args: args{sampleText}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeRanges(t *testing.T) {
	type args struct {
		a idrange
		b idrange
	}
	tests := []struct {
		name string
		args args
		want idrange
	}{
		//{name: "disjoint", args: args{idrange{5, 10}, idrange{11, 20}}, want: false},
		{name: "overlap right", args: args{idrange{5, 10}, idrange{9, 20}}, want: idrange{5, 20}},
		{name: "overlap left", args: args{idrange{5, 10}, idrange{1, 6}}, want: idrange{1, 10}},
		{name: "touching", args: args{idrange{5, 10}, idrange{10, 20}}, want: idrange{5, 20}},
		{name: "identical", args: args{idrange{5, 10}, idrange{5, 10}}, want: idrange{5, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeRanges(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rangesOverlap(t *testing.T) {
	type args struct {
		a idrange
		b idrange
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "disjoint", args: args{idrange{5, 10}, idrange{11, 20}}, want: false},
		{name: "overlap", args: args{idrange{5, 10}, idrange{9, 20}}, want: true},
		{name: "touching", args: args{idrange{5, 10}, idrange{10, 20}}, want: true},
		{name: "identical", args: args{idrange{5, 10}, idrange{5, 10}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangesOverlap(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("rangesOverlap() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{sampleText}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
