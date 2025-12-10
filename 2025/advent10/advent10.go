package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	combos "github.com/mxschmitt/golang-combinations"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

type machine struct {
	numLights  int
	finalState int64
	buttons    [][]int
	joltages   []int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	machines := []machine{}
	for _, line := range parsers.SplitByLines(input) {
		numLights, fs, b, j := parseMachine(line)
		machines = append(machines, machine{numLights, fs, b, j})
	}

	sum := 0
	for _, m := range machines {
		sum += solveOneMachine(m)
	}

	return sum
}

func solveOneMachine(m machine) int {
	// we know pushing a button more than once is pointless
	// so the solution set is all orderings of all the buttons
	// just try 'em all. and it works in 80ms.
	trials := combos.Combinations(m.buttons, -1)
	minPresses := math.MaxInt
	for _, trial := range trials {
		var state int64 = 0
		for n, p := range trial {
			state = pushAButtonBitmap(state, m.numLights, p)
			if state == m.finalState {
				minPresses = tools.MinInt(minPresses, n+1)
			}
		}

	}

	return minPresses
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	/*
		brute force method- this is effectively BFS:
		  check state to see if we have been here before and what buttons have been pressed here
		  if there are any unpressed, check state to see if it's been done before and press one that hasn't been tried
		  if we have what we want, store the win and start over
		  if we're over, store the failure and start over
		  store the try

		  once done, find the shortest win and return that value

	*/
	machines := []machine{}
	for _, line := range parsers.SplitByLines(input) {
		numLights, fs, b, j := parseMachine(line)
		machines = append(machines, machine{numLights, fs, b, j})
	}

	sum := 0
	for _, m := range machines {
		sum += solveOneMachine(m)
	}

	return 0
}

func parseMachine(s string) (int, int64, [][]int, []int) {
	re := regexp.MustCompile(`\[([#.]+)] (.*?) {([0-9,]+)}`)
	matches := re.FindStringSubmatch(s)
	numLights := len(matches[1])
	fs := parseFinalState(matches[1])
	buttons := parseButtons(matches[2])
	joltages := parsers.StringsWithCommasToIntSlice(matches[3])
	return numLights, fs, buttons, joltages
}

func parseFinalState(state string) int64 {
	s := strings.Trim(state, "[]")
	s = strings.ReplaceAll(s, ".", "0")
	s = strings.ReplaceAll(s, "#", "1")
	finalState, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		panic(err)
	}
	return finalState
}

func parseButtons(bstr string) [][]int {
	buttons := [][]int{}
	// bstr is someting like "(3) (1,3) (2) (2,3) (0,2) (0,1)"
	// for each (x) turn it into a list of ints
	for _, s := range strings.Fields(bstr) {
		b := parsers.StringsWithCommasToIntSlice(strings.Trim(s, "()"))
		buttons = append(buttons, b)
	}
	return buttons
}

func pushAButtonBitmap(state int64, numLights int, button []int) int64 {
	for _, b := range button {
		rlButton := numLights - b - 1 // because the buttons count l-->r but the bitmap is r-->
		state ^= int64(tools.PowInt(2, rlButton))
	}
	return state
}
