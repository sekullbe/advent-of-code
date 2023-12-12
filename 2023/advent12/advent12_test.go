package main

import "testing"

func Test_validateRow(t *testing.T) {
	type args struct {
		n nomogramRow
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sample 1", args: args{n: nomogramRow{spring: "#.#.###", groups: []int{1, 1, 3}}}, want: true},
		{name: "sample 2", args: args{n: nomogramRow{spring: ".#...#....###.", groups: []int{1, 1, 3}}}, want: true},
		{name: "sample 3", args: args{n: nomogramRow{spring: ".#.###.#.######", groups: []int{1, 3, 1, 6}}}, want: true},
		{name: "sample 4", args: args{n: nomogramRow{spring: "####.#...#...", groups: []int{4, 1, 1}}}, want: true},
		{name: "sample 5", args: args{n: nomogramRow{spring: "#....######..#####.", groups: []int{1, 6, 5}}}, want: true},
		{name: "sample 6", args: args{n: nomogramRow{spring: ".###.##....#", groups: []int{3, 2, 1}}}, want: true},
		{name: "false 1", args: args{n: nomogramRow{spring: "#.###.##....#", groups: []int{3, 2, 1}}}, want: false},
		{name: "false 2", args: args{n: nomogramRow{spring: "#.###.##....#", groups: []int{3, 2, 1, 1}}}, want: false},
		{name: "false 3", args: args{n: nomogramRow{spring: "#.....#", groups: []int{1, 1, 1}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateRow(tt.args.n); got != tt.want {
				t.Errorf("validateRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateRowWithNumberDigits(t *testing.T) {
	type args struct {
		n nomogramRow
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sample 1", args: args{n: nomogramRow{spring: "1010111", groups: []int{1, 1, 3}}}, want: true},
		{name: "sample 2", args: args{n: nomogramRow{spring: "01000100001110", groups: []int{1, 1, 3}}}, want: true},
		{name: "sample 3", args: args{n: nomogramRow{spring: "010111010111111", groups: []int{1, 3, 1, 6}}}, want: true},
		{name: "sample 4", args: args{n: nomogramRow{spring: "1111010001000", groups: []int{4, 1, 1}}}, want: true},
		{name: "sample 5", args: args{n: nomogramRow{spring: "1000011111100111110", groups: []int{1, 6, 5}}}, want: true},
		{name: "sample 6", args: args{n: nomogramRow{spring: "011101100001", groups: []int{3, 2, 1}}}, want: true},
		{name: "false 1", args: args{n: nomogramRow{spring: "1011101100001", groups: []int{3, 2, 1}}}, want: false},
		{name: "false 2", args: args{n: nomogramRow{spring: "1011101100001", groups: []int{3, 2, 1, 1}}}, want: false},
		{name: "false 3", args: args{n: nomogramRow{spring: "1000001", groups: []int{1, 1, 1}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateRowWithNumberDigits(tt.args.n); got != tt.want {
				t.Errorf("validateRowWithNumberDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPossibleArrangements(t *testing.T) {
	type args struct {
		n nomogramRow
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "trivial", args: args{n: nomogramRow{spring: ".?.", groups: []int{1}}}, want: 1},
		{name: "sample1", args: args{n: nomogramRow{spring: "???.###", groups: []int{1, 1, 3}}}, want: 1},
		{name: "sample2", args: args{n: nomogramRow{spring: ".??..??...?##.", groups: []int{1, 1, 3}}}, want: 4},
		{name: "sample3", args: args{n: nomogramRow{spring: "?#?#?#?#?#?#?#?", groups: []int{1, 3, 1, 6}}}, want: 1},
		{name: "sample4", args: args{n: nomogramRow{spring: "????.#...#...", groups: []int{4, 1, 1}}}, want: 1},
		{name: "sample5", args: args{n: nomogramRow{spring: "????.######..#####.", groups: []int{1, 6, 5}}}, want: 4},
		{name: "sample6", args: args{n: nomogramRow{spring: "?###????????", groups: []int{3, 2, 1}}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPossibleArrangements(tt.args.n); got != tt.want {
				t.Errorf("countPossibleArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
