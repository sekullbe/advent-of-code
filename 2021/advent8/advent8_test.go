package main

import (
	"reflect"
	"testing"
)

func Test_newDigitPattern(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want digitPattern
	}{
		{
			name: "1",
			args: args{p: "cf"},
			want: digitPattern{pattern: "cf", possibleNums: []int{1}, num: 1},
		},
		{
			name: "235",
			args: args{p: "acedg"},
			want: digitPattern{pattern: "acdeg", possibleNums: []int{2, 3, 5}, num: -1},
		},
		{
			name: "4",
			args: args{p: "bcdf"},
			want: digitPattern{pattern: "bcdf", possibleNums: []int{4}, num: 4},
		},
		{
			name: "7",
			args: args{p: "acf"},
			want: digitPattern{pattern: "acf", possibleNums: []int{7}, num: 7},
		},
		{
			name: "690",
			args: args{p: "abcefg"},
			want: digitPattern{pattern: "abcefg", possibleNums: []int{6, 9, 0}, num: -1},
		},
		{
			name: "8",
			args: args{p: "abcdefg"},
			want: digitPattern{pattern: "abcdefg", possibleNums: []int{8}, num: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDigitPattern(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDigitPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name        string
		args        args
		wantInputs  []digitPattern
		wantOutputs []digitPattern
	}{
		{
			name:        "example",
			args:        args{line: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"},
			wantInputs:  []digitPattern{{"be", []int{1}, 1}, {"abcdefg", []int{8}, 8}, {"bcdefg", []int{6, 9, 0}, -1}, {"acdefg", []int{6, 9, 0}, -1}, {"bceg", []int{4}, 4}, {"cdefg", []int{2, 3, 5}, -1}, {"abdefg", []int{6, 9, 0}, -1}, {"bcdef", []int{2, 3, 5}, -1}, {"abcdf", []int{2, 3, 5}, -1}, {"bde", []int{7}, 7}},
			wantOutputs: []digitPattern{{"abcdefg", []int{8}, 8}, {"bcdef", []int{2, 3, 5}, -1}, {"bcdefg", []int{6, 9, 0}, -1}, {"bceg", []int{4}, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInputs, gotOutputs := parseLine(tt.args.line)
			if !reflect.DeepEqual(gotInputs, tt.wantInputs) {
				t.Errorf("parseLine() gotInputs = %v, want %v", gotInputs, tt.wantInputs)
			}
			if !reflect.DeepEqual(gotOutputs, tt.wantOutputs) {
				t.Errorf("parseLine() gotOutputs = %v, want %v", gotOutputs, tt.wantOutputs)
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
		want int
	}{
		{
			name: "example",
			args: args{inputText: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\nedbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc\nfgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg\nfbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb\naecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea\nfgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb\ndbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe\nbdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef\negadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb\ngcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"},
			want: 26,
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
			name: "example",
			args: args{inputText: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\nedbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc\nfgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg\nfbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb\naecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea\nfgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb\ndbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe\nbdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef\negadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb\ngcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"},
			want: 61229,
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

func Test_deduceCodesFromInputLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "example",
			args: args{line: "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"},
			want: map[string]int{"ab": 1, "abcdef": 9, "abcdefg": 8, "abd": 7, "abef": 4, "acdfg": 2, "bcdef": 5, "abcdf": 3, "bcdefg": 6, "abcdeg": 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deduceCodesFromInputLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deduceCodesFromInputLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
