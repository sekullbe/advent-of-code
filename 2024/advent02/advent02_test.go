package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed sample.txt
var sampleText string

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sampletext", args: args{input: sampleText}, want: 2},
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
		{name: "sampletext", args: args{input: sampleText}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseReports(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "simple", args: args{input: "1 2 3 4 5\n2 4 6\n7 1 2\n"}, want: [][]int{{1, 2, 3, 4, 5}, {2, 4, 6}, {7, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseReports(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testallOneDirection(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "inc", args: args{s: []int{1, 2, 3}}, want: true},
		{name: "dec", args: args{s: []int{3, 2, 1}}, want: true},
		{name: "inceq", args: args{s: []int{1, 3, 5, 5, 7, 9}}, want: false},
		{name: "deceq", args: args{s: []int{9, 7, 6, 5, 5, 5, 4, 3}}, want: false},
		{name: "incno", args: args{s: []int{1, 3, 5, 6, 7, 9, 8}}, want: false},
		{name: "decno", args: args{s: []int{10, 9, 8, 7, 6, 5, 4, 3, 4, 1}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testallOneDirection(tt.args.s); got != tt.want {
				t.Errorf("testallOneDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testDiff(t *testing.T) {
	type args struct {
		s   []int
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "inc1", args: args{s: []int{1, 2, 3, 4, 5}, min: 1, max: 3}, want: true},
		{name: "inc1", args: args{s: []int{1, 3, 5, 7, 9}, min: 1, max: 3}, want: true},
		{name: "inc3", args: args{s: []int{1, 4, 7, 10, 13}, min: 1, max: 3}, want: true},
		{name: "inc3+1", args: args{s: []int{1, 4, 7, 11, 13}, min: 1, max: 3}, want: false},
		{name: "dec3", args: args{s: []int{13, 10, 9, 8, 7, 4, 1}, min: 1, max: 3}, want: true},
		{name: "upanddown", args: args{s: []int{1, 2, 3, 4, 5, 4, 3, 1}, min: 1, max: 3}, want: true},
		{name: "upanddown+eq", args: args{s: []int{1, 2, 3, 4, 5, 5, 4, 3, 1}, min: 1, max: 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testDiff(tt.args.s, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("testDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateDampenedReports(t *testing.T) {
	type args struct {
		r []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "inc3", args: args{r: []int{1, 4, 7, 10, 13}},
			want: [][]int{
				{4, 7, 10, 13},
				{1, 7, 10, 13},
				{1, 4, 10, 13},
				{1, 4, 7, 13},
				{1, 4, 7, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateDampenedReports(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateDampenedReports() = %v, want %v", got, tt.want)
			}
		})
	}
}
