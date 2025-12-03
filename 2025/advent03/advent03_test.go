package main

import (
	"reflect"
	"testing"
)

func Test_parseBanks(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "simple", args: args{input: "123\n456"}, want: [][]int{{1, 2, 3}, {4, 5, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBanks(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBanks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxOutputFromBank(t *testing.T) {
	type args struct {
		bank []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "123456", args: args{bank: []int{1, 2, 3, 4, 5, 6}}, want: 56},
		{name: "4569288888", args: args{bank: []int{4, 5, 6, 9, 2, 8, 8, 8, 8, 8}}, want: 98},
		{name: "4569288888", args: args{bank: []int{4, 5, 6, 9, 2, 8, 8, 8, 8, 8}}, want: 98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOutputFromBank(tt.args.bank); got != tt.want {
				t.Errorf("maxOutputFromBank() = %v, want %v", got, tt.want)
			}
		})
	}
}

const sampletext = `987654321111111
811111111111119
234234234234278
818181911112111`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{sampletext}, want: 357},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxOutputFromBank2(t *testing.T) {
	type args struct {
		bank    []int
		howMany int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "987654321111111", args: args{bank: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, howMany: 12}, want: 987654321111},
		{name: "811111111111119", args: args{bank: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, howMany: 12}, want: 811111111119},
		{name: "234234234234278", args: args{bank: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, howMany: 12}, want: 434234234278},
		{name: "818181911112111", args: args{bank: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, howMany: 12}, want: 888911112111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOutputFromBank2(tt.args.bank, tt.args.howMany); got != tt.want {
				t.Errorf("maxOutputFromBank2() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{sampletext}, want: 3121910778619},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
