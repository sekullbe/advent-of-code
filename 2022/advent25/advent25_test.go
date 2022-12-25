package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleText string

func Test_snafuToInt(t *testing.T) {
	type args struct {
		snafu string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{"1"}, want: 1},
		{name: "2", args: args{"2"}, want: 2},
		{name: "3", args: args{"1="}, want: 3},
		{name: "4", args: args{"1-"}, want: 4},
		{name: "5", args: args{"10"}, want: 5},
		{name: "6", args: args{"11"}, want: 6},
		{name: "7", args: args{"12"}, want: 7},
		{name: "8", args: args{"2="}, want: 8},
		{name: "9", args: args{"2-"}, want: 9},
		{name: "10", args: args{"20"}, want: 10},
		{name: "15", args: args{"1=0"}, want: 15},
		{name: "20", args: args{"1-0"}, want: 20},
		{name: "2022", args: args{"1=11-2"}, want: 2022},
		{name: "12345", args: args{"1-0---0"}, want: 12345},
		{name: "31459265", args: args{"1121-1110-1=0"}, want: 314159265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snafuToInt(tt.args.snafu); got != tt.want {
				t.Errorf("snafuToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_log5(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{1.0}, want: 0.0},
		{name: "5", args: args{5.0}, want: 1.0},
		{name: "25", args: args{25.0}, want: 2.0},
		{name: "1/5", args: args{1.0 / 5.0}, want: -1.0},
		{name: "1/25", args: args{1.0 / 25.0}, want: -2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := log5(tt.args.n); got != tt.want {
				t.Errorf("log5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToSnafu(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantSnafu string
	}{
		{name: "0", args: args{0}, wantSnafu: "0"},
		{name: "1", args: args{1}, wantSnafu: "1"},
		{name: "2", args: args{2}, wantSnafu: "2"},
		{name: "3", args: args{3}, wantSnafu: "1="},
		{name: "4", args: args{4}, wantSnafu: "1-"},
		{name: "5", args: args{5}, wantSnafu: "10"},
		{name: "6", args: args{6}, wantSnafu: "11"},
		{name: "7", args: args{7}, wantSnafu: "12"},
		{name: "8", args: args{8}, wantSnafu: "2="},
		{name: "9", args: args{9}, wantSnafu: "2-"},
		{name: "10", args: args{10}, wantSnafu: "20"},
		{name: "15", args: args{15}, wantSnafu: "1=0"},
		{name: "20", args: args{20}, wantSnafu: "1-0"},
		{name: "2022", args: args{2022}, wantSnafu: "1=11-2"},
		{name: "12345", args: args{12345}, wantSnafu: "1-0---0"},
		{name: "314159265", args: args{314159265}, wantSnafu: "1121-1110-1=0"},
		{name: "1=-2-2", args: args{1747}, wantSnafu: "1=-0-2"},
		{name: "12111", args: args{906}, wantSnafu: "12111"},
		{name: "2=0=", args: args{198}, wantSnafu: "2=0="},
		{name: "111", args: args{31}, wantSnafu: "111"},
		{name: "20012", args: args{1257}, wantSnafu: "20012"},
		{name: "112", args: args{32}, wantSnafu: "112"},
		{name: "1=-1=", args: args{353}, wantSnafu: "1=-1="},
		{name: "1-12", args: args{107}, wantSnafu: "1-12"},
		{name: "122", args: args{37}, wantSnafu: "122"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSnafu := intToSnafu(tt.args.n); gotSnafu != tt.wantSnafu {
				t.Errorf("intToSnafu() = %v, want %v", gotSnafu, tt.wantSnafu)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "sample", args: args{sampleText}, want: "2=-1=0"},
		{name: "live", args: args{inputText}, want: "20=212=1-12=200=00-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
