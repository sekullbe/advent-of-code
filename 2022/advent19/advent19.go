package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 24))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText, 32))
}

const (
	OREBOT = iota
	CLAYBOT
	OBSBOT
	GEODEBOT
)

type resources struct {
	ore    int
	clay   int
	obs    int
	geodes int
}

type bots struct {
	orebots   int
	claybots  int
	obsbots   int
	geodebots int
}

type state struct {
	timeRemaining int
	resources
	bots
}

// this will be pretty sparse so maybe I should make it simpler, but it looks so pretty and clean
type blueprint struct {
	number   int
	orebot   resources
	claybot  resources
	obsbot   resources
	geodebot resources
}

func tick(s state) state {
	s.timeRemaining -= 1
	s.clay += s.claybots
	s.ore += s.orebots
	s.obs += s.obsbots
	s.geodes += s.geodebots
	return s
}

var triangular [33]int

func init() {
	triangular[0] = 0
	for i := 1; i < 32; i++ {
		triangular[i] = triangular[i-1] + i
	}
}

var maximumbest = 0

// input in example is multline but real input is one line
// Blueprint 2: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obs robot costs 2 ore and 5 clay. Each geode robot costs 2 ore and 10 obs.
func parseBlueprint(line string) (bp blueprint) {
	n, err := fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&bp.number, &bp.orebot.ore, &bp.claybot.ore, &bp.obsbot.ore, &bp.obsbot.clay, &bp.geodebot.ore, &bp.geodebot.obs)

	if n != 7 || err != nil {
		log.Panicf("parse error, parsed %d elements, error:%s", n, err)
	}
	return
}

func (s *state) prune(goalbot int, bp blueprint, best int) bool {
	//  don't build orebot if we can't use more ore
	if goalbot == OREBOT && (s.orebots >= bp.orebot.ore && s.orebots > bp.claybot.ore && s.orebots > bp.obsbot.ore) {
		return true
	}
	// don't build claybot if we can't use more clay
	if goalbot == CLAYBOT && s.claybots >= bp.obsbot.clay {
		return true
	}
	// don't build obsbot if we have no claybots or if we can't use more
	if goalbot == OBSBOT && (s.claybots == 0 || s.obsbots >= bp.geodebot.obs) {
		return true
	}
	// don't build geodebot if we have no obsbots
	if goalbot == GEODEBOT && s.obsbots == 0 {
		return true
	}

	// punt if we can't catch up with the best yet
	now := s.geodes + (s.geodebots * s.timeRemaining)
	future := triangular[s.timeRemaining]

	if now+future < best {
		return true
	}

	return false
}

func simulate(s state, goalbot int, bp blueprint) int {

	if s.prune(goalbot, bp, maximumbest) { // this branch won't work, prune it
		return 0
	}

	best := 0
	for s.timeRemaining > 0 {

		if goalbot == OREBOT && s.ore >= bp.orebot.ore {
			ns := tick(s)
			ns.orebots = ns.orebots + 1
			ns.ore -= bp.orebot.ore
			for goal := 0; goal <= 3; goal++ {
				g := simulate(ns, goal, bp)
				best = tools.MaxInt(best, g)
			}
			return best
		}
		if goalbot == CLAYBOT && s.ore >= bp.claybot.ore {
			ns := tick(s)
			ns.claybots = ns.claybots + 1
			ns.ore -= bp.claybot.ore
			for goal := 0; goal <= 3; goal++ {
				g := simulate(ns, goal, bp)
				best = tools.MaxInt(best, g)
			}
			return best
		}
		if goalbot == OBSBOT && s.ore >= bp.obsbot.ore && s.clay >= bp.obsbot.clay {
			ns := tick(s)
			ns.obsbots = ns.obsbots + 1
			ns.ore -= bp.obsbot.ore
			ns.clay -= bp.obsbot.clay
			for goal := 0; goal <= 3; goal++ {
				g := simulate(ns, goal, bp)
				best = tools.MaxInt(best, g)
			}
			return best
		}
		if goalbot == GEODEBOT && s.ore >= bp.geodebot.ore && s.obs >= bp.geodebot.obs {
			ns := tick(s)
			ns.geodebots = ns.geodebots + 1
			ns.ore -= bp.geodebot.ore
			ns.obs -= bp.geodebot.obs
			for goal := 0; goal <= 3; goal++ {
				g := simulate(ns, goal, bp)
				best = tools.MaxInt(best, g)
			}
			return best
		}

		s = tick(s)
	}
	best = tools.MaxInt(s.geodes, best)
	maximumbest = tools.MaxInt(best, maximumbest)
	return best
}

func run1(inputText string, ticks int) int {
	var blueprints []blueprint

	for _, line := range parsers.SplitByLines(inputText) {
		blueprints = append(blueprints, parseBlueprint(line))
	}
	totalScore := 0
	for _, bp := range blueprints {
		best := 0
		for goal := 0; goal <= 3; goal++ {
			s := state{
				timeRemaining: ticks,
				resources:     resources{},
				bots:          bots{orebots: 1},
			}
			geodes := simulate(s, goal, bp)
			//log.Printf("BP %d produced %d geodes stating with goal %d", bp.number, best, goal)
			best = tools.MaxInt(best, geodes)
		}
		log.Printf("BP %d produced %d geodes", bp.number, best)
		totalScore += bp.number * best
	}

	return totalScore
}

func run2(inputText string, ticks int) int {
	var blueprints []blueprint

	for _, line := range parsers.SplitByLines(inputText) {
		blueprints = append(blueprints, parseBlueprint(line))
	}
	totalScore := 1
	for i, bp := range blueprints {
		maximumbest = 0
		best := 0
		for goal := 0; goal <= 3; goal++ {
			s := state{
				timeRemaining: ticks,
				resources:     resources{},
				bots:          bots{orebots: 1},
			}
			geodes := simulate(s, goal, bp)
			//log.Printf("BP %d produced %d geodes stating with goal %d", bp.number, best, goal)
			best = tools.MaxInt(best, geodes)
		}
		log.Printf("BP %d produced %d geodes", bp.number, best)
		totalScore *= best
		if i == 2 {
			break
		}
	}
	return totalScore
}
