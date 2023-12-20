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

	pulsers := parseModules(parsers.SplitByLines(input))
	for i := 0; i < 1000; i++ {
		//fmt.Println("button -low -> broadcaster")
		//pulsers["broadcaster"].queuedPulse("button", LOW)

		pulseQueue.Enqueue(queuedPulse{
			sourceName: "broadcaster", // doesn't matter what this is so long as it exists
			level:      LOW,
			target:     "broadcaster",
		})
		runPulses(pulsers, pulseQueue)
		//fmt.Println("-----------------")
	}
	var lows, highs int
	for _, p := range pulsers {
		lows += p.getLows()
		highs += p.getHighs()
	}

	// +1 for the initial button sending a LOW queuedPulse
	fmt.Printf("lows=%d, highs=%d\n", lows, highs)
	return lows * highs // for sample 2 should be 4250, 2750
}
func run2(input string) int {

	return 0
}

// given a queue, pop and queuedPulse, adding any new pulses to the queue
// not sure what to return yet
func runPulses(pulsers map[string]pulser, queue *lane.Queue[queuedPulse]) {
	for {
		nextPulse, ok := queue.Dequeue()
		if ok {
			nextPulses := pulsers[nextPulse.target].pulse(nextPulse.sourceName, nextPulse.level)
			for _, np := range nextPulses {
				queue.Enqueue(np)
			}
		} else {
			break
		}
	}
}
