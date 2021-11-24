package advent4

import (
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed input.txt
var inputs string



func Run() {

	passports := parsePassports(inputs)

	var validCount int
	for _, passport := range passports {
		if passport.isValid() {
			validCount++
		}
	}

	fmt.Printf("Valid passports: %d\n", validCount)
}


func parsePassports(inputs string) []passportValues {

	// passports are separated by newlines
	// lines are key:valvalval , space separated
	var passports []passportValues

	passportChunks := SplitByEmptyNewline(inputs)

	for _, chunk := range passportChunks {
		pass := newPassportValues(chunk)
		passports = append(passports, *pass)
	}

	_ = passportChunks

	return passports
}

func SplitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}
