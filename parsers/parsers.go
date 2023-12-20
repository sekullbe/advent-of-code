package parsers

import (
	"regexp"
	"strconv"
	"strings"
)

// TODO Many of the things here can be done with fmt.Fscanf, fmt.Sscan, etc.

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

func StringsWithCommasToInt64Slice(inputText string) []int64 {
	dataSetStr := strings.Split(inputText, ",")
	var dataSet []int64
	for _, s := range dataSetStr {
		if i, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
			dataSet = append(dataSet, int64(i))
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

func SplitByEmptyNewlineToSlices(str string) [][]string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	strGroups := regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
	var ret [][]string
	for _, group := range strGroups {
		splitGroup := SplitByLines(group)
		ret = append(ret, splitGroup)
	}
	return ret
}

func SplitByLines(str string) []string {
	return strings.Split(strings.TrimSpace(str), "\n")
}

func SplitByLinesNoTrim(str string) []string {
	return strings.Split(str, "\n")
}

func SplitByCommasAndTrim(str string) []string {
	return strings.Split(strings.ReplaceAll(str, " ", ""), ",")
}
