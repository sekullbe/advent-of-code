package advent12

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

const (
	N int = iota
	E
	S
	W
)

func Run() {
	run1(inputText)
	run2(inputText)
}

func run1(inputText string) ship {

	// define starting location as 0,0 - that's in the ship already
	minnow := ship{x: 0, y: 0, facing: E}

	for _, s := range strings.Fields(inputText) {
		instr, arg := parseCommand(s)
		switch instr {
		case "R":
			minnow.rotate(true, arg)
		case "L":
			minnow.rotate(false, arg)
		case "F":
			minnow.moveForward(arg)
		case "N":
			minnow.move(N, arg)
		case "E":
			minnow.move(E, arg)
		case "S":
			minnow.move(S, arg)
		case "W":
			minnow.move(W, arg)
		default:
			panic("bad command: "+ s)
		}
	}

	fmt.Printf("Movement complete. Ship is at (%d,%d)\n", minnow.x, minnow.y)
	fmt.Printf("Distance = %0.0f\n", math.Abs(float64(minnow.x)) + math.Abs(float64(minnow.y)))
	return minnow
}

func run2(inputText string) ship {

	minnow := ship{x: 0, y: 0, facing: E}
	waypoint := &waypoint{x: 10, y: 1}

	for _, s := range strings.Fields(inputText) {
		instr, arg := parseCommand(s)
		switch instr {
		case "R":
			waypoint.rotate(true, arg)
		case "L":
			waypoint.rotate(false, arg)
		case "F":
			minnow.moveToWaypoint(waypoint, arg)
		case "N":
			waypoint.move(N, arg)
		case "E":
			waypoint.move(E, arg)
		case "S":
			waypoint.move(S, arg)
		case "W":
			waypoint.move(W, arg)
		default:
			panic("bad command: "+ s)
		}
	}

	fmt.Printf("Movement complete. Ship is at (%d,%d)\n", minnow.x, minnow.y)
	fmt.Printf("Distance = %0.0f\n", math.Abs(float64(minnow.x)) + math.Abs(float64(minnow.y)))
	return minnow
}

func parseCommand(command string) (instr string, arg int) {
	instr = command[0:1]
	arg,err := strconv.Atoi(command[1:])
	if err != nil {
		panic("bad parse")
	}
	return instr,arg
}
