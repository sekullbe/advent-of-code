package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"strconv"
	"time"
)

//go:embed input.txt
var inputText string

const doLog = false

type stoneAndCount struct {
	stone, count int
}

var stoneScores map[stoneAndCount]int
var cacheHits, cacheMisses int

func main() {

	stoneScores = make(map[stoneAndCount]int)

	start := time.Now()
	fmt.Printf("Magic number: %d\n", run1(inputText))
	duration := time.Since(start)
	fmt.Printf("duration %d\n", duration/1_000)

	start = time.Now()
	fmt.Printf("Magic number new way: %d\n", run2(inputText, 25))
	duration = time.Since(start)
	fmt.Printf("duration %d\n", duration/1_000)

	fmt.Println("-------------")
	clear(stoneScores)

	start = time.Now()
	fmt.Printf("Magic number: %d\n", run2(inputText, 75))
	duration = time.Since(start)
	fmt.Printf("duration %d\n", duration/1_000)
	fmt.Printf("Cache Hits: %d\n", cacheHits)
	fmt.Printf("Cache Misses: %d\n", cacheMisses)
}

func run1(input string) int {
	stones := parsers.StringsToIntSlice(input)

	if doLog {
		log.Println(stones)
	}
	for i := 0; i < 25; i++ {
		stones = blink(stones)
		if doLog {
			log.Println(stones)
		}
	}
	return len(stones)
}

func run2(input string, iterations int) int {
	initStones := parsers.StringsToIntSlice(input)

	// let's go further with the idea of stones individually
	// we know that score([a,b]) == score([a]) + score([b])
	// so each time a stone splits, just recurse
	// and memoize for efficiency- that really matters!
	score := 0
	for _, stone := range initStones {
		score += scoreOneStone(stone, iterations)
	}

	return score

}

func scoreOneStone(stone int, iterations int) int {

	if iterations == 0 {
		return 1
	}
	score := 0
	for i := 0; i < iterations; i++ {
		if cacheScore, ok := stoneScores[stoneAndCount{stone, iterations}]; ok {
			cacheHits += 1
			return cacheScore
		}
		cacheMisses += 1
		newstones := blink([]int{stone})
		if len(newstones) == 1 {
			ns := scoreOneStone(newstones[0], iterations-1)
			stoneScores[stoneAndCount{newstones[0], iterations - 1}] = ns
			return ns
		} else {
			ns0 := scoreOneStone(newstones[0], iterations-1)
			stoneScores[stoneAndCount{newstones[0], iterations - 1}] = ns0
			ns1 := scoreOneStone(newstones[1], iterations-1)
			stoneScores[stoneAndCount{newstones[1], iterations - 1}] = ns1
			return ns0 + ns1
		}
	}
	return score
}

func blink(stones []int) []int {

	newstones := []int{}
	for _, stone := range stones {
		newstones = append(newstones, processOneStoneOnce(stone)...)
	}

	return newstones
}

func countDigits(num int) (count int) {
	s := strconv.Itoa(num)
	return len(s)
}

func splitDigits(num int) (left, right int) {
	s := strconv.Itoa(num)
	r := []rune(s)
	if len(s) == 1 {
		return num, 0
	}
	return tools.Atoi(string(r[0 : len(s)/2])), tools.Atoi(string(r[len(s)/2:]))
}

func processOneStoneOnce(stone int) []int {

	if stone == 0 {
		return []int{1}
	}
	if countDigits(stone)%2 == 0 {
		s1, s2 := splitDigits(stone)
		return []int{s1, s2}
	}
	return []int{stone * 2024}
}
