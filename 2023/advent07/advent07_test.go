package main

import (
	"github.com/sekullbe/advent/tools"
	"strconv"
	"testing"
)

const sampleText = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_handType(t *testing.T) {
	type args struct {
		h hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "5", args: args{h: hand{cardstr: "AAAAA"}}, want: FIVE},
		{name: "4a", args: args{h: hand{cardstr: "TAAAA"}}, want: FOUR},
		{name: "4b", args: args{h: hand{cardstr: "AAAA3"}}, want: FOUR},
		{name: "4c", args: args{h: hand{cardstr: "AA2AA"}}, want: FOUR},
		{name: "full a", args: args{h: hand{cardstr: "AAA22"}}, want: FULL},
		{name: "full b", args: args{h: hand{cardstr: "AA22A"}}, want: FULL},
		{name: "full c", args: args{h: hand{cardstr: "22AAA"}}, want: FULL},
		{name: "full d", args: args{h: hand{cardstr: "2A2AA"}}, want: FULL},
		{name: "full e", args: args{h: hand{cardstr: "2A2A2"}}, want: FULL},
		{name: "three a", args: args{h: hand{cardstr: "2AK22"}}, want: THREE},
		{name: "three b", args: args{h: hand{cardstr: "2A2K2"}}, want: THREE},
		{name: "three d", args: args{h: hand{cardstr: "2A2K2"}}, want: THREE},
		{name: "two a", args: args{h: hand{cardstr: "2AAQ2"}}, want: TWO_PAIR},
		{name: "two b", args: args{h: hand{cardstr: "2AAQ2"}}, want: TWO_PAIR},
		{name: "two c", args: args{h: hand{cardstr: "22AA5"}}, want: TWO_PAIR},
		{name: "two d", args: args{h: hand{cardstr: "A52A2"}}, want: TWO_PAIR},
		{name: "two e", args: args{h: hand{cardstr: "2AA28"}}, want: TWO_PAIR},
		{name: "one a", args: args{h: hand{cardstr: "2A428"}}, want: ONE_PAIR},
		{name: "one b", args: args{h: hand{cardstr: "AA428"}}, want: ONE_PAIR},
		{name: "one c", args: args{h: hand{cardstr: "2AA68"}}, want: ONE_PAIR},
		{name: "high", args: args{h: hand{cardstr: "23456"}}, want: HIGH},
		{name: "high", args: args{h: hand{cardstr: "AKQJT"}}, want: HIGH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handType(tt.args.h); got != tt.want {
				t.Errorf("handType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreHand(t *testing.T) {
	type args struct {
		h hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "five", args: args{h: hand{cardstr: "AAAAA"}}, want: 7270126},
		{name: "five hex", args: args{h: hand{cardstr: "AAAAA"}}, want: int(tools.Must(strconv.ParseInt("6EEEEE", 16, 0)))},
		{name: "four max", args: args{h: hand{cardstr: "AAAAK"}}, want: 6221549},
		{name: "four max -1", args: args{h: hand{cardstr: "AAAAQ"}}, want: 6221548},
		{name: "four min", args: args{h: hand{cardstr: "22223"}}, want: 5382691},
		{name: "four min +1", args: args{h: hand{cardstr: "22224"}}, want: 5382692},
		{name: "three a ", args: args{h: hand{cardstr: "22234"}}, want: 3285556},
		{name: "three a hex", args: args{h: hand{cardstr: "22234"}}, want: int(tools.Must(strconv.ParseInt("322234", 16, 0)))},
		{name: "three a hex", args: args{h: hand{cardstr: "22234"}}, want: int(tools.Must(strconv.ParseInt("322234", 16, 0)))},
		{name: "three b", args: args{h: hand{cardstr: "22235"}}, want: 3285557},
		{name: "three b hex", args: args{h: hand{cardstr: "22235"}}, want: int(tools.Must(strconv.ParseInt("322235", 16, 0)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreHand(tt.args.h, handType); got != tt.want {
				t.Errorf("scoreHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sampleText}, want: 6440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handTypeWithJoker(t *testing.T) {
	type args struct {
		h hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "5", args: args{h: hand{cardstr: "AAAAA"}}, want: FIVE},
		{name: "4a", args: args{h: hand{cardstr: "TAAAA"}}, want: FOUR},
		{name: "4b", args: args{h: hand{cardstr: "AAAA3"}}, want: FOUR},
		{name: "4c", args: args{h: hand{cardstr: "AA2AA"}}, want: FOUR},
		{name: "full a", args: args{h: hand{cardstr: "AAA22"}}, want: FULL},
		{name: "full b", args: args{h: hand{cardstr: "AA22A"}}, want: FULL},
		{name: "full c", args: args{h: hand{cardstr: "22AAA"}}, want: FULL},
		{name: "full d", args: args{h: hand{cardstr: "2A2AA"}}, want: FULL},
		{name: "full e", args: args{h: hand{cardstr: "2A2A2"}}, want: FULL},
		{name: "three a", args: args{h: hand{cardstr: "2AK22"}}, want: THREE},
		{name: "three b", args: args{h: hand{cardstr: "2A2K2"}}, want: THREE},
		{name: "three d", args: args{h: hand{cardstr: "2A2K2"}}, want: THREE},
		{name: "two a", args: args{h: hand{cardstr: "2AAQ2"}}, want: TWO_PAIR},
		{name: "two b", args: args{h: hand{cardstr: "2AAQ2"}}, want: TWO_PAIR},
		{name: "two c", args: args{h: hand{cardstr: "22AA5"}}, want: TWO_PAIR},
		{name: "two d", args: args{h: hand{cardstr: "A52A2"}}, want: TWO_PAIR},
		{name: "two e", args: args{h: hand{cardstr: "2AA28"}}, want: TWO_PAIR},
		{name: "one a", args: args{h: hand{cardstr: "2A428"}}, want: ONE_PAIR},
		{name: "one b", args: args{h: hand{cardstr: "AA428"}}, want: ONE_PAIR},
		{name: "one c", args: args{h: hand{cardstr: "2AA68"}}, want: ONE_PAIR},
		{name: "high", args: args{h: hand{cardstr: "23456"}}, want: HIGH},
		{name: "high", args: args{h: hand{cardstr: "AKQ9T"}}, want: HIGH},
		{name: "5a", args: args{h: hand{cardstr: "AAAAJ"}}, want: FIVE},
		{name: "5b", args: args{h: hand{cardstr: "JJJJJ"}}, want: FIVE},
		{name: "5c", args: args{h: hand{cardstr: "JJJJA"}}, want: FIVE},
		{name: "4a", args: args{h: hand{cardstr: "TAJAA"}}, want: FOUR},
		{name: "4b", args: args{h: hand{cardstr: "JAAA3"}}, want: FOUR},
		{name: "4b", args: args{h: hand{cardstr: "JJJ23"}}, want: FOUR},
		{name: "4c", args: args{h: hand{cardstr: "JJJ25"}}, want: FOUR},
		{name: "fulla", args: args{h: hand{cardstr: "2233J"}}, want: FULL},
		{name: "fullb", args: args{h: hand{cardstr: "2J233"}}, want: FULL},
		{name: "threea", args: args{h: hand{cardstr: "22J34"}}, want: THREE},
		{name: "threeb", args: args{h: hand{cardstr: "2245J"}}, want: THREE},
		// no TWO- you can't get two pair with a joker because it'll preferably make a 3

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handTypeWithJoker(tt.args.h); got != tt.want {
				t.Errorf("handTypeWithJoker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sampleText}, want: 5905},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
