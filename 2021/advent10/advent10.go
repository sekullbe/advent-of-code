package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
	"log"
	"sort"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	score := 0
	for _, line := range parsers.SplitByLines(inputText) {
		corrupt, _, badRune, _ := syntaxCheck(line)
		if corrupt {
			score += scoreIllegalCharacter(badRune)
		}
	}

	return score
}

func run2(inputText string) int {

	scores := []int{}
	for _, line := range parsers.SplitByLines(inputText) {
		corrupt, incomplete, _, completion := syntaxCheck(line)
		if corrupt {
			continue
		} else {
			if incomplete {
				scores = append(scores, scoreCompletionCharacters(completion))
			}
		}
	}

	sort.Ints(scores)
	score := scores[len(scores)/2]

	return score
}

func syntaxCheck(line string) (corrupt bool, incomplete bool, badchar rune, completion []rune) {
	stack := stack.New()

	for _, ch := range line {
		switch ch {
		case '(', '[', '{', '<':
			stack.Push(ch)
		case ')', ']', '}', '>':
			t := stack.Pop()
			if !(t == inverseOf(ch)) {
				return true, false, ch, []rune{}
			}
		default:
			panic(fmt.Sprintf("I don't know what to do with %c", ch))
		}
	}

	if stack.Empty() {
		return false, false, '0', []rune{}
	}
	// also return the completion characters
	for !stack.Empty() {
		completion = append(completion, inverseOf(stack.Pop().(rune)))
	}
	return false, true, '0', completion
}

func inverseOf(c rune) rune {
	switch c {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	default:
		return ' '
	}
}

func scoreIllegalCharacter(c rune) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		log.Printf("super-illegal character '%c'\n", c)
		return 0
	}
}
func scoreCompletionCharacters(chs []rune) int {
	score := 0
	for _, ch := range chs {
		score *= 5
		switch ch {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}
	return score
}
