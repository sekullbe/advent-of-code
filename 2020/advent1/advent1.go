package advent1

import (
	_ "embed"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputs string

func Run() {

	inputInts := splitStringToInts(inputs)

	sort.Ints(inputInts)

	inputRev := make([]int, len(inputInts))
	copy(inputRev, inputInts)
	sort.Sort(sort.Reverse(sort.IntSlice(inputRev)))

	// fmt.Println(inputInts)
	// fmt.Println(inputRev)

	answer, err := findsum(inputInts, inputRev, 2020)
	if err == nil {
		fmt.Printf("The answer is: %d\n", answer)
	} else {
		panic(err)
	}

	answer2, err := findthree(inputInts, 2020)
	if err == nil {
		fmt.Printf("The answer is: %d\n", answer2)
	} else {
		panic(err)
	}

}

func splitStringToInts(str string) []int {

	var ints []int
	s := strings.Fields(str)
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		ints = append(ints, i)
	}
	return ints
}

func findsum(forward, reverse []int, desiredSum int) (int, error) {
	for _, a := range forward {
		for _, b := range reverse {
			sum := a + b
			if sum == desiredSum {
				//fmt.Printf("%d + %d = %d\n", a, b, sum)
				return a * b, nil
			}
		}
	}
	return 0, errors.New("No answer found")
}

func findthree(forward []int, desiredSum int) (int, error) {
	for _, a := range forward {
		for _, b := range forward {
			for _, c := range forward {
				if a == b || a == c || b == c {
					continue
				}
				sum := a + b + c
				if sum == desiredSum {
					fmt.Printf("%d + %d + %d = %d\n", a, b, c, sum)
					return a * b * c, nil
				}
			}
		}
	}
	return 0, errors.New("No answer found")
}
