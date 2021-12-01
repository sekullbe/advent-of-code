package advent10

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"

	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string



func Run() {
	Run1Doit(inputText)
	Run2Doit(inputText)
}

func Run1Doit(inputText string) {

	adapters := parsers.StringsToIntSlice(inputText)
	sort.Ints(adapters)

	// setup
	jolts := 0
	var jumps [4]int // ignore the 0 box
	deviceJolts := adapters[len(adapters)-1] + 3
	adapters = append(adapters, deviceJolts)

	fmt.Printf("Starting at 0 jolts, looking for device input at %d jolts\n", deviceJolts)
	for _, adapter := range adapters {
		jump := adapter - jolts
		//fmt.Printf("Adding %d jolt adapter for jump of %d: %d -> %d\n", adapter, jolts, jump, jolts + jump)
		jumps[jump]++
		jolts += jump
	}

	fmt.Printf("Jumped to final adapter at %d==%d\n", jolts, deviceJolts)
	fmt.Printf("1-jumps: %d\n", jumps[1])
	fmt.Printf("3-jumps: %d\n", jumps[3])
	fmt.Printf("product for some reason: %d\n", jumps[1] * jumps[3])
}

type memomap map[string]int

func Run2Doit(inputTest string) {
	adapters := parsers.StringsToIntSlice(inputText)
	sort.Ints(adapters)


	var memo = memomap{}
	count := memoCountPossibilities(adapters, 0, memo)
	fmt.Printf("Possibilities: %d\n", count)
	fmt.Printf("Entries in the memo map: %d\n", len(memo))


}

// idea here stolen from https://github.com/alexchao26/advent-of-code-go/blob/main/2020/day10/main.go
// rewritten a little so I feel better about stealing it

func memoCountPossibilities(nums []int, lastJolt int, memo memomap) int {
	// if in memo, return that value
	str := makeMemoKey(nums, lastJolt)
	if v, ok := memo[str]; ok {
		return v
	}

	// if all adapters used up, return 1
	if len(nums) == 0 {
		return 1
	}

	// create a recursive call for each adapter within 3 of the lastJoltage
	var count int
	for i, v := range nums {
		if v-lastJolt <= 3 {
			count += memoCountPossibilities(nums[i+1:], v, memo)
		} else { // stop counting if the joltage diff is too larger (>3)
			break
		}
	}

	// update memo
	memo[str] = count

	return count
}
func makeMemoKey(nums []int, lastJolt int) string {
	ans := strconv.Itoa(lastJolt) + "x"
	for _, v := range nums {
		ans += strconv.Itoa(v)
	}
	return ans
}
