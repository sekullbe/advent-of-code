package main

import "testing"

var testinput = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "examplepart1", args: args{testinput}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adjacent(t *testing.T) {
	type args struct {
		a point3
		b point3
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "adjx+", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 2, y: 1, z: 1}}, want: true},
		{name: "adjx-", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 0, y: 1, z: 1}}, want: true},
		{name: "adjy+", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 2, z: 1}}, want: true},
		{name: "adjy-", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 0, z: 1}}, want: true},
		{name: "adjz+", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 1, z: 2}}, want: true},
		{name: "adjz-", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 1, z: 0}}, want: true},
		{name: "2 apart", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 3, y: 1, z: 1}}, want: false},
		{name: "up and over", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 2, y: 2, z: 1}}, want: false},
		{name: "corner", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 2, y: 2, z: 2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjacent(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("adjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adjacent2(t *testing.T) {
	type args struct {
		a point3
		b point3
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "adjx+", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 2, y: 1, z: 1}}, want: true},
		{name: "adjx-", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 0, y: 1, z: 1}}, want: true},
		{name: "adjy+", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 2, z: 1}}, want: true},
		{name: "adjy-", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 0, z: 1}}, want: true},
		{name: "adjz+", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 1, z: 2}}, want: true},
		{name: "adjz-", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 1, y: 1, z: 0}}, want: true},
		{name: "2 apart", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 3, y: 1, z: 1}}, want: false},
		{name: "up and over", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 2, y: 2, z: 1}}, want: false},
		{name: "corner", args: args{a: point3{x: 1, y: 1, z: 1}, b: point3{x: 2, y: 2, z: 2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjacent2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("adjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
