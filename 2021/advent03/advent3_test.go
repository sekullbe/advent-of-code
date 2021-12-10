package main

import (
	"strings"
	"testing"
)

func Test_setBitNFromLeft(t *testing.T) {
	type args struct {
		num    int
		place  int
		length int
		value  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic 00 -> 10",
			args: args{num: 0, place: 0, length: 2, value: 1},
			want: 2,
		},
		{
			name: "basic 01 -> 11",
			args: args{num: 1, place: 0, length: 2, value: 1},
			want: 3,
		},
		{
			name: "basic 01 -> 11",
			args: args{num: 1, place: 0, length: 2, value: 1},
			want: 3,
		},
		{
			name: "1001 -> 0001",
			args: args{num: 9, place: 0, length: 4, value: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setBitNFromLeft(tt.args.num, tt.args.place, tt.args.length, tt.args.value); got != tt.want {
				t.Errorf("setBitNFromLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateOxygen(t *testing.T) {

	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{nums: strings.Fields("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010")},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateOxygen(tt.args.nums); got != tt.want {
				t.Errorf("calculateOxygen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateScrubber(t *testing.T) {

	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{nums: strings.Fields("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010")},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScrubber(tt.args.nums); got != tt.want {
				t.Errorf("calculateScrubber() = %v, want %v", got, tt.want)
			}
		})
	}
}
