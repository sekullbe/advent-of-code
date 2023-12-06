package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type race struct {
	duration       int
	recordDistance int
}

func run1(input string) int {
	winProd := 1
	races := parseRaces(parsers.SplitByLines(input))
	for i := 0; i < len(races); i++ {
		winProd *= evalOneRaceOptimized(races[i])
	}

	return winProd
}

func run2(input string) int {
	race := parseRacesToOneRace(parsers.SplitByLines(input))
	return evalOneRaceOptimized(race)
}

func evalOneRace(r race) int {
	wins := 0
	for hold := 1; hold < r.duration; hold++ { // we know hold = 0 and hold = duration are 0 and can't win
		if calculateDistance(hold, r.duration) > r.recordDistance {
			wins++
		}
	}
	return wins
}

// in tests this is about 30% faster
func evalOneRaceOptimized(r race) int {
	wins := 0
	everWon := false
	for hold := 1; hold < r.duration; hold++ { // we know hold = 0 and hold = duration are 0 and can't win
		if calculateDistance(hold, r.duration) > r.recordDistance {
			wins++
			everWon = true
		} else if everWon {
			break // if we've ever won and now we're losing, stop because we'll never win again
		}
	}
	return wins
}

func calculateDistance(hold int, dur int) (distance int) {
	if dur == 0 || dur == hold {
		return 0
	}
	return hold * (dur - hold)
}

func parseRaces(lines []string) []race {
	times := parsers.StringsToIntSlice(strings.Split(lines[0], ":")[1])
	dists := parsers.StringsToIntSlice(strings.Split(lines[1], ":")[1])
	races := []race{}
	for i := 0; i < len(times); i++ {
		races = append(races, race{duration: times[i], recordDistance: dists[i]})
	}
	return races
}

func parseRacesToOneRace(lines []string) race {
	timeStr := strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), "")
	distStr := strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), "")
	dur := tools.Atoi(timeStr)
	dist := tools.Atoi(distStr)
	return race{duration: dur, recordDistance: dist}
}
