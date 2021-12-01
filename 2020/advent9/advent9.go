package advent9

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

// store the preamble
// push new value & remove old value
// search it for pairs that add up to a value

// Or just scan from n to n-25 rather than tracking the preamble


func Run() {
	Run2Doit(inputText, 1124361034)
}

func Run1Doit(inputText string) {

	dataSetStr := strings.Fields(inputText)
	var dataSet []int
	for _, s := range dataSetStr {
		if i,err := strconv.Atoi(s); err == nil  {
			dataSet = append(dataSet, i)
		}
	}

	// indices of the preamble
	var p0 int = 0
	var p1 int = 25
	// current pointer, kind of redundant because it's always current

	for isSumOfPairInSlice(dataSet[p0:p1], dataSet[p1]) {
		p0++
		p1++
	}
	fmt.Printf("Crashed out at %d\n", dataSet[p1])
}

func Run2Doit(inputText string, target int) {

	dataSetStr := strings.Fields(inputText)
	var dataSet []int
	for _, s := range dataSetStr {
		if i,err := strconv.Atoi(s); err == nil  {
			dataSet = append(dataSet, i)
		}
	}
	contigs :=  findContiguousSum(dataSet, target)
	sort.Ints(contigs)
	sum := contigs[0] + contigs[len(contigs)-1]
	fmt.Printf("small + large = %d\n",sum)



}

func isSumOfPairInSlice(preamble []int, target int) bool {

	for idx1, a  := range preamble {
		for idx2, b := range preamble {
			if idx1 != idx2 && a != b && (a+b == target) {
				fmt.Printf("%d + %d = %d\n", a, b, target)
				return true
			}
		}
	}
	return false
}

func findContiguousSum(data []int, target int) []int {
	for idx1, a := range data {
		sum := a
		for idx2, b := range data[idx1+1:] {
			sum += b
			if sum == target {
				return data[idx1:idx1+idx2+1]
			}
			if sum > target {
				break
			}
		}
		if sum > target {
			continue
		}
	}
	return nil
}


