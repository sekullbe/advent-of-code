package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleText string

func Test_run1BruteForce(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample text",
			args: args{
				input: sampleText,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1BruteForce(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
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
		{
			name: "sample text",
			args: args{
				input: sampleText,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSectionComputeOneSection(t *testing.T) {
	type args struct {
		key     int
		section mapSection
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple out of range", args: args{key: 97, section: mapSection{sourceStart: 98, destStart: 50, rangeLength: 2}}, want: 97},
		{name: "simple start of range", args: args{key: 98, section: mapSection{sourceStart: 98, destStart: 50, rangeLength: 2}}, want: 50},
		{name: "simple in range", args: args{key: 99, section: mapSection{sourceStart: 98, destStart: 50, rangeLength: 2}}, want: 51},
		{name: "simple past end of range", args: args{key: 100, section: mapSection{sourceStart: 98, destStart: 50, rangeLength: 2}}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapSectionComputeOneSection(tt.args.key, tt.args.section); got != tt.want {
				t.Errorf("mapSectionComputeOneSection() = %v, want %v", got, tt.want)
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
		{
			name: "sample text",
			args: args{
				input: sampleText,
			},
			want: 46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
