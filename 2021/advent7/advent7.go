package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
	"math"
	"sort"
)

//go:embed input.txt
var inputText string
var pyrMem FuncIntInt

func main() {
	// initialize
	pyrMem = memorized(func(n int) int {
		if n <= 1 {
			return 1
		}
		return n + pyrMem(n-1)
	})

	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	crabs := parsers.StringsWithCommasToIntSlice(inputText)
	sort.Ints(crabs)
	bestFuel := math.MaxInt
	target := -1
	for i := crabs[0]; i < crabs[len(crabs)-1]; i++ {
		fuelUse := sumsOfDifferences(i, crabs)
		if fuelUse < bestFuel {
			bestFuel = fuelUse
			target = i
		}
	}
	fmt.Printf("Best is location %d with %d fuel used\n", target, bestFuel)
	return bestFuel
}

func run2(inputText string) int {
	crabs := parsers.StringsWithCommasToIntSlice(inputText)
	sort.Ints(crabs)
	bestFuel := math.MaxInt
	target := -1
	for i := crabs[0]; i < crabs[len(crabs)-1]; i++ {
		fuelUse := pyramidSumsOfDifferences(i, crabs)
		if fuelUse < bestFuel {
			bestFuel = fuelUse
			target = i
		}
	}
	fmt.Printf("Best is location %d with %d fuel used\n", target, bestFuel)
	return bestFuel

}

// 1+2+3+...x
func pyramid(x int) int {
	if x < 0 {
		panic("x<0")
	}
	if x <= 1 {
		return x
	} else {
		return x + pyramid(x-1)
	}
}

type FuncIntInt func(int) int

func memorized(fn FuncIntInt) FuncIntInt {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1

	return func(input int) int {
		if val, found := cache[input]; found {
			//log.Println("Read from cache")
			return val
		}
		result := fn(input)
		cache[input] = result
		log.Printf("added n=%d result=%d to cache", input, result)
		return result
	}
}

func sumsOfDifferences(center int, nums []int) (sum int) {
	for _, num := range nums {
		sum += absint(num - center)
	}
	return
}

func pyramidSumsOfDifferences(center int, nums []int) (sum int) {
	for _, num := range nums {
		sum += pyrMem(absint(num - center))
		//sum += pyramid(absint(num - center))
	}
	return
}

func absint(i int) int {
	return int(math.Abs(float64(i)))
}
