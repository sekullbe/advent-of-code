package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"strings"
	"time"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	// myDirKp -> coldDirKp -> radDirKp -> numKp

	numkp := initKeypad(numKeypad)
	radDirkp := initKeypad(dirKeypad)
	coldDirkp := initKeypad(dirKeypad)

	complexity := 0
	codes := parsers.SplitByLines(input)
	for _, code := range codes {

		codeRunes := []rune(code)
		radPresses := numkp.targetKeypadPresses(codeRunes)
		for _, radPress := range radPresses {
			coldPresses := radDirkp.targetKeypadPresses(radPress)
			for _, coldPress := range coldPresses {
				myPresses := coldDirkp.targetKeypadPresses(coldPress)
				complexCoeff := 0
				for _, p := range myPresses {
					printPressSequence(p)
					complexCoeff += len(p)
				}
				complexity += complexCoeff * numericPart(code)
			}
		}

		fmt.Println()
	}
	return complexity
}

func printPressSequence(p []rune) {
	for _, r := range p {
		fmt.Printf("%c", r)
	}
}

// always do > before ^ and v before <
func sortSequences(steps []rune) []rune {

	// if moving right, always move horizontally first
	// if moving left always move vertically first

	s := string(steps)
	for {
		olds := s
		s = strings.ReplaceAll(s, "^>", ">^")
		s = strings.ReplaceAll(s, "v>", ">v")
		s = strings.ReplaceAll(s, "<^", "^<")
		s = strings.ReplaceAll(s, "<v", "v<")

		if s == olds {
			break
		}
	}
	// it's always safe to go right and down
	// and safe to go left IF you've gone down first
	// and safe to go up IF you're gone right first
	// AND you'll never have right and left, or up and down, in the same list
	// so right, up, left, down

	return []rune(s)
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	return 0
}

const numKeypad = `789
456
123
#0A`

const dirKeypad = `#^A
<v>`

type fromto struct {
	from, to geometry.Point2
}

type keypad struct {
	layout    *grid.Board
	loc       geometry.Point2
	curKey    rune
	keyCache  map[rune]geometry.Point2
	pathCache map[fromto][]rune
}

func initKeypad(layout string) *keypad {
	b := grid.ParseBoardString(layout)
	aloc, err := b.Find('A')
	if err != nil {
		panic("no A!")
	}

	// set up the keyCache so we don't have to keep calling b.Find()
	keyCache := make(map[rune]geometry.Point2)
	for point, tile := range b.Grid {
		keyCache[tile.Contents] = point
	}

	keypad := &keypad{layout: b, loc: aloc, curKey: 'A', keyCache: keyCache, pathCache: make(map[fromto][]rune)}
	return keypad
}

func (k *keypad) press(r rune) []rune {

	pressStr := keysFromTo[keyPair{k.curKey, r}]
	presses := []rune(pressStr)
	k.curKey = r
	return presses

	/*
		// where are we going?
		toPt := k.keyCache[r]

		// check the cache
		if presses, ok := k.pathCache[fromto{k.loc, toPt}]; ok {
			return presses
		}

		// cache miss, compute the steps
		path, _, found := k.layout.FindPath(k.loc, toPt)
		if !found {
			panic("no path, that shouldn't be possible") // I believe in sanity clause
		}

		// convert the path of locations to a path of moves
		presses := grid.PathToSteps(path)
		presses = sortSequences(presses)
		// cache it
		k.pathCache[fromto{from: k.loc, to: toPt}] = presses

		// move the robot finger
		k.loc = toPt
		return presses

	*/
}

// first look at the target keypad and get the sequence of presses on a directional keypad to press each of those buttons
// for 029A on numeric keypad starting with loc='A' this will return <A, ^A, ^^>A, vvvA
func (k *keypad) targetKeypadPresses(buttons []rune) [][]rune {
	// i think what I am not taking into account is pushing A twice in a row
	allKeypresses := [][]rune{}
	for _, button := range buttons {
		presses := k.press(button)
		allKeypresses = append(allKeypresses, presses)
	}
	return allKeypresses
}

// given a bunch of presses necessary to move on some target keypad
// generate the presses for _this_ keypad to execute that, including the 'A' enter presses
// ie. we've generated ^^> for a step on the numeric keypad, so this will generate
// the dir steps for ^A^A>A
// not called? do i need it? or does targetKeypadPresses do the job?
func (k *keypad) executeTargetOnDirKeypad(targetMoves []rune) []rune {

	presses := []rune{}
	for _, move := range targetMoves {
		presses = append(presses, k.press(move)...)
		//presses = append(presses, k.press('A')...)
	}
	return presses
}

func numericPart(code string) int {
	c, _ := strings.CutSuffix(code, "A")
	return tools.Atoi(c)
}

type keyPair struct {
	first  rune
	second rune
}

