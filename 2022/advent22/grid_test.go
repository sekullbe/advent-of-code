package main

import (
	"image"
	"testing"
)

func Test_coordsToSide(t *testing.T) {
	type args struct {
		p image.Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{image.Point{75, 25}}, want: 1},
		{name: "2", args: args{image.Point{125, 25}}, want: 2},
		{name: "3", args: args{image.Point{75, 75}}, want: 3},
		{name: "4", args: args{image.Point{23, 110}}, want: 4},
		{name: "5", args: args{image.Point{62, 114}}, want: 5},
		{name: "6", args: args{image.Point{12, 151}}, want: 6},
		{name: "1c", args: args{image.Point{51, 1}}, want: 1},
		{name: "2c", args: args{image.Point{101, 1}}, want: 2},
		{name: "3c", args: args{image.Point{51, 51}}, want: 3},
		{name: "4c", args: args{image.Point{1, 101}}, want: 4},
		{name: "5c", args: args{image.Point{51, 101}}, want: 5},
		{name: "6c", args: args{image.Point{1, 151}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coordsToSide(tt.args.p); got != tt.want {
				t.Errorf("coordsToSide() = %v, want %v", got, tt.want)
			}
		})
	}
}
