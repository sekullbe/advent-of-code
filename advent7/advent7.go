package advent7

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/yourbasic/graph"
)

//go:embed input.txt
var rulesText string

type contents struct {
	bagName string
	bagNumber int64
}

func newContents(bagName string, bagNumber int64) *contents {
	c := contents{bagName: bagName, bagNumber: bagNumber}
	return &c
}

func Run1() {

	rulesText = strings.TrimSuffix(rulesText, "\n")
	rules := splitRules(rulesText)
	n2i, i2n := createNodeNameMap(rules)
	rulesGraph := parseRulesToGraph(rules, n2i)

	// now, for each bag, do a search for "shiny gold"
	indexOfTarget := n2i["shiny gold"]
	count := 0
	for index, _ := range i2n {
		_, connected := graph.ShortestPath(rulesGraph, index, indexOfTarget)
		if connected > 0 {
			count++
		}
	}

	fmt.Printf("Bags that can contain 'shiny gold': %d\n", count)
}

func Run2() {
	count := Run2_doit(rulesText)
	fmt.Printf("Total bags in shiny gold: %d\n", count)
	// 7794 is too low
}

func Run2_doit(theRules string) int64 {
	theRules = strings.TrimSuffix(theRules, "\n")
	rules := splitRules(theRules)
	n2i, _ := createNodeNameMap(rules)
	rulesGraph := parseRulesToGraph(rules, n2i)

	var bagCount int64
	bagCount = countSubBags(n2i["shiny gold"], rulesGraph)
	return bagCount -1
}

func countSubBags(v int, g *graph.Mutable) int64 {
	var bagCount int64 = 1

	g.Visit(v, func(w int, c int64) (skip bool) {
		bagCount += c * countSubBags(w,g)
		return
	})
	return bagCount
}


func splitRules (rulesText string) []string {
	rules := strings.Split(rulesText, "\n")
	return rules
}

func parseRulesToGraph(rules []string, n2i map[string]int) *graph.Mutable {
	g := graph.New(len(rules))
	// a rule looks like this:
	// dull tan bags contain 4 faded blue bags, 3 faded olive bags, 5 dull salmon bags.
	// Nodes are numbered not named
	for _, rule := range rules {
		containerName := parseContainerFromRule(rule)
		contents := parseContentsFromRule(rule)
		for _, con := range contents {
			g.AddCost(n2i[containerName], n2i[con.bagName], con.bagNumber)
		}
	}
	return g
}

func createNodeNameMap(rules []string) (map[string]int, map[int]string) {
	nameToIndex := make(map[string]int)
	indexToName := make(map[int]string)
	for i, rule := range rules {
		bagname := parseContainerFromRule(rule)
		nameToIndex[bagname] = i
		indexToName[i] = bagname
	}
	return nameToIndex, indexToName

}
func parseContainerFromRule(rule string) string {
	return rule[:strings.Index(rule, " bags")]
}

func parseContentsFromRule(rule string) []contents {
	// eg  dull tan bags contain 4 faded blue bags, 3 faded olive bags, 5 dull salmon bags.
	ruleSegmentContentsSlice := strings.Split(rule, "contain ")
	ruleSegmentContents := ruleSegmentContentsSlice[1]
	if ruleSegmentContents == "no other bags." {
		return make([]contents, 0)
	}
	contentStrings := strings.Split(ruleSegmentContents, ", ")
	// parse each element into a 'contents'- they look like "3 faded olive bags"
	cons := make([]contents, 0)
	for _, contentString := range contentStrings {
		re := regexp.MustCompile("^(\\d) (.*?) bags?\\.?$")
		matches := re.FindStringSubmatch(contentString)
		num, _ := strconv.Atoi(matches[1])
		con := newContents(matches[2], int64(num))
		cons = append(cons, *con)
	}
	return cons
}
