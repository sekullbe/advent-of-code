package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type rule struct {
	before, after int
}
type update []int

func run1(input string) int {
	inputComponents := parsers.SplitByEmptyNewlineToSlices(input)

	rules := parseRules(inputComponents[0])
	updates := inputComponents[1]

	// make a regexp for each rule: /Y\DX/ and if it matches the update string it's a violation
	sum := 0
	for _, u := range updates {
		pass := true
		for _, r := range rules {
			pass = pass && testRuleWithRegex(r, u)
		}
		if pass {
			sum += extractMiddleOfUpdate(u)
		}
	}
	return sum
}

func run2(input string) int {
	inputComponents := parsers.SplitByEmptyNewlineToSlices(input)

	rules := parseRules(inputComponents[0])
	updates := inputComponents[1]
	sum := 0
	for _, u := range updates {
		fixed := false
	fix:
		for _, r := range rules {
			if !testRuleWithRegex(r, u) { // only look at the failures
				// fix it by swapping the numbers in place
				//fmt.Printf("Fixing %s for rule %d|%d: ", u, r.before, r.after)
				// This would break very badly if some numbers were substrings of others.
				u = fixUpdate(r, u)
				//fmt.Println(u)
				fixed = true
				goto fix // this fix could have broken other rules, so start over. This risks infloop if two rules break each other so let's hope that doesn't happen
			}
		}
		if fixed {
			sum += extractMiddleOfUpdate(u)
		}
	}
	return sum
}

func parseRules(ruleLines []string) []rule {
	rules := []rule{}
	for _, line := range ruleLines {
		rule := rule{}
		_, err := fmt.Sscanf(line, "%d|%d", &rule.before, &rule.after)
		if err != nil {
			panic(err)
		}
		rules = append(rules, rule)
	}
	return rules
}

func parseUpdates(updateLines []string) []update {
	updates := []update{}
	for _, line := range updateLines {
		update := parsers.StringsWithCommasToIntSlice(line)
		updates = append(updates, update)
	}
	return updates
}

// returns true if the rule is NOT violated in the update
func testRuleWithRegex(rule rule, update string) bool {
	re := regexp.MustCompile(`,` + strconv.Itoa(rule.after) + `,.*?` + strconv.Itoa(rule.before) + `,`)
	return !re.MatchString("," + update + ",")
}

func extractMiddleOfUpdate(update string) int {
	updatePages := parsers.StringsWithCommasToIntSlice(update)
	return updatePages[len(updatePages)/2]
}

func fixUpdate(r rule, u string) string {
	replacer := strings.NewReplacer(strconv.Itoa(r.before), strconv.Itoa(r.after), strconv.Itoa(r.after), strconv.Itoa(r.before))
	return replacer.Replace(u)
}
