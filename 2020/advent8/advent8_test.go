package advent8

import "testing"

func TestRun2Doit(t *testing.T) {
	type args struct {
		inputs string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "example",
			args: args{inputs:"nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6\n"},
			want: 8,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Run2Doit(tt.args.inputs)
			if got != tt.want {
				t.Errorf("Run2Doit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Run2Doit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
