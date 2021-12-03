package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

type columnInfo struct {
	zeroes int
	ones   int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	nums := strings.Fields(inputText)
	gamma := calculateGamma(nums)
	epsilon := calculateEpsilon(nums)
	return gamma * epsilon
}

func run2(inputText string) int {
	nums := strings.Fields(inputText)
	oxygen := calculateOxygen(nums)
	scrubber := calculateScrubber(nums)
	return oxygen * scrubber
}

func calculateGamma(nums []string) int {
	// for each column in each string, see if there are more 1s or 0s
	// then the bit is the most of those
	var gamma int
	dataLength := len(nums[0])
	for i := 0; i < dataLength; i++ {
		zeroes, ones := countZerosAndOnesInColumn(nums, i)
		if ones >= zeroes {
			gamma = setBitNFromLeft(gamma, i, dataLength, 1)
		} else {
			gamma = setBitNFromLeft(gamma, i, dataLength, 0)
		}
	}
	return gamma
}

// this could be optimized, I don't care
func calculateEpsilon(nums []string) int {
	// for each column in each string, see if there are more 1s or 0s
	// then the bit is the most of those
	var epsilon int
	dataLength := len(nums[0])
	for i := 0; i < dataLength; i++ {
		zeroes, ones := countZerosAndOnesInColumn(nums, i)
		if ones >= zeroes {
			epsilon = setBitNFromLeft(epsilon, i, dataLength, 0)
		} else {
			epsilon = setBitNFromLeft(epsilon, i, dataLength, 1)
		}
	}
	return epsilon
}

func calculateOxygen(nums []string) int {
	// start from the left bit; count most common bit (1 wins ties)
	// keep all that match that bit
	// repeat with 2nd bit
	dataLength := len(nums[0])
	keptNumbers := createMapFromListOfNumbers(nums)
	for i := 0; i < dataLength; i++ {

		zeroes, ones := countZerosAndOnesInColumn(nums, i)
		zap := '0'
		if ones < zeroes {
			zap = '1'
		}
		for _, num := range nums {
			if rune(num[i]) == zap {
				keptNumbers[num] = false
			}
		}
		nums = extractRemainingNumbersFromMap(keptNumbers)
		if len(nums) == 1 {
			break
		}
	}
	// making some big assumptions here but the data should match
	oxygen, _ := strconv.ParseInt(extractRemainingNumbersFromMap(keptNumbers)[0], 2, 64)

	return int(oxygen)
}

func calculateScrubber(nums []string) int {
	// start from the left bit; count most common bit (1 wins ties)
	// keep all that match that bit
	// repeat with 2nd bit
	dataLength := len(nums[0])
	keptNumbers := createMapFromListOfNumbers(nums)
	for i := 0; i < dataLength; i++ {
		zeroes, ones := countZerosAndOnesInColumn(nums, i)
		zap := '1'
		if ones < zeroes {
			zap = '0'
		}
		for _, num := range nums {
			if rune(num[i]) == zap {
				keptNumbers[num] = false
			}
		}
		nums = extractRemainingNumbersFromMap(keptNumbers)
		if len(nums) == 1 {
			break
		}
	}
	// making some big assumptions here but the data should match
	scrubber, _ := strconv.ParseInt(extractRemainingNumbersFromMap(keptNumbers)[0], 2, 64)

	return int(scrubber)
}

func createMapFromListOfNumbers(nums []string) map[string]bool {
	keptNumbers := make(map[string]bool)
	for _, num := range nums {
		keptNumbers[num] = true
	}
	return keptNumbers
}

func extractRemainingNumbersFromMap(keptNumbers map[string]bool) (nums []string) {
	for k, v := range keptNumbers {
		if v == true {
			nums = append(nums, k)
		}
	}
	return nums
}

func setBitNFromLeft(num int, place int, length int, value int) int {
	num = num | 1<<(length-place-1)
	if value == 0 {
		newBit := 1 << (length - place - 1)
		num = num ^ newBit
	}
	return num
}

func countZerosAndOnesInColumn(nums []string, column int) (zeros, ones int) {
	for _, num := range nums {
		c := num[column]
		switch c {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}
	return
}
