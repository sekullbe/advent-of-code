package advent9

import (
	"testing"
)

func Test_isSumOfPairInSlice(t *testing.T) {
	type args struct {
		preamble []int
		target   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args {target: 40, preamble: []int{35,20,15,25,47}},
			want: true,
		},
		{
			name: "error1",
			args: args{target: 403, preamble:[]int{133,281,137,134,154,148,150,161,179,211,181,182,215,203,217,232,233,237,234,241,267,254,256,257,270}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSumOfPairInSlice(tt.args.preamble, tt.args.target); got != tt.want {
				t.Errorf("isSumOfPairInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findContiguousSum(t *testing.T) {
	type args struct {
		data   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "error1",
			args: args{target: 854, data:[]int{133,281,137,134,154,148,150,161,179,211,181,182,215,203,217,232,233,237,234,241,267,254,256,257,270}},
			want: 281,
			want1: 148,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findContiguousSum(tt.args.data, tt.args.target)
			if got != tt.want {
				t.Errorf("findContiguousSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findContiguousSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
