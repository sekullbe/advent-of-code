package advent6

import "testing"

func Test_countMatchesInString(t *testing.T) {
	type args struct {
		respons string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{respons: "0"},
			want: 0,
		},
		{
			name: "trivial",
			args: args{respons: "a"},
			want: 1,
		},
		{
			name: "oneline",
			args: args{respons: "abc"},
			want: 3,
		},
		{
			name: "twolines_all",
			args: args{respons: "abc\nabc"},
			want: 3,
		},
		{
			name: "twolines_none",
			args: args{respons: "abc\ndef"},
			want: 0,
		},
		{
			name: "twolines_some",
			args: args{respons: "abc\nade"},
			want: 1,
		},
		{
			name: "threelines_none",
			args: args{respons: "a\nb\nc"},
			want: 0,
		},
		{
			name: "threelines_example",
			args: args{respons: "ybcgtxznorvjel\nbrlyvoexnjtgcz\nlnbgtxvoiyecjrz"},
			want: 14,
		},
		{
			name: "threelines_example2",
			args: args{respons: "mw\nwm\nmw"},
			want: 2,
		},
		{
			name: "threelines_example3",
			args: args{respons: "ks\nskh\nsk"},
			want: 2,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatchesInString(tt.args.respons); got != tt.want {
				t.Errorf("countMatchesInString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2_doit(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{inp:"abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2_doit(tt.args.inp); got != tt.want {
				t.Errorf("run2_doit() = %v, want %v", got, tt.want)
			}
		})
	}
}
