package main

import (
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func floormod(n int, mod int) int {
	/*
		Do the modulo who returns the negative value. (return -1)
		Add the value of the modulo (here 4, return 3).
		Re-do a second modulo, and now we are sure to have a positive value. (return 3)
	*/
	return ((n % mod) + mod) % mod
}

// Computes the 'real' index in slice s for supplied index idx
func computeNewIndex[T any](s []T, idx int) int {
	return floormod(idx, len(s))
	// if len is 10,  -1 ->9, -2 - 9, -10 -> 0
	// try idx = 1
	//if idx >= 0 {
	//return idx % (len(s))
	//} else {
	//return (len(s) + idx) % len(s)
	// len is 7, index of -2 is 5, index of -9 is 5, etc
	// 1 - 3 under mod 7 == -2 == 5... but we are zero based
	// 1 - 10 under mod 7 == -2 = 5... we want 4 which is maxidx -2 not 7-2
	// how about
	//negIdx := idx % len(s)
	//return (len(s) + negIdx) % len(s)
	//}
}

// Move element at index x 'move' places
func rotateElt[T any](s []T, idx int, move int) []T {
	if move == 0 {
		return s
	}
	// example: 1, -3, 2, 3, -2, 0, 4  . move -3 by -3, want 1, 2, 3, -2, -3, 0, 4. from index 1 to index 4
	// length here is 7  idx + move = 1-3 = -2. that's maxidx -2
	newIdx := computeNewIndex(s, idx+move)
	//if idx+move < -1 {
	//		newIdx -= 1
	//}
	//if idx+move > len(s)+1 {
	//	newIdx += 1
	//}
	if move < -1 && newIdx >= idx {
		newIdx -= 1
	}
	if move > 1 && newIdx < idx {
		newIdx += 1
	}

	return tools.MoveElt(s, idx, newIdx)
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

func mix(key []*int) []*int {
	data := make([]*int, len(key))
	copy(data, key) // why why why is it this order?
	for _, np := range key {
		// Find where in the data the key value is
		idx := indexOf(data, np)
		// and move that element *np places left/right
		data = rotateElt(data, idx, *np)
		//log.Printf("key:%v elt:%2d file:%v", values(key), *np, values(data))
	}
	return data
}

func loadFile(inputText string) []*int {
	var thefile []*int
	nums := parsers.StringsToIntSlice(inputText)
	for _, num := range nums {
		n := num
		thefile = append(thefile, &n)
	}
	return thefile
}

func run1(inputText string) int {
	// what I want here is a slice of *pointers to integers
	encryption := loadFile(inputText)
	decryption := mix(encryption)
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
	firstValues := loadFile(inputText)
	key := 811589153
	encryption := lo.Map(firstValues, func(n *int, idx int) *int { nn := *n; nn *= key; n = &nn; return n })
	decryption := mix(encryption)
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
