package main

import "testing"

func Test_calcCalibrationDigitsOnly(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1abc2", args: args{line: "1abc2"}, want: 12},
		{name: "pqr3stu8vwx", args: args{line: "pqr3stu8vwx"}, want: 38},
		{name: "a1b2c3d4e5f", args: args{line: "a1b2c3d4e5f"}, want: 15},
		{name: "treb7uchet", args: args{line: "treb7uchet"}, want: 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcCalibrationDigitsOnly(tt.args.line); got != tt.want {
				t.Errorf("calcCalibrationDigitsOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "two1nine", args: args{inputText: "two1nine"}, want: 29},
		{name: "eightwothree", args: args{inputText: "eightwothree"}, want: 83},
		{name: "abcone2threexyz", args: args{inputText: "abcone2threexyz"}, want: 13},
		{name: "xtwone3four", args: args{inputText: "xtwone3four"}, want: 24},
		{name: "4nineeightseven2", args: args{inputText: "4nineeightseven2"}, want: 42},
		{name: "zoneight234", args: args{inputText: "zoneight234"}, want: 14},
		{name: "7pqrstsixteen", args: args{inputText: "7pqrstsixteen"}, want: 76},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
