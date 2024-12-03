package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/tools"
	"regexp"
	"strings"
)

//go:embed input.txt
var inputText string

// var sampleText string ="xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
	fmt.Printf("Magic number: %d\n", run2b(inputText)) // 44339270 too low
}

func run1(input string) int {

	pairs := extractMuls(input)
	sum := 0
	for _, p := range pairs {
		sum += p.a * p.b
	}

	return sum
}

func run2(input string) int {
	sum := 0
	dos := strings.SplitAfter(input, "do()")
	for _, do := range dos {
		enabled := strings.Split(do, "don't")
		sum += run1(enabled[0])
	}
	return sum
}

// this didn't work at first but after trying it the other way I fixed it; the problem was that the
// extraction regexp didn't work across newlines; once I removed them it worked fine
func run2b(input string) int {

	// now we have to handle do() and don't()
	enabledSegments := extractEnabled(input)
	sum := 0
	for _, s := range enabledSegments {
		sum += run1(s)
	}

	return sum
}

type pair struct{ a, b int }

func extractMuls(input string) []pair {
	var muls = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`).FindAllStringSubmatch(input, -1)
	// mul is ["mul(1,2)", "1","2"
	pairs := make([]pair, 0, len(muls))
	for _, mul := range muls {
		pairs = append(pairs, pair{tools.Atoi(mul[1]), tools.Atoi(mul[2])})
	}
	return pairs
}

func extractEnabled(input string) []string {
	input = strings.ReplaceAll(input, "\n", "")
	enabledSegments := regexp.MustCompile(`do\(\)(.*?)don't\(\)`).FindAllStringSubmatch("do()"+input+"don't()", -1)
	instructions := []string{}
	for _, segment := range enabledSegments {
		instructions = append(instructions, strings.TrimSpace(segment[1]))
	}
	return instructions
}
