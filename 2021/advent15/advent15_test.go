package main

import (
	"reflect"
	"testing"
)

func Test_point_neighbors(t *testing.T) {
	type fields struct {
		row int
		col int
	}
	type args struct {
		maxRow int
		maxCol int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []point
	}{
		{
			name:   "center",
			fields: fields{row: 2, col: 2},
			args:   args{maxRow: 10, maxCol: 10},
			want:   []point{{1, 2}, {3, 2}, {2, 1}, {2, 3}},
		},
		{
			name:   "row0",
			fields: fields{row: 0, col: 2},
			args:   args{maxRow: 10, maxCol: 10},
			want:   []point{{1, 2}, {0, 1}, {0, 3}},
		},
		{
			name:   "row 10",
			fields: fields{row: 10, col: 2},
			args:   args{maxRow: 10, maxCol: 10},
			want:   []point{{9, 2}, {10, 1}, {10, 3}},
		},
		{
			name:   "col 10",
			fields: fields{row: 2, col: 10},
			args:   args{maxRow: 10, maxCol: 10},
			want:   []point{{1, 10}, {3, 10}, {2, 9}},
		},
		{
			name:   "maxcorner",
			fields: fields{row: 10, col: 10},
			args:   args{maxRow: 10, maxCol: 10},
			want:   []point{{9, 10}, {10, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				row: tt.fields.row,
				col: tt.fields.col,
			}
			if got := p.neighbors(tt.args.maxRow, tt.args.maxCol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point_uniquePointId(t *testing.T) {
	type fields struct {
		row int
		col int
	}
	type args struct {
		rows int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1,1 =10*1 + 2",
			fields: fields{row: 1, col: 2},
			args:   args{rows: 10},
			want:   12,
		},
		{
			name:   "40*30 =50*40 +30",
			fields: fields{row: 40, col: 30},
			args:   args{rows: 50},
			want:   2030,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				row: tt.fields.row,
				col: tt.fields.col,
			}
			if got := p.uniquePointId(tt.args.rows); got != tt.want {
				t.Errorf("uniquePointId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
		graphSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trivial",
			args: args{inputText: "122\n992\n992\n", graphSize: 3},
			want: 8,
		},
		{
			name: "example",
			args: args{inputText: "1163751742\n1381373672\n2136511328\n3694931569\n7463417111\n1319128137\n1359912421\n3125421639\n1293138521\n2311944581", graphSize: 10},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText, tt.args.graphSize); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copyLineForTiles(t *testing.T) {
	type args struct {
		line  string
		tiles int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no-op",
			args: args{line: "987654321", tiles: 1},
			want: "987654321",
		},
		{
			name: "two",
			args: args{line: "987654321", tiles: 2},
			want: "987654321198765432",
		},
		{
			name: "three",
			args: args{line: "999888777", tiles: 3},
			want: "999888777111999888222111999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyLineForTiles(tt.args.line, tt.args.tiles); got != tt.want {
				t.Errorf("copyLineForTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point_uniquePointId1(t *testing.T) {
	type fields struct {
		row int
		col int
	}
	type args struct {
		rows int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "origin",
			fields: fields{row: 0, col: 0},
			args:   args{rows: 10},
			want:   0,
		},
		{
			name:   "0,9",
			fields: fields{row: 0, col: 9},
			args:   args{rows: 10},
			want:   9,
		},
		{
			name:   "1,0",
			fields: fields{row: 1, col: 0},
			args:   args{rows: 10},
			want:   10,
		},
		{
			name:   "reallybig 145,475",
			fields: fields{row: 145, col: 475},
			args:   args{rows: 500},
			want:   72975,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				row: tt.fields.row,
				col: tt.fields.col,
			}
			if got := p.uniquePointId(tt.args.rows); got != tt.want {
				t.Errorf("uniquePointId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
		graphSize int
		tiles     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "super trivial",
			args: args{inputText: "1", graphSize: 1, tiles: 2},
			want: 5,
		},
		{
			name: "trivial",
			args: args{inputText: "12\n92\n", graphSize: 2, tiles: 2},
			want: 14,
		},
		{
			name: "example",
			args: args{inputText: "1163751742\n1381373672\n2136511328\n3694931569\n7463417111\n1319128137\n1359912421\n3125421639\n1293138521\n2311944581", graphSize: 10, tiles: 5},
			want: 315,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText, tt.args.graphSize, tt.args.tiles); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeRollingRisk(t *testing.T) {
	type args struct {
		risk int
		roll int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1,0", args: args{1, 0}, want: 1},
		{name: "1,9", args: args{1, 9}, want: 1},
		{name: "1,8", args: args{1, 8}, want: 9},
		{name: "1,18", args: args{1, 18}, want: 1},
		{name: "1,10", args: args{1, 10}, want: 2},
		{name: "8,8", args: args{8, 8}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeRollingRisk(tt.args.risk, tt.args.roll); got != tt.want {
				t.Errorf("computeRollingRisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
