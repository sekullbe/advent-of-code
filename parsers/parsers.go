package parsers

import (
	"regexp"
	"strconv"
	"strings"
)

func StringsToIntSlice(inputText string) []int {
	dataSetStr := strings.Fields(inputText)
	var dataSet []int
	for _, s := range dataSetStr {
		if i, err := strconv.Atoi(s); err == nil {
			dataSet = append(dataSet, i)
		}
	}
	return dataSet
}

func StringsWithCommasToIntSlice(inputText string) []int {
	dataSetStr := strings.Split(inputText, ",")
	var dataSet []int
	for _, s := range dataSetStr {
		if i, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
			dataSet = append(dataSet, i)
		}
	}
	return dataSet
}

func SplitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
}

func SplitByLines(str string) []string {
	return strings.Split(str, "\n")
}
