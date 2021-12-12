package tools

import "math"

// yes I know 'util' etc is bad practice as a package
// so maybe I'll refactor this as I split things up

func CountElementsInString(elts []rune, s string) (count int) {
	for _, r := range s {
		if ContainsRune(elts, r) {
			count++
		}
	}
	return count
}

func ContainsInt(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func ContainsRune(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}