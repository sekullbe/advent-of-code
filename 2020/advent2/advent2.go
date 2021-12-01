package advent2

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// of course it's not that easy, need to parse it into an array
//go:embed input.txt
var inputs string

type PasswordRule struct {
	min    int
	max    int
	letter string
}

func (r *PasswordRule) checkPasswordOldRule(password string) bool {
	// count how many letters in the string
	c := strings.Count(password, r.letter)
	return (c >= r.min && c <= r.max)
}

func (r *PasswordRule) checkPasswordNewRule(password string) bool {
	// min XOR max must contain the letter
	// count is 1-based not zero based

	pwRunes := []rune(password)
	//fmt.Println(pwRunes)
	firstRune := pwRunes[r.min-1]
	secondRune := pwRunes[r.max-1]

	requiredRune := []rune(r.letter)[0]

	//fmt.Printf("Comparing %c %c = %c\n", firstRune, secondRune, requiredRune)

	return (firstRune == requiredRune || secondRune == requiredRune) && !(firstRune == secondRune)
}

func Run() {

	var validPasswords int

	for _, v := range strings.Split(inputs, "\n") {
		if r, p, err := parsePasswordLine(v); err != nil {
			fmt.Println(err)
			panic(err)
		} else {
			passwordOK := r.checkPasswordNewRule(p)
			if passwordOK {
				validPasswords++
			}
		}

	}
	fmt.Printf("Valid Passwords: %d\n", validPasswords)

}

func parsePasswordLine(s string) (PasswordRule, string, error) {
	r := PasswordRule{}
	// eg "1-3 a: abcde"
	// split on space, drop the :
	pwElements := strings.Fields(s)
	if len(pwElements) != 3 {
		return r, "", errors.New("Bad password string: " + s)
	}
	pwRange := pwElements[0]
	pwChar := pwElements[1]
	password := pwElements[2]

	rangeElts := splitStringToInts(pwRange, "-")
	if len(rangeElts) != 2 {
		return r, "", errors.New("Bad password string: " + s)
	}
	r.min = rangeElts[0]
	r.max = rangeElts[1]

	r.letter = strings.TrimSuffix(pwChar, ":")

	return r, password, nil
}

func splitStringToInts(str string, sep string) []int {

	var ints []int
	s := strings.Split(str, sep)
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		ints = append(ints, i)
	}
	return ints
}
