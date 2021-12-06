package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type fish int

func (f *fish) iterate() *fish {
	*f = *f - 1
	if *f == -1 {
		*f = 6
		var newFish fish = 8
		return &newFish
	}
	return nil
}

func run1(inputText string) int {

	return runMap(inputText, 80)
}

func run2(inputText string) int {

	return runMap(inputText, 256)
}

func runMap(inputText string, days int) int {
	//return run(inputText, 256)
	fishByDay := make(map[int]int)

	totalFish := 0
	for _, f := range parsers.StringsWithCommasToIntSlice(inputText) {
		fishByDay[f]++
		totalFish++
	}

	for time := 1; time <= days; time++ {
		newFishByDay := make(map[int]int)
		for i := 0; i <= 8; i++ {
			newFishByDay[i] = 0
		}
		newborns := 0
		for i := 0; i <= 8; i++ {
			if i == 0 {
				newborns = fishByDay[0]
				newFishByDay[6] = fishByDay[0]
				newFishByDay[8] = newborns
				totalFish += newborns
			} else {
				newFishByDay[i-1] += fishByDay[i]
			}
		}
		//fmt.Printf("After %d: %d fish (%d newborn)\n", time, totalFish, newborns)
		fishByDay = newFishByDay
	}

	return totalFish
}

// This does 80 days just fine but stalls out around day 175 on my macbook.
func runIteratively(inputText string, days int) int {
	var fishes []*fish
	for _, f := range parsers.StringsWithCommasToIntSlice(inputText) {
		fp := fish(f)
		fishes = append(fishes, &fp)
	}
	for time := 1; time <= days; time++ {
		newborns := []*fish{}
		for _, f := range fishes {
			baby := f.iterate()
			if baby != nil {
				newborns = append(newborns, baby)
			}
		}
		fishes = append(fishes, newborns...)
		//fmt.Printf("After %d: %d fish (%d newborn)\n", time, len(fishes), len(newborns))
	}
	return len(fishes)
}
