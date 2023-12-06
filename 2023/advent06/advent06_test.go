package main

import "testing"

const sampleData = `Time:      7  15   30
Distance:  9  40  200`

func Test_calculateDistance(t *testing.T) {
	type args struct {
		dur  int
		hold int
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
	}{
		{name: "0 dur", args: args{dur: 0, hold: 0}, wantDistance: 0},
		{name: "0/7", args: args{dur: 0, hold: 0}, wantDistance: 0},
		{name: "1/7", args: args{dur: 7, hold: 1}, wantDistance: 6},
		{name: "2/7", args: args{dur: 7, hold: 2}, wantDistance: 10},
		{name: "3/7", args: args{dur: 7, hold: 3}, wantDistance: 12},
		{name: "4/7", args: args{dur: 7, hold: 4}, wantDistance: 12},
		{name: "5/7", args: args{dur: 7, hold: 5}, wantDistance: 10},
		{name: "6/7", args: args{dur: 7, hold: 6}, wantDistance: 6},
		{name: "7/7", args: args{dur: 7, hold: 7}, wantDistance: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := calculateDistance(tt.args.hold, tt.args.dur); gotDistance != tt.wantDistance {
				t.Errorf("calculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}

func Test_evalOneRace(t *testing.T) {
	type args struct {
		r race
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample 1", args: args{r: race{duration: 7, recordDistance: 9}}, want: 4},
		{name: "sample 2", args: args{r: race{duration: 15, recordDistance: 40}}, want: 8},
		{name: "sample 3", args: args{r: race{duration: 30, recordDistance: 200}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalOneRace(tt.args.r); got != tt.want {
				t.Errorf("evalOneRace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evalOneRaceOptimized(t *testing.T) {
	type args struct {
		r race
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample 1", args: args{r: race{duration: 7, recordDistance: 9}}, want: 4},
		{name: "sample 2", args: args{r: race{duration: 15, recordDistance: 40}}, want: 8},
		{name: "sample 3", args: args{r: race{duration: 30, recordDistance: 200}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalOneRaceOptimized(tt.args.r); got != tt.want {
				t.Errorf("evalOneRaceOptimized() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{input: sampleData}, want: 288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
