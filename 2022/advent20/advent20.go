package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func floormod(n int, mod int) int {
	return (n%mod + mod) % mod
}

// Computes the 'real' index in slice s for supplied index idx
func computeNewIndex[T any](s []T, idx int) int {
	return floormod(idx, len(s))
}

func values(s []*int) []int {
	r := []int{}
	for _, p := range s {
		r = append(r, *p)
	}
	return r
}

func indexOf(s []*int, np *int) int {
	for i, i2 := range s {
		if i2 == np {
			return i
		}
	}
	return -1
}

func mix(key []*int, data []*int) []*int {
	// reworked this in part 2; had trouble with the data resetting while
	//being passed around. I need to grok slice representation better.
	for _, np := range key {
		// Find where in the data the key value is
		// appending magic from https://github.com/mnml/aoc/
		idx := indexOf(data, np)
		d := data[idx]
		data = append(data[:idx], data[idx+1:]...)
		newIdx := floormod(idx+*d, len(data))
		data = append(data[:newIdx], append([]*int{d}, data[newIdx:]...)...)
		if len(key) < 10 {
			log.Printf("key:%v elt:%2d file:%v", values(key), *np, values(data))
		}
	}
	return data
}

func loadFile(inputText string, multiplier int) []*int {
	var thefile []*int
	nums := parsers.StringsToIntSlice(inputText)
	for _, num := range nums {
		n := num * multiplier
		thefile = append(thefile, &n)
	}
	return thefile
}

func run1(inputText string) int {
	// what I want here is a slice of *pointers to integers
	key := loadFile(inputText, 1)
	data := make([]*int, len(key))
	copy(data, key) // why why why is it this order?
	decryption := mix(key, data)
	// find the 0
	zeroIdx := 0
	for i, num := range decryption {
		if *num == 0 {
			zeroIdx = i
			break
		}
	}
	a := decryption[computeNewIndex(decryption, zeroIdx+1000)]
	b := decryption[computeNewIndex(decryption, zeroIdx+2000)]
	c := decryption[computeNewIndex(decryption, zeroIdx+3000)]
	return *a + *b + *c
}

func run2(inputText string) int {
	// what I want here is a slice of *pointers to integers
	multiplier := 811589153
	key := loadFile(inputText, multiplier)
	decryption := make([]*int, len(key))
	copy(decryption, key) // why why why is it this order?
	for i := 0; i <= 9; i++ {
		decryption = mix(key, decryption)
	}

	// find the 0
	zeroIdx := 0
	for i, num := range decryption {
		if *num == 0 {
			zeroIdx = i
			break
		}
	}
	a := decryption[computeNewIndex(decryption, zeroIdx+1000)]
	b := decryption[computeNewIndex(decryption, zeroIdx+2000)]
	c := decryption[computeNewIndex(decryption, zeroIdx+3000)]
	return *a + *b + *c
}
