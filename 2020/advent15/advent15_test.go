package main

import "testing"

func Test_playGame(t *testing.T) {
	type args struct {
		starters []int
		lastTurn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1-4", args{[]int{0, 3, 6}, 4}, 0},
		{"example1-5", args{[]int{0, 3, 6}, 5}, 3},
		{"example1-6", args{[]int{0, 3, 6}, 6}, 3},
		{"example1-7", args{[]int{0, 3, 6}, 7}, 1},
		{"example1-8", args{[]int{0, 3, 6}, 8}, 0},
		{"example1-9", args{[]int{0, 3, 6}, 9}, 4},
		{"example1-10", args{[]int{0, 3, 6}, 10}, 0},
		{"example1-11", args{[]int{0, 3, 6}, 11}, 2},
		{"example1-12", args{[]int{0, 3, 6}, 12}, 0},
		{"example1-13", args{[]int{0, 3, 6}, 13}, 2},
		{"example1-2020", args{[]int{0, 3, 6}, 2020}, 436},
		{"example2-2020", args{[]int{1, 3, 2}, 2020}, 1},
		{"example3-2020", args{[]int{2, 1, 3}, 2020}, 10},
		{"example4-2020", args{[]int{1, 2, 3}, 2020}, 27},
		{"example5-2020", args{[]int{2, 3, 1}, 2020}, 78},
		{"example6-2020", args{[]int{3, 2, 1}, 2020}, 438},
		{"example7-2020", args{[]int{3, 1, 2}, 2020}, 1836},
		{"part2-1", args{[]int{0, 3, 6}, 30000000}, 175594},
		{"part2-2", args{[]int{1, 3, 2}, 30000000}, 2578},
		{"part2-3", args{[]int{2, 1, 3}, 30000000}, 3544142},
		{"part2-4", args{[]int{1, 2, 3}, 30000000}, 261214},
		{"part2-5", args{[]int{2, 3, 1}, 30000000}, 6895259},
		{"part2-6", args{[]int{3, 2, 1}, 30000000}, 18},
		{"part2-7", args{[]int{3, 1, 2}, 30000000}, 362},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := playGame(tt.args.starters, tt.args.lastTurn); got != tt.want {
				t.Errorf("playGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
