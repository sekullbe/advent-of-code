package main

import "testing"

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
			name: "Example1", args: args{inputText: "A Y\nB X\nC Z\n"}, want: 15,
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
			name: "Example1", args: args{inputText: "A Y\nB X\nC Z\n"}, want: 12,
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

func Test_parseOneGame(t *testing.T) {
	type args struct {
		game string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Example1", args: args{game: "A Y"}, want: "A", want1: "Y",
		},
		{
			name: "Example2", args: args{game: "B X"}, want: "B", want1: "X",
		},
		{
			name: "Example3", args: args{game: "C Z"}, want: "C", want1: "Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOneGame(tt.args.game)
			if got != tt.want {
				t.Errorf("parseOneGame() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseOneGame() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_scoreGame(t *testing.T) {
	type args struct {
		move    string
		counter string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "win", args: args{"A", "Y"}, want: 6,
		},
		{
			name: "loss", args: args{"B", "X"}, want: 0,
		},
		{
			name: "draw", args: args{"C", "Z"}, want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreGame(tt.args.move, tt.args.counter); got != tt.want {
				t.Errorf("scoreGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreMove(t *testing.T) {
	type args struct {
		move string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A", args: args{"A"}, want: 1,
		},
		{
			name: "B", args: args{"B"}, want: 2,
		},
		{
			name: "C", args: args{"C"}, want: 3,
		},
		{
			name: "X", args: args{"X"}, want: 1,
		},
		{
			name: "Y", args: args{"Y"}, want: 2,
		},
		{
			name: "Z", args: args{"Z"}, want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreMove(tt.args.move); got != tt.want {
				t.Errorf("scoreMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeCounterFor(t *testing.T) {
	type args struct {
		move         string
		desiredState string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Example1", args: args{"A", "X"}, want: "Z"},
		{name: "Example1", args: args{"A", "Y"}, want: "X"},
		{name: "Example1", args: args{"A", "Z"}, want: "Y"},
		{name: "Example1", args: args{"B", "X"}, want: "X"},
		{name: "Example1", args: args{"B", "Y"}, want: "Y"},
		{name: "Example1", args: args{"B", "Z"}, want: "Z"},
		{name: "Example1", args: args{"C", "X"}, want: "Y"},
		{name: "Example1", args: args{"C", "Y"}, want: "Z"},
		{name: "Example1", args: args{"C", "Z"}, want: "X"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeCounterFor(tt.args.move, tt.args.desiredState); got != tt.want {
				t.Errorf("computeCounterFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
