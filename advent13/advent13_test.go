package advent13

import (
	"reflect"
	"testing"
)

func Test_calculateWaitFor(t *testing.T) {
	type args struct {
		timestamp int
		busId     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trivial",
			args: args{ timestamp: 2, busId: 1},
			want: 0,
		},
		{
			name: "larger",
			args: args{ timestamp: 5, busId: 4},
			want: 3,
		},
		{
			name: "example 1",
			args: args{ timestamp: 939, busId: 59},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWaitFor(tt.args.timestamp, tt.args.busId); got != tt.want {
				t.Errorf("calculateWaitFor() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_parseInput(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "example",
			args: args {inputText: "939\n7,13,x,x,59,x,31,19\n"},
			want:939,
			want1: []int{7,13,59,31,19},
		},

	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInput(tt.args.inputText)
			if got != tt.want {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
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
			name: "example1",
			args: args {inputText: "939\n7,13,x,x,59,x,31,19\n"},
			want: 1068781,
		},
		{
			name: "example2",
			args: args {inputText: "939\n17,x,13,19\n"},
			want: 3417,
		},
		{
			name: "example3",
			args: args {inputText: "939\n67,7,59,61\n"},
			want: 754018,
		},
		{
			name: "example4",
			args: args {inputText: "939\n67,x,7,59,61\n"},
			want: 779210,
		},
		{
			name: "example5",
			args: args {inputText: "939\n67,7,x,59,61\n"},
			want: 1261476,
		},
		{
			name: "example6",
			args: args {inputText: "939\n1789,37,47,1889\n"},
			want: 1202161486,
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

func Test_run2Faster(t *testing.T) {
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
			args: args {inputText: "939\n7,13,x,x,59,x,31,19\n"},
			want: 1068781,
		},
		{
			name: "example2",
			args: args {inputText: "939\n17,x,13,19\n"},
			want: 3417,
		},
		{
			name: "example3",
			args: args {inputText: "939\n67,7,59,61\n"},
			want: 754018,
		},
		{
			name: "example4",
			args: args {inputText: "939\n67,x,7,59,61\n"},
			want: 779210,
		},
		{
			name: "example5",
			args: args {inputText: "939\n67,7,x,59,61\n"},
			want: 1261476,
		},
		{
			name: "example6",
			args: args {inputText: "939\n1789,37,47,1889\n"},
			want: 1202161486,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2Faster(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
