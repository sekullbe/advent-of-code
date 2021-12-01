package advent12

import (
	"reflect"
	"testing"
)

func Test_parseCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name      string
		args      args
		wantInstr string
		wantArg   int
	}{
		{
			name: "N30",
			args: args{command: "N30"},
			wantInstr: "N",
			wantArg: 30,
		},
		{
			name: "S1",
			args: args{command: "S1"},
			wantInstr: "S",
			wantArg: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInstr, gotArg := parseCommand(tt.args.command)
			if gotInstr != tt.wantInstr {
				t.Errorf("parseCommand() gotInstr = %v, want %v", gotInstr, tt.wantInstr)
			}
			if gotArg != tt.wantArg {
				t.Errorf("parseCommand() gotArg = %v, want %v", gotArg, tt.wantArg)
			}
		})
	}
}


func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want ship
	}{
		{
			name: "example",
			args: args{inputText: "F10\nN3\nF7\nR90\nF11"},
			want: ship{x:17, y:-8, facing: S},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); !reflect.DeepEqual(got, tt.want) {
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
		want ship
	}{
		{
			name: "example",
			args: args{inputText: "F10\nN3\nF7\nR90\nF11"},
			want: ship{x: 214, y:-72, facing: E},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
