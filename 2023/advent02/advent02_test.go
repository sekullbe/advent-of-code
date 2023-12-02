package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleText string

func Test_extractnum(t *testing.T) {
	type args struct {
		color         string
		contentString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "green", args: args{color: "green", contentString: "8 green, 6 blue, 20 red"}, want: 8},
		{name: "blue", args: args{color: "blue", contentString: "8 green, 6 blue, 20 red"}, want: 6},
		{name: "red", args: args{color: "red", contentString: "8 green, 6 blue, 20 red"}, want: 20},
		{name: "missing", args: args{color: "fuschia", contentString: "8 green, 6 blue, 20 red"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractnum(tt.args.color, tt.args.contentString); got != tt.want {
				t.Errorf("extractnum() = %v, want %v", got, tt.want)
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
		want int
	}{
		{
			name: "sampletext",
			args: args{
				inputText: sampleText,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
