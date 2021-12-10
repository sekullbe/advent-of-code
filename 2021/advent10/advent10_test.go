package main

import (
	"reflect"
	"testing"
)

func Test_scoreIllegalCharacter(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreIllegalCharacter(tt.args.c); got != tt.want {
				t.Errorf("scoreIllegalCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_syntaxCheck(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name           string
		args           args
		wantCorrupt    bool
		wantIncomplete bool
		wantBadchar    rune
		wantCompletion []rune
	}{
		{
			name:           "trivial pass",
			args:           args{line: "([])"},
			wantCorrupt:    false,
			wantIncomplete: false,
			wantBadchar:    '0',
			wantCompletion: []rune{},
		},
		{
			name:           "trivial corrupt",
			args:           args{line: "([}])"},
			wantCorrupt:    true,
			wantIncomplete: false,
			wantBadchar:    '}',
			wantCompletion: []rune{},
		},
		{
			name:           "trivial incomplete",
			args:           args{line: "([]){["},
			wantCorrupt:    false,
			wantIncomplete: true,
			wantBadchar:    '0',
			wantCompletion: []rune{']', '}'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCorrupt, gotIncomplete, gotBadchar, gotCompletion := syntaxCheck(tt.args.line)
			if gotCorrupt != tt.wantCorrupt {
				t.Errorf("syntaxCheck() gotCorrupt = %v, want %v", gotCorrupt, tt.wantCorrupt)
			}
			if gotIncomplete != tt.wantIncomplete {
				t.Errorf("syntaxCheck() gotIncomplete = %v, want %v", gotIncomplete, tt.wantIncomplete)
			}
			if gotBadchar != tt.wantBadchar {
				t.Errorf("syntaxCheck() gotBadchar = %v, want %v", gotBadchar, tt.wantBadchar)
			}
			if !reflect.DeepEqual(gotCompletion, tt.wantCompletion) {
				t.Errorf("syntaxCheck() gotCompletion = %v, want %v", gotCompletion, tt.wantCompletion)
			}
		})
	}
}

func Test_scoreCompletionCharacters(t *testing.T) {
	type args struct {
		chs []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{chs: []rune{}},
			want: 0,
		},
		{
			name: "example",
			args: args{chs: []rune{']', ')', '}', '>'}},
			want: 294,
		},
		{
			name: "example ]]}}]}]}>",
			args: args{chs: []rune{']', ']', '}', '}', ']', '}', ']', '}', '>'}},
			want: 995444,
		},
		{
			name: "example )}>]})",
			args: args{chs: []rune{')', '}', '>', ']', '}', ')'}},
			want: 5566,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreCompletionCharacters(tt.args.chs); got != tt.want {
				t.Errorf("scoreCompletionCharacters() = %v, want %v", got, tt.want)
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
			name: "example",
			args: args{inputText: "[({(<(())[]>[[{[]{<()<>>\n[(()[<>])]({[<{<<[]>>(\n{([(<{}[<>[]}>{[]{[(<()>\n(((({<>}<{<{<>}{[]{[]{}\n[[<[([]))<([[{}[[()]]]\n[{[{({}]{}}([{[{{{}}([]\n{<[[]]>}<{[{[{[]{()[[[]\n[<(<(<(<{}))><([]([]()\n<{([([[(<>()){}]>(<<{{\n<{([{{}}[<[[[<>{}]]]>[]]"},
			want: 288957,
		},
		{
			name: "real",
			args: args{inputText: inputText},
			want: 4329504793,
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
