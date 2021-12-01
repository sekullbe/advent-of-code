package advent13

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

func Run() {
	run1(inputText)
	run2Faster(inputText)
}

func run1(inputText string) {

	shortestWait := math.MaxInt
	soonestBus := 0

	timestamp, busIds := parseInput(inputText)
	for _, id := range busIds {
		wait := calculateWaitFor(timestamp, id)
		if wait < shortestWait {
			shortestWait = wait
			soonestBus = id
		}
	}
	fmt.Printf("First bus is %d after %d minutes\n", soonestBus,shortestWait)
	fmt.Printf("Magic number is %d\n", soonestBus * shortestWait)
}

func calculateWaitFor(timestamp int, busId int) int {

	// last bus came this many min before timestamp
	// eg ts = 10, bus = 1, lbd = 1
	lastBusDelta := timestamp % busId
	if lastBusDelta == 0 {
		return 0 // just bail on this case instead of doing math
	}
	// yeah this is redundant i don't care
	nextBus := timestamp - lastBusDelta + busId
	return nextBus - timestamp

}

// return timestamp and slice of buses
func parseInput(inputText string) (int, []int) {
	var busIds []int
	inputLines := strings.Fields(inputText) // should split into ts and bus lines
	timestamp, err := strconv.Atoi(inputLines[0])
	if err != nil {
		panic("bad input parsing timestamp")
	}
	for _, s := range strings.Split(inputLines[1],",") {
		if s == "x" {
			continue
		}
		bus, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error parsing bus '%s'\n", s)
			continue
		}
		busIds = append(busIds, bus)
	}

	return timestamp,busIds
}


func run2(inputText string) int {
	busIds := parseInputWithGaps(inputText)

	// for every multiple M of the first bus ID
	// look at each number that isn't 0
	// see if M % id = idx
	// if no, continue

	// assumes first input is a bus, but in the input it is, so meh
	firstBus := busIds[0]
	var time int
	var oom float64
	for time = firstBus; ; time+=firstBus {
		//fmt.Print(".")
		if time > math.MaxInt / 10 {
			panic("that wasn't supposed to happen")
		}
		good := true
		for idx, bus := range busIds {
			if bus == 0 || idx == 0 { // skipping first bus here sets the index correctly
				continue
			}
			// look ahead at time+idx and see if the bus comes at that time
			busDelta := (time+idx) % bus
			if busDelta != 0 {
				good = false
				break
			}
		}
		if math.Log(float64(time)) > oom {
			oom +=1
			fmt.Print(".")
		}
		if good {
			break
		}
	}

	fmt.Printf("\nMagic number is %d\n", time)
	return time
}

func run2Faster(inputText string) int {
	busIds := parseInputWithGaps(inputText)

	// for every multiple M of the first bus ID
	// look at each number that isn't 0
	// see if M % id = idx
	// if no, continue

	// assumes first input is a bus, but in the input it is, so meh
	checkInterval := busIds[0]
	var time int
	dings := make(map[int]bool)
	dings[0] = true
	for time = checkInterval ; ; time+=checkInterval {
		//fmt.Print(".")

		good := true
		for idx, bus := range busIds {

			if bus == 0 || idx == 0 { // skipping first bus here sets the index correctly
				continue
			}
			// look ahead at time+idx and see if the bus comes at that time
			busDelta := (time+idx) % bus
			if busDelta != 0 {
				good = false
				break
			} else {
				// now we know to check every N minutes because it took N minutes to align what we have
				_, dingExists := dings[idx]
				if !dingExists {
					checkInterval = checkInterval * bus // don't multiply every time, only if we've dinged this match before
					dings[idx]= true
					fmt.Printf("DING %d bus=%d t=%d interval is now %d\n", idx, bus, time, checkInterval)
				}
				//fmt.Printf("ding%d t=%d\n", idx, time)

			}
		}
		if good {
			break
		}
	}

	fmt.Printf("\nMagic number is %d\n", time)
	return time
}

// return timestamp and slice of buses
func parseInputWithGaps(inputText string) []int {
	var busIds []int
	inputLines := strings.Fields(inputText) // should split into ts and bus lines
	for _, s := range strings.Split(inputLines[1],",") {
		if s == "x" {
			s = "0"
		}
		bus, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error parsing bus '%s'\n", s)
			continue
		}
		busIds = append(busIds, bus)
	}

	return busIds
}
