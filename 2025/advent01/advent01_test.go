package main

import "testing"

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		start int
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				start: 50,
				input: "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82\n",
			},
			want: 6,
		},
		{
			name: "testR100",
			args: args{
				start: 50,
				input: "R100",
			},
			want: 1,
		},
		{
			name: "testL100",
			args: args{
				start: 50,
				input: "L100",
			},
			want: 1,
		},
		{
			name: "testL100 from zero",
			args: args{
				start: 0,
				input: "L100",
			},
			want: 1,
		},
		{
			name: "testR1000",
			args: args{
				start: 50,
				input: "R1000",
			},
			want: 10,
		},
		{
			name: "testR100fromzero",
			args: args{
				start: 0,
				input: "R100",
			},
			want: 1, // 10 rotations, final zero
		},
		{
			name: "testR1000fromzero",
			args: args{
				start: 0,
				input: "R1000",
			},
			want: 10, // 10 rotations, final zero
		},
		{
			name: "testL1000",
			args: args{
				start: 50,
				input: "L1000",
			},
			want: 10,
		},
		{
			name: "testL1000fromzero",
			args: args{
				start: 0,
				input: "L1000",
			},
			want: 10, // 10 rotations, final zero
		},
		{
			name: "testL110 from 1",
			args: args{
				start: 1,
				input: "L110",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.start, tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSpins(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{d: 0}, want: 0},
		{name: "1", args: args{d: 1}, want: 0},
		{name: "99", args: args{d: 99}, want: 0},
		{name: "100", args: args{d: 100}, want: 1},
		{name: "101", args: args{d: 101}, want: 1},
		{name: "199", args: args{d: 199}, want: 1},
		{name: "-100", args: args{d: -100}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSpins(tt.args.d); got != tt.want {
				t.Errorf("countSpins() = %v, want %v", got, tt.want)
			}
		})
	}
}
