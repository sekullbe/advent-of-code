package main

import (
	_ "embed"
	"fmt"
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

var sequenceCache = make(map[string][]int)

func run1(input string) int {
	sequenceCache = make(map[string][]int)
	sum := 0
	for _, line := range parsers.SplitByLines(input) {
		seq := parsers.StringsToIntSlice(line)
		seqs := repeatdiff(seq)
		e := extrapolate(seqs)
		sum += e
	}
	return sum
}

func run2(input string) int {
	sequenceCache = make(map[string][]int)
	sequenceCache = make(map[string][]int)
	sum := 0
	for _, line := range parsers.SplitByLines(input) {
		seq := parsers.StringsToIntSlice(line)
		seqs := repeatdiff(seq)
		e := backxtrapolate(seqs)
		sum += e
	}
	return sum
}

// given a slice of numbers produce the slice of differences
func diffsequence(seq []int) []int {

	if cs, ok := sequenceCache[fmt.Sprint(seq)]; ok {
		return cs
	}
	dseq := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		dseq[i] = seq[i+1] - seq[i]
	}
	sequenceCache[fmt.Sprint(seq)] = dseq
	return dseq
}

// repeats diffsequence until the final sequence is all 0s
// returns the original and all intermediate sequences
func repeatdiff(seq []int) [][]int {
	sequences := [][]int{seq} // can't assume length and make() it
	for {
		ns := diffsequence(sequences[len(sequences)-1])
		sequences = append(sequences, ns)
		if isZeroSlice(ns) {
			break
		}
	}
	return sequences
}

func extrapolate(sequences [][]int) int {
	// starting at the bottom...
	// add a zero to the last sequence, call it seq[n]
	// now append new value X to seq[n-1] such that X= lastElt(seq[n-1]) + lastElt(seq[n])
	// repeat up the chain and return X of the original seq[0]
	sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)

	newElt := 0
	for row := len(sequences) - 2; row >= 0; row-- {
		// could do this more efficiently without intermediates but it's easier to read and debug
		seq := sequences[row]
		seqBelow := sequences[row+1]
		newElt = tools.LastElt(seq) + tools.LastElt(seqBelow)
		sequences[row] = append(sequences[row], newElt)
	}
	return newElt
}

func backxtrapolate(sequences [][]int) int {
	// starting at the bottom...
	// *prepend* a zero to the last sequence, call it seq[n]
	// now append new value X to seq[n-1] such that X= lastElt(seq[n-1]) + lastElt(seq[n])
	// repeat up the chain and return X of the original seq[0]
	sequences[len(sequences)-1] = tools.PrependElt(sequences[len(sequences)-1], 0)
	newElt := 0
	for row := len(sequences) - 2; row >= 0; row-- {
		// could do this more efficiently without intermediates but it's easier to read and debug
		seq := sequences[row]
		seqBelow := sequences[row+1]
		newElt = seq[0] - seqBelow[0]
		sequences[row] = tools.PrependElt(sequences[row], newElt)
	}
	return newElt
}

func isZeroSlice(s []int) bool {
	if len(s) == 0 {
		panic("can't get zeroness of an empty slice")
	}
	for _, se := range s {
		if se != 0 {
			return false
		}
	}
	return true
}
