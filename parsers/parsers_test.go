package parsers

import (
	"reflect"
	"testing"
)

func TestSplitByEmptyNewline(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitByEmptyNewline(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitByEmptyNewline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitByLines(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitByLines(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitByLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringsToIntSlice(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringsToIntSlice(tt.args.inputText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringsToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringsWithCommasToIntSlice(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "trailing newline",
			args: args{inputText: "1,1,1,1,3\n"},
			want: []int{1, 1, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringsWithCommasToIntSlice(tt.args.inputText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringsWithCommasToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
