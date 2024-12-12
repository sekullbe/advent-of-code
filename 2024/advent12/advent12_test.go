package main

import "testing"

const sampleText1 = `AAAA
BBCD
BBCC
EEEC`

const sampleText2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

const sampleText3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const sampleTextE = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: sampleText1}, want: 140},
		{name: "sample2", args: args{input: sampleText2}, want: 772},
		{name: "sample3", args: args{input: sampleText3}, want: 1930},
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
		{name: "sample1", args: args{input: sampleText1}, want: 80},
		{name: "sample2", args: args{input: sampleText2}, want: 436},
		{name: "sampleE", args: args{input: sampleTextE}, want: 236},
		{name: "sample3", args: args{input: sampleText3}, want: 1206},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
