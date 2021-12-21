package main

import (
	"fmt"
	"log"
	"testing"
)

func TestRun1(t *testing.T) {
	num := run1(inputText, 100)
	fmt.Printf("num: %d\n", num)
}

func TestRun2(t *testing.T) {
	p1w, p2w := run2DoIt(4, 8, 21)
	_ = p1w
	_ = p2w
	log.Println("done")
}
