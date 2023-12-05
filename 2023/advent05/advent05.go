package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"maps"
	"math"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

var processingOrder = []string{"seed-to-soil map", "soil-to-fertilizer map", "fertilizer-to-water map", "water-to-light map",
	"light-to-temperature map", "temperature-to-humidity map", "humidity-to-location map"}

func run1(input string) int {

	seedNums, mapOfSections := parseAllMapSectionsToStructs(input)
	minLoc := math.MaxInt
	for _, seed := range seedNums {
		working := seed
		for _, mapName := range processingOrder {
			working = mapSectionLookup(working, mapOfSections[mapName])
		}
		minLoc = min(working, minLoc)
	}

	return minLoc

}

func run1BruteForce(input string) int {

	seedNums, mapOfMaps := parseAllMaps(input)
	// hmmm is it useful to store the intermediate values in a map? i doubt it
	minLoc := math.MaxInt
	for _, seed := range seedNums {
		working := seed
		for _, mapName := range processingOrder {
			working = mapLookupWithDefault(working, mapOfMaps[mapName])
		}
		minLoc = min(working, minLoc)
	}

	return minLoc
}

func run2(inputText string) int {

	return 0
}

// preloading all the maps works and solves the sample problem, but it is dog slow and uses vast memory. do better.

// return a map of maps? not that crazy...
func parseAllMaps(allLines string) (seedNums []int, mapOfMaps map[string]map[int]int) {
	var sections [][]string = parsers.SplitByEmptyNewlineToSlices(allLines)
	mapOfMaps = make(map[string]map[int]int)
	for i, section := range sections {
		if i == 0 {
			seedNums = parsers.StringsToIntSlice(strings.Split(section[0], ":")[1])
			continue
		}
		sectionName := strings.TrimSuffix(section[0], ":")
		sectionMap := parseMapSection(section[1:])
		mapOfMaps[sectionName] = sectionMap
	}
	return seedNums, mapOfMaps
}

func parseMapSection(lines []string) map[int]int {
	sectionMap := make(map[int]int)
	for _, line := range lines {
		m := parseMapLine(line)
		maps.Copy(sectionMap, m)
	}
	return sectionMap
}

func parseMapLine(mapstr string) map[int]int {
	var destStart, sourceStart, rangeLength int
	n, err := fmt.Sscanf(mapstr, "%d %d %d", &destStart, &sourceStart, &rangeLength)
	if n != 3 || err != nil {
		panic("parse error: " + mapstr)
	}
	outMap := make(map[int]int)
	for i := 0; i < +rangeLength; i++ {
		outMap[sourceStart+i] = destStart + i
	}
	return outMap
}

// echoes the key if the key is not in the map
func mapLookupWithDefault(key int, theMap map[int]int) int {
	val, ok := theMap[key]
	if ok {
		return val
	}
	return key
}

type mapSection struct {
	sourceStart int
	destStart   int
	rangeLength int
}

func parseMapSectionToStructs(lines []string) []mapSection {
	var destStart, sourceStart, rangeLength int
	mapSections := []mapSection{}
	for _, mapstr := range lines {
		n, err := fmt.Sscanf(mapstr, "%d %d %d", &destStart, &sourceStart, &rangeLength)
		if n != 3 || err != nil {
			panic("parse error: " + mapstr)
		}
		mapSections = append(mapSections, mapSection{sourceStart, destStart, rangeLength})
	}
	return mapSections
}

func parseAllMapSectionsToStructs(allLines string) (seedNums []int, mapOfSections map[string][]mapSection) {
	var sections [][]string = parsers.SplitByEmptyNewlineToSlices(allLines)
	mapOfSections = make(map[string][]mapSection)
	for i, section := range sections {
		if i == 0 {
			seedNums = parsers.StringsToIntSlice(strings.Split(section[0], ":")[1])
			continue
		}
		sectionName := strings.TrimSuffix(section[0], ":")
		sectionMap := parseMapSectionToStructs(section[1:])
		mapOfSections[sectionName] = sectionMap
	}
	return seedNums, mapOfSections
}

func mapSectionLookup(key int, sections []mapSection) int {
	for _, section := range sections {
		val := mapSectionLookupOneSection(key, section)
		if val != key {
			return val
		}
	}
	return key
}

func mapSectionLookupOneSection(key int, section mapSection) int {
	if key >= section.sourceStart && key < section.sourceStart+(section.rangeLength) {
		return section.destStart + (key - section.sourceStart)
	}
	return key
}
