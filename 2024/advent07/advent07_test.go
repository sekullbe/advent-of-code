package main

import "testing"

const sampleText = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sampletext", args: args{input: sampleText}, want: 3749},
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
		{name: "sampletext", args: args{input: sampleText}, want: 11387},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvable(t *testing.T) {
	type args struct {
		e equation
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sample1", args: args{e: equation{190, []int{10, 19}}}, want: true},
		{name: "sample2", args: args{e: equation{3267, []int{81, 40, 27}}}, want: true},
		{name: "sample3", args: args{e: equation{83, []int{17, 5}}}, want: false},
		{name: "sample4", args: args{e: equation{156, []int{15, 6}}}, want: false},
		{name: "sample5", args: args{e: equation{7290, []int{6, 8, 6, 15}}}, want: false},
		{name: "sample6", args: args{e: equation{161011, []int{16, 10, 13}}}, want: false},
		{name: "sample7", args: args{e: equation{192, []int{17, 8, 14}}}, want: false},
		{name: "sample8", args: args{e: equation{21037, []int{9, 7, 18, 13}}}, want: false},
		{name: "sample9", args: args{e: equation{292, []int{11, 6, 16, 20}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvable(tt.args.e); got != tt.want {
				t.Errorf("solvable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvableWithConcat(t *testing.T) {
	type args struct {
		e equation
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sample1", args: args{e: equation{190, []int{10, 19}}}, want: true},
		{name: "sample2", args: args{e: equation{3267, []int{81, 40, 27}}}, want: true},
		{name: "sample3", args: args{e: equation{83, []int{17, 5}}}, want: false},
		{name: "sample4", args: args{e: equation{156, []int{15, 6}}}, want: true},
		{name: "sample5", args: args{e: equation{7290, []int{6, 8, 6, 15}}}, want: true},
		{name: "sample6", args: args{e: equation{161011, []int{16, 10, 13}}}, want: false},
		{name: "sample7", args: args{e: equation{192, []int{17, 8, 14}}}, want: true},
		{name: "sample8", args: args{e: equation{21037, []int{9, 7, 18, 13}}}, want: false},
		{name: "sample9", args: args{e: equation{292, []int{11, 6, 16, 20}}}, want: true},
		{name: "badconcat", args: args{e: equation{112387, []int{4, 4, 8, 1, 9, 8, 32, 2, 6, 1, 1, 601}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvableWithConcat(tt.args.e); got != tt.want {
				t.Errorf("solvableWithConcat() = %v, want %v", got, tt.want)
			}
		})
	}
}
