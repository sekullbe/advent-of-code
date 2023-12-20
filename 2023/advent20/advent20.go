package main

import (
	_ "embed"
	"fmt"
	"github.com/oleiade/lane/v2"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	pulseQueue := lane.NewQueue[queuedPulse]()
	magic := make(map[string]int)
	pulsers := parseModules(parsers.SplitByLines(input))
	for i := 0; i < 1000; i++ {

		pulseQueue.Enqueue(queuedPulse{
			sourceName: "broadcaster", // doesn't matter what this is so long as it exists
			level:      LOW,
			target:     "broadcaster",
		})
		runPulses(pulsers, magic, pulseQueue, 0)
		//fmt.Println("-----------------")
	}
	var lows, highs int
	for _, p := range pulsers {
		lows += p.getLows()
		highs += p.getHighs()
	}

	fmt.Printf("lows=%d, highs=%d\n", lows, highs)
	return lows * highs // for sample 2 should be 4250, 2750
}
func run2(input string) int {
	pulseQueue := lane.NewQueue[queuedPulse]()

	pulsers := parseModules(parsers.SplitByLines(input))
	presses := 0
	/*
		From observing the input we see (bt, dl, rv, fr) -> rs ->rx
		rs needs all high outputs from those to rs to send low to rx
		so find the cycles of each of those, or at least the first press count where they go high

		This isn't quite a general solution, but the puzzle is specific enough in supplying four cycles that all
		feed rs that it wouldn't be very general anyway.
	*/
	magic := make(map[string]int)
	for {
		presses++
		pulseQueue.Enqueue(queuedPulse{
			sourceName: "broadcaster", // doesn't matter what this is so long as it exists
			level:      LOW,
			target:     "broadcaster",
		})
		magic := runPulses(pulsers, magic, pulseQueue, presses)
		if magic > 0 {
			return magic
		}
		//fmt.Println("-----------------")
	}

}

// given a queue, pop and queuedPulse, adding any new pulses to the queue
// not sure what to return yet
func runPulses(pulsers map[string]pulser, magic map[string]int, queue *lane.Queue[queuedPulse], presses int) int {

	for {
		nextPulse, ok := queue.Dequeue()
		if ok {
			nextPulses := pulsers[nextPulse.target].pulse(nextPulse.sourceName, nextPulse.level)
			for _, np := range nextPulses {
				// announce if it's one of the magic ones
				if presses > 0 && np.level == HIGH && (np.sourceName == "bt" || np.sourceName == "dl" || np.sourceName == "rv" || np.sourceName == "fr") {
					//fmt.Printf("Magic conjunction %s low\n", np.sourceName)
					if _, ok := magic[np.sourceName]; !ok {
						magic[np.sourceName] = presses
						fmt.Printf("Magic conjunction %s high at press %d\n", np.sourceName, presses)
					}
					if len(magic) == 4 {
						prod := 1
						for _, i := range magic {
							prod *= i
						}
						return prod
					}
				}

				queue.Enqueue(np)
			}
		} else {
			break
		}
	}
	return 0
}
