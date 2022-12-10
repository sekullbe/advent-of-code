package main

import (
	"testing"
)

func Test_adjacent(t *testing.T) {
	type args struct {
		p1 point
		p2 point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "same0", args: args{p1: point{0, 0}, p2: point{0, 0}}, want: true},
		{name: "samepos", args: args{p1: point{3, 3}, p2: point{3, 3}}, want: true},
		{name: "sameneg", args: args{p1: point{-4, -4}, p2: point{-4, -4}}, want: true},
		{name: "samemix", args: args{p1: point{1, -2}, p2: point{1, -2}}, want: true},
		{name: "adjx+", args: args{p1: point{1, 1}, p2: point{2, 1}}, want: true},
		{name: "adjx-", args: args{p1: point{1, 1}, p2: point{2, 1}}, want: true},
		{name: "adjy+", args: args{p1: point{1, 1}, p2: point{1, 2}}, want: true},
		{name: "adjy-", args: args{p1: point{1, 1}, p2: point{1, 0}}, want: true},
		{name: "adjx+y+", args: args{p1: point{2, 2}, p2: point{3, 3}}, want: true},
		{name: "adjx-y+", args: args{p1: point{2, 2}, p2: point{1, 3}}, want: true},
		{name: "adjx+y-", args: args{p1: point{2, 2}, p2: point{3, 1}}, want: true},
		{name: "adjx-y-", args: args{p1: point{2, 2}, p2: point{1, 1}}, want: true},
		{name: "notadj1", args: args{p1: point{2, 2}, p2: point{0, 2}}, want: false},
		{name: "notadj2", args: args{p1: point{2, 2}, p2: point{2, 0}}, want: false},
		{name: "notadj3", args: args{p1: point{2, 2}, p2: point{0, 4}}, want: false},
		{name: "notadj4", args: args{p1: point{2, 2}, p2: point{4, 0}}, want: false},
		{name: "notadj5", args: args{p1: point{2, 2}, p2: point{0, 0}}, want: false},
		{name: "notadj6", args: args{p1: point{2, 2}, p2: point{4, 4}}, want: false},
		{name: "notadj7", args: args{p1: point{2, 2}, p2: point{30000, -12343}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjacent(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("adjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