var keysFromTo = map[keyPair]string{
	{'A', '0'}: "<A",
	{'0', 'A'}: ">A",
	{'A', '1'}: "^<<A",
	{'1', 'A'}: ">>vA",
	{'A', '2'}: "<^A",
	{'2', 'A'}: "v>A",
	{'A', '3'}: "^A",
	{'3', 'A'}: "vA",
	{'A', '4'}: "^^<<A",
	{'4', 'A'}: ">>vvA",
	{'A', '5'}: "<^^A",
	{'5', 'A'}: "vv>A",
	{'A', '6'}: "^^A",
	{'6', 'A'}: "vvA",
	{'A', '7'}: "^^^<<A",
	{'7', 'A'}: ">>vvvA",
	{'A', '8'}: "<^^^A",
	{'8', 'A'}: "vvv>A",
	{'A', '9'}: "^^^A",
	{'9', 'A'}: "vvvA",
	{'0', '1'}: "^<A",
	{'1', '0'}: ">vA",
	{'0', '2'}: "^A",
	{'2', '0'}: "vA",
	{'0', '3'}: "^>A",
	{'3', '0'}: "<vA",
	{'0', '4'}: "^<^A",
	{'4', '0'}: ">vvA",
	{'0', '5'}: "^^A",
	{'5', '0'}: "vvA",
	{'0', '6'}: "^^>A",
	{'6', '0'}: "<vvA",
	{'0', '7'}: "^^^<A",
	{'7', '0'}: ">vvvA",
	{'0', '8'}: "^^^A",
	{'8', '0'}: "vvvA",
	{'0', '9'}: "^^^>A",
	{'9', '0'}: "<vvvA",
	{'1', '2'}: ">A",
	{'2', '1'}: "<A",
	{'1', '3'}: ">>A",
	{'3', '1'}: "<<A",
	{'1', '4'}: "^A",
	{'4', '1'}: "vA",
	{'1', '5'}: "^>A",
	{'5', '1'}: "<vA",
	{'1', '6'}: "^>>A",
	{'6', '1'}: "<<vA",
	{'1', '7'}: "^^A",
	{'7', '1'}: "vvA",
	{'1', '8'}: "^^>A",
	{'8', '1'}: "<vvA",
	{'1', '9'}: "^^>>A",
	{'9', '1'}: "<<vvA",
	{'2', '3'}: ">A",
	{'3', '2'}: "<A",
	{'2', '4'}: "<^A",
	{'4', '2'}: "v>A",
	{'2', '5'}: "^A",
	{'5', '2'}: "vA",
	{'2', '6'}: "^>A",
	{'6', '2'}: "<vA",
	{'2', '7'}: "<^^A",
	{'7', '2'}: "vv>A",
	{'2', '8'}: "^^A",
	{'8', '2'}: "vvA",
	{'2', '9'}: "^^>A",
	{'9', '2'}: "<vvA",
	{'3', '4'}: "<<^A",
	{'4', '3'}: "v>>A",
	{'3', '5'}: "<^A",
	{'5', '3'}: "v>A",
	{'3', '6'}: "^A",
	{'6', '3'}: "vA",
	{'3', '7'}: "<<^^A",
	{'7', '3'}: "vv>>A",
	{'3', '8'}: "<^^A",
	{'8', '3'}: "vv>A",
	{'3', '9'}: "^^A",
	{'9', '3'}: "vvA",
	{'4', '5'}: ">A",
	{'5', '4'}: "<A",
	{'4', '6'}: ">>A",
	{'6', '4'}: "<<A",
	{'4', '7'}: "^A",
	{'7', '4'}: "vA",
	{'4', '8'}: "^>A",
	{'8', '4'}: "<vA",
	{'4', '9'}: "^>>A",
	{'9', '4'}: "<<vA",
	{'5', '6'}: ">A",
	{'6', '5'}: "<A",
	{'5', '7'}: "<^A",
	{'7', '5'}: "v>A",
	{'5', '8'}: "^A",
	{'8', '5'}: "vA",
	{'5', '9'}: "^>A",
	{'9', '5'}: "<vA",
	{'6', '7'}: "<<^A",
	{'7', '6'}: "v>>A",
	{'6', '8'}: "<^A",
	{'8', '6'}: "v>A",
	{'6', '9'}: "^A",
	{'9', '6'}: "vA",
	{'7', '8'}: ">A",
	{'8', '7'}: "<A",
	{'7', '9'}: ">>A",
	{'9', '7'}: "<<A",
	{'8', '9'}: ">A",
	{'9', '8'}: "<A",
	{'<', '^'}: ">^A",
	{'^', '<'}: "v<A",
	{'<', 'v'}: ">A",
	{'v', '<'}: "<A",
	{'<', '>'}: ">>A",
	{'>', '<'}: "<<A",
	{'<', 'A'}: ">>^A",
	{'A', '<'}: "v<<A",
	{'^', 'v'}: "vA",
	{'v', '^'}: "^A",
	{'^', '>'}: "v>A",
	{'>', '^'}: "<^A",
	{'^', 'A'}: ">A",
	{'A', '^'}: "<A",
	{'v', '>'}: ">A",
	{'>', 'v'}: "<A",
	{'v', 'A'}: "^>A",
	{'A', 'v'}: "<vA",
	{'>', 'A'}: "^A",
	{'A', '>'}: "vA",
	{'A', 'A'}: "A",
}
