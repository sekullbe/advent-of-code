package main

import (
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

func Test_seekLeftAndReplace(t *testing.T) {
	type args struct {
		num    string
		newNum int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nothing to the left",
			args: args{num: "[[[[", newNum: 9},
			want: "[[[[",
		},
		{
			name: "basic to the left",
			args: args{num: "[7,[6,[5,[4,", newNum: 3},
			want: "[7,[6,[5,[7,",
		},
		{
			name: "multidigit to the left",
			args: args{num: "[7,[6,[5,[29,", newNum: 3},
			want: "[7,[6,[5,[32,",
		},
		{
			name: "another example",
			args: args{num: "[[3,[2,[8,0]]],[9,[5,[4,", newNum: 3},
			want: "[[3,[2,[8,0]]],[9,[5,[7,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seekLeftAndReplace(tt.args.num, tt.args.newNum); got != tt.want {
				t.Errorf("seekLeftAndReplace() = '%s', want '%s'", got, tt.want)
			}
		})
	}
}

func Test_seekRightAndReplace(t *testing.T) {
	type args struct {
		num    string
		newNum int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nothing to the right",
			args: args{num: "]]]]", newNum: 2},
			want: "]]]]",
		},
		{
			name: "simple to the right",
			args: args{num: "1],2],3],4]", newNum: 8},
			want: "9],2],3],4]",
		},
		{
			name: "distant to the right",
			args: args{num: "]]],1]", newNum: 8},
			want: "]]],9]",
		},
		{
			name: "distant, additive to the right",
			args: args{num: "]]],32]", newNum: 9},
			want: "]]],41]",
		},
		{
			name: "additive to the right",
			args: args{num: "12],2],3],4]", newNum: 8},
			want: "20],2],3],4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seekRightAndReplace(tt.args.num, tt.args.newNum); got != tt.want {
				t.Errorf("seekRightAndReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_explodeAtIndex(t *testing.T) {
	type args struct {
		num     string
		explodo int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{num: "[[[[[9,8],1],2],3],4]", explodo: 4},
			want: "[[[[0,9],2],3],4]",
		},
		{
			name: "example 2",
			args: args{num: "[7,[6,[5,[4,[3,2]]]]]", explodo: 12},
			want: "[7,[6,[5,[7,0]]]]",
		},
		{
			name: "example 3",
			args: args{num: "[[6,[5,[4,[3,2]]]],1]", explodo: 10},
			want: "[[6,[5,[7,0]]],3]",
		},
		{
			name: "example 4",
			args: args{num: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", explodo: 10},
			want: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			name: "example 5",
			args: args{num: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", explodo: 24},
			want: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := explodeAtIndex(tt.args.num, tt.args.explodo); got != tt.want {
				t.Errorf("explodeAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_explode(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name         string
		args         args
		wantNewNum   string
		wantExploded bool
	}{
		{
			name:         "example 1",
			args:         args{num: "[[[[[9,8],1],2],3],4]"},
			wantNewNum:   "[[[[0,9],2],3],4]",
			wantExploded: true,
		},
		{
			name:         "example 2",
			args:         args{num: "[7,[6,[5,[4,[3,2]]]]]"},
			wantNewNum:   "[7,[6,[5,[7,0]]]]",
			wantExploded: true,
		},
		{
			name:         "example 3",
			args:         args{num: "[[6,[5,[4,[3,2]]]],1]"},
			wantNewNum:   "[[6,[5,[7,0]]],3]",
			wantExploded: true,
		},
		{
			name:         "example 4",
			args:         args{num: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"},
			wantNewNum:   "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			wantExploded: true,
		},
		{
			name:         "example 5",
			args:         args{num: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
			wantNewNum:   "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
			wantExploded: true,
		},
		{
			name:         "no boom today",
			args:         args{num: "[[3,[2,[8,0]]],[9,[5,[4,3]]]]"},
			wantNewNum:   "[[3,[2,[8,0]]],[9,[5,[4,3]]]]",
			wantExploded: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewNum, gotExploded := explode(tt.args.num)
			if gotNewNum != tt.wantNewNum {
				t.Errorf("explode() gotNewNum = %v, want %v", gotNewNum, tt.wantNewNum)
			}
			if gotExploded != tt.wantExploded {
				t.Errorf("explode() gotExploded = %v, want %v", gotExploded, tt.wantExploded)
			}
		})
	}
}

func Test_doTheFirstSplit(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name:  "no-op",
			args:  args{num: "[[[[0,7],4],[5,[0,3]]],[1,1]]"},
			want:  "[[[[0,7],4],[5,[0,3]]],[1,1]]",
			want1: false,
		},
		{
			name:  "trivial even",
			args:  args{num: "[14,3]"},
			want:  "[[7,7],3]",
			want1: true,
		},
		{
			name:  "trivial odd",
			args:  args{num: "[15,3]"},
			want:  "[[7,8],3]",
			want1: true,
		},
		{
			name:  "example 1",
			args:  args{num: "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
			want:  "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			want1: true,
		},
		{
			name:  "example 2",
			args:  args{num: "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
			want:  "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := doTheFirstSplit(tt.args.num)
			if got != tt.want {
				t.Errorf("doTheFirstSplit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("doTheFirstSplit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reduce(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{num: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"},
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduce(tt.args.num); got != tt.want {
				t.Errorf("reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "trivial",
			args: args{num1: "[1,2]", num2: "[3,4]"},
			want: "[[1,2],[3,4]]",
		},
		{
			name: "example 1",
			args: args{num1: "[[[[4,3],4],4],[7,[[8,4],9]]]", num2: "[1,1]"},
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatAdd(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "trivial 2 pairs",
			args: args{nums: []string{"[1,2]", "[3,4]"}},
			want: "[[1,2],[3,4]]",
		},
		{
			name: "trivial example 1",
			args: args{nums: []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}},
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			name: "trivial example 2",
			args: args{nums: []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}},
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			name: "trivial example 3",
			args: args{nums: []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}},
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			name: "example 1",
			args: args{nums: parsers.SplitByLines("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]\n[7,[5,[[3,8],[1,4]]]]\n[[2,[2,2]],[8,[8,1]]]\n[2,9]\n[1,[[[9,3],9],[[9,0],[0,7]]]]\n[[[5,[7,4]],7],1]\n[[[[4,2],2],6],[8,7]]\n")},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatAdd(tt.args.nums); got != tt.want {
				t.Errorf("repeatAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_magnitude(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trivial",
			args: args{num: "[9,1]"},
			want: 29,
		},
		{
			name: "trivial recursive",
			args: args{num: "[[9,1],[1,9]]"},
			want: 129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magnitude(tt.args.num); got != tt.want {
				t.Errorf("magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitPair(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "trivial",
			args: args{num: "[1,2]"},
			want: []string{"1", "2"},
		},
		{
			name: "pairs",
			args: args{num: "[[1,2],[3,4]]"},
			want: []string{"[1,2]", "[3,4]"},
		},
		{
			name: "more complex pairs",
			args: args{num: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
			want: []string{"[[[8,7],[7,7]],[[8,6],[7,7]]]", "[[[0,7],[6,6]],[8,7]]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitPair(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitPair() = %v, want %v", got, tt.want)
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
			args: args{"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]\n[[[5,[2,8]],4],[5,[[9,9],0]]]\n[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]\n[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]\n[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]\n[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]\n[[[[5,4],[7,7]],8],[[8,3],8]]\n[[9,3],[[9,9],[6,[4,9]]]]\n[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]\n[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"},
			want: 3993,
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

func Test_getAllNumPairs(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want []numPair
	}{
		{
			name: "simple",
			args: args{nums: []string{"a", "b", "c"}},
			want: []numPair{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"c", "a"}, {"b", "c"}, {"c", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllNumPairs(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllNumPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
