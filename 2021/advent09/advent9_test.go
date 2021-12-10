package main

import (
	"reflect"
	"testing"
)

func Test_parseChart(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want caveChart
	}{
		{
			name: "simple",
			args: args{lines: "123\n456\n789\n"},
			want: caveChart{
				{{0, 0, 1, 0}, {0, 1, 2, 0}, {0, 2, 3, 0}},
				{{1, 0, 4, 0}, {1, 1, 5, 0}, {1, 2, 6, 0}},
				{{2, 0, 7, 0}, {2, 1, 8, 0}, {2, 2, 9, -1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseChart(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseChart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neighborDepths(t *testing.T) {
	type args struct {
		chart caveChart
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple1",
			args: args{
				chart: caveChart{
					{{0, 0, 1, 0}, {0, 1, 2, 0}, {0, 2, 3, 0}},
					{{1, 0, 4, 0}, {1, 1, 5, 0}, {1, 2, 6, 0}},
					{{2, 0, 7, 0}, {2, 1, 8, 0}, {2, 2, 9, 0}},
				}, row: 1, col: 1},
			want: []int{2, 8, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := neighborHeights(tt.args.chart, tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neighborHeights() = %v, want %v", got, tt.want)
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
		{
			name: "example",
			args: args{inputText: "2199943210\n3987894921\n9856789892\n8767896789\n9899965678"},
			want: 1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
