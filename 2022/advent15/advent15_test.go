package main

import "testing"

var testinput string = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

func Test_run1_countingOnly(t *testing.T) {
	type args struct {
		inputText string
		yToCheck  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{inputText: testinput, yToCheck: 10}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1_countingOnly(tt.args.inputText, tt.args.yToCheck); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1_withGrid(t *testing.T) {
	type args struct {
		inputText string
		yToCheck  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{inputText: testinput, yToCheck: 10}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1_withGrid(tt.args.inputText, tt.args.yToCheck); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
