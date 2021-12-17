package main

import (
	"reflect"
	"testing"
)

func Test_parseTarget(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want target
	}{
		{
			name: "example",
			args: args{inputText: "target area: x=20..30, y=-10..-5"},
			want: target{20, 30, -10, -5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTarget(tt.args.inputText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_launch(t *testing.T) {
	type args struct {
		s shot
		t target
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 7,2",
			args: args{s: shot{0, 0, 7, 2}, t: target{20, 30, -10, -5}},
			want: true,
		},
		{
			name: "example 6,3",
			args: args{s: shot{0, 0, 6, 3}, t: target{20, 30, -10, -5}},
			want: true,
		},
		{
			name: "example 9,0",
			args: args{s: shot{0, 0, 9, 0}, t: parseTarget("target area: x=20..30, y=-10..-5")},
			want: true,
		},
		{
			name: "example 17,-4",
			args: args{s: shot{0, 0, 17, -4}, t: parseTarget("target area: x=20..30, y=-10..-5")},
			want: false,
		},
		{
			name: "cheaty real target",
			args: args{s: shot{0, 0, 50, -200}, t: parseTarget(inputText)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := launch(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("launch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_target_findPossibleDxRange(t1 *testing.T) {
	type fields struct {
		minX int
		maxX int
		minY int
		maxY int
	}
	type args struct {
		initX int
	}
	tests := []struct {
		name           string
		fields         fields
		wantVelocities []int
	}{
		{
			name:           "10-30",
			fields:         fields{10, 12, 0, 0},
			wantVelocities: []int{4, 5, 6, 10, 11, 12},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := target{
				minX: tt.fields.minX,
				maxX: tt.fields.maxX,
				minY: tt.fields.minY,
				maxY: tt.fields.maxY,
			}
			if gotVelocities := t.findPossibleDxRange(); !reflect.DeepEqual(gotVelocities, tt.wantVelocities) {
				t1.Errorf("findPossibleDxRange() = %v, want %v", gotVelocities, tt.wantVelocities)
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
			name: "example1",
			args: args{inputText: "target area: x=20..30, y=-10..-5"},
			want: 45,
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

func Test_target_findPossibleDyRange(t1 *testing.T) {
	type fields struct {
		minX int
		maxX int
		minY int
		maxY int
	}
	tests := []struct {
		name           string
		fields         fields
		wantVelocities []int
	}{
		{
			name:           "basic",
			fields:         fields{0, 0, -10, -5},
			wantVelocities: []int{-11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := target{
				minX: tt.fields.minX,
				maxX: tt.fields.maxX,
				minY: tt.fields.minY,
				maxY: tt.fields.maxY,
			}
			if gotVelocities := t.findPossibleDyRange(); !reflect.DeepEqual(gotVelocities, tt.wantVelocities) {
				t1.Errorf("findPossibleDyRange() = %v, want %v", gotVelocities, tt.wantVelocities)
			}
		})
	}
}
