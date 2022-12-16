package main

import (
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"github.com/yourbasic/graph"
	"log"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type volcano struct {
	time                  int
	openValveRate         int // total rate of all open valves
	totalPressureRelieved int
	// or should these be not pointers?
	valves     []*valve
	valveMap   map[string]*valve
	valveIdMap map[int]*valve
	// or some kind of graph of the valves?
}

func newVolcano() volcano {
	return volcano{
		time:                  1,
		openValveRate:         0,
		totalPressureRelieved: 0,
		valves:                []*valve{},
		valveMap:              make(map[string]*valve),
		valveIdMap:            make(map[int]*valve),
	}
}

type valve struct {
	name     string
	id       int
	flowRate int
	// one or the other of these
	leadsTo      []*valve
	leadsToNames []string
	reachability map[*valve]int
}

func parseValve(line string) valve {
	//Valve RU has flow rate=0; tunnels lead to valves YH, ID
	var name string
	var flowRate int
	// can't use scanf to parse the variable bit at the end
	n, err := fmt.Sscanf(line, "Valve %s has flow rate=%d;", &name, &flowRate)
	if err != nil || n != 2 {
		log.Panicf("parse error, got %d/4, error:%s", n, err)
	}
	// we can't fill in leadsTo until we have all the valve objects! so there needs to be two passses
	// split the string on valve/valves, and split on , after that
	tmpa := strings.Split(line, "to valve")
	tmpb := strings.Fields(strings.ReplaceAll(tmpa[1], ",", " "))
	var ltn []string
	if tmpb[0] == "s" {
		ltn = tmpb[1:]
	} else {
		ltn = tmpb
	}

	return valve{
		name:         name,
		flowRate:     flowRate,
		leadsTo:      nil,
		leadsToNames: ltn,
		reachability: make(map[*valve]int),
	}
}

func parseAllValves(valveLines []string) volcano {
	vol := newVolcano()
	for i, s := range valveLines {
		v := parseValve(s)
		v.id = i
		// probably only going to need one of these
		vol.valves = append(vol.valves, &v)
		vol.valveMap[v.name] = &v
		vol.valveIdMap[i] = &v

	}
	// fix up the links
	for _, v := range vol.valves {
		for _, name := range v.leadsToNames {
			v.leadsTo = append(v.leadsTo, vol.valveMap[name])
		}

	}
	// add reachability from every valve to other valves
	// first set up the graph...
	g := graph.New(len(vol.valves))
	for id, vp := range vol.valveIdMap {
		for _, toP := range vp.leadsTo {
			g.AddBothCost(id, toP.id, 1)
		}
	}
	// ... and compute cost S->T for all S,T | S!=T
	for _, start := range vol.valves {
		_, dists := graph.ShortestPaths(g, start.id)
		for i := 0; i < len(dists); i++ {
			start.reachability[vol.valveIdMap[i]] = int(dists[i])
		}
		/*
			for _, target := range vol.valves {
				_, d := graph.ShortestPath(g, start.id, target.id)
				start.reachability[target] = int(d)
				//ps := lo.Map(p, func(id int, index int) string { return vol.valveIdMap[id].name }) //ids to names
				//log.Printf("path from %s to %s: %s\n", start.name, target.name, ps)
			}
		*/
	}

	return vol
}

// pressure is current total pressure expected
// currentflow is current flow/tick
// currentunnel is where we are now
// targets are all useful end valves
// limit is time limit (30 in part 1)
func (vol *volcano) findMaxPressure(tick int, pressure int, flow int, currentValve *valve, targets []*valve, timeLimit int) int {
	// what will we score if we stop looking now
	bestScoreYet := pressure + (timeLimit-tick)*flow
	// for every valve we can reach from here...
	for _, targetVP := range targets {
		// how long will it take us to get to it and open it
		timeSpend := currentValve.reachability[targetVP] + 1
		// is that too long?
		if tick+timeSpend > timeLimit {
			continue
		}
		// compute what happens if we do this move
		timeAfterMove := tick + timeSpend              // when will we get there and open it?
		pressureAfterMove := pressure + timeSpend*flow // while moving, rack up pressure
		flowAfterMove := flow + targetVP.flowRate
		// now that we've done that, what are our next moves starting from here
		scoreAfterMove := vol.findMaxPressure(timeAfterMove, pressureAfterMove, flowAfterMove, targetVP, tools.RemoveFromSlice(targets, targetVP), timeLimit)
		// and were they any good?
		if scoreAfterMove > bestScoreYet {
			bestScoreYet = scoreAfterMove // winnah!
		}
	}

	return bestScoreYet
}

func run1(inputText string) int {
	const timeLimit int = 30

	vol := parseAllValves(parsers.SplitByLines(inputText))
	// you never want to end up in a valve that does nothing, so don't bother with them as endpoints
	targets := lo.Filter(vol.valves, func(v *valve, idx int) bool { return v.flowRate > 0 })
	maxPressure := vol.findMaxPressure(0, 0, 0, vol.valveMap["AA"], targets, timeLimit)

	return maxPressure
}

func run2(inputText string) int {
	const timeLimit int = 26
	vol := parseAllValves(parsers.SplitByLines(inputText))
	_ = vol

	return 0
}
