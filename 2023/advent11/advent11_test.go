package main

import "testing"

const sampleGrid = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func Test_run(t *testing.T) {
	type args struct {
		input             string
		expansionConstant int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 galaxy",
			args: args{
				input:             sampleGrid,
				expansionConstant: 2,
			},
			want: 374,
		},
		{
			name: "10 galaxy",
			args: args{
				input:             sampleGrid,
				expansionConstant: 10,
			},
			want: 1030,
		},
		{
			name: "100 galaxy",
			args: args{
				input:             sampleGrid,
				expansionConstant: 100,
			},
			want: 8410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.input, tt.args.expansionConstant); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
