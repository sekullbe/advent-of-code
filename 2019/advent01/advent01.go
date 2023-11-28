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
	fmt.Printf("Magic number: %d\n", run1(parsers.StringsToIntSlice(inputText)))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(parsers.StringsToIntSlice(inputText)))
}

func run1(moduleMasses []int) int {

	totalFuel := 0
	for _, moduleMass := range moduleMasses {
		totalFuel += calculateFuel(moduleMass)
	}

	return totalFuel
}

func run2(moduleMasses []int) int {
	// (Calculate the fuel requirements for each module separately, then add them all up at the end.)
	totalFuel := 0
	for _, moduleMass := range moduleMasses {
		totalFuel += calculateFuelAccountingForFuel(moduleMass)
	}
	return totalFuel
}

func calculateFuel(mass int) int {
	return tools.MaxInt(mass/3-2, 0)
}

func calculateFuelAccountingForFuel(modMass int) int {
	fuelForModule := calculateFuel(modMass)
	fuelForFuel := calculateFuel(fuelForModule)
	newFuel := fuelForFuel
	for {
		newFuel = calculateFuel(newFuel)
		fuelForFuel += newFuel
		if newFuel == 0 {
			return fuelForFuel + fuelForModule
		}
	}
}
