package tools

import (
	"golang.org/x/exp/constraints"
	"log"
	"math"
	"sort"
	"strconv"
)

type AnyInt interface {
	constraints.Signed | constraints.Unsigned
}

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

// inlineable atoi
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("bad atoi: %s", s)
	}
	return i
}

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Generic tools, from https://bitfieldconsulting.com/golang/functional

// Contains returns true if the slice contains the value
func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}

// Reverse a slice's order
func Reverse[E any](s []E) []E {
	result := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

// Sort a slice
func Sort[E constraints.Ordered](s []E) []E {
	result := make([]E, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// SumSlice sums any kind of integer value
func SumSlice[T AnyInt](measurements []T) T {
	var sum T
	for _, measurement := range measurements {
		sum += measurement
	}
	return sum
}

// KeepFunc is used in filtering
type KeepFunc[E any] func(E) bool

// Filter the slice to all values where the KeepFunc is true
func Filter[E any](s []E, f KeepFunc[E]) []E {
	result := []E{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func IsEven[T AnyInt](v T) bool {
	return v%2 == 0
}

type reduceFunc[E any] func(E, E) E

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init
	for _, v := range s {
		cur = f(cur, v)
	}
	return cur
}

func RemoveFromSlice[T comparable](theslice []T, doomed T) []T {
	newslice := []T{}
	for _, t := range theslice {
		if t != doomed {
			newslice = append(newslice, t)
		}
	}
	return newslice
}

// SliceSubtract returns a-b ; all elements in a that are not in b
func SliceSubtract[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	diff := []T{}
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func KeyExists[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func Triangular(n int) (triangle int) {
	for i := 1; i <= n; i++ {
		triangle += i
	}
	return
}

// these do no bounds checking. TODO make them return error
func InsertElt[T any](array []T, value T, index int) []T {
	return append(array[:index], append([]T{value}, array[index:]...)...)
}

func RemoveElt[T any](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}

func MoveElt[T any](array []T, srcIndex int, dstIndex int) []T {
	// clever appending from https://github.com/mnml/aoc/
	elt := array[srcIndex]
	array = append(array[:srcIndex], array[srcIndex+1:]...)
	array = append(array[:dstIndex], append([]T{elt}, array[dstIndex:]...)...)
	return array

}

func ReverseString(s string) string {
	return string(Reverse([]byte(s)))
}

func BaseConvert(x int, base int) []int {
	if x == 0 {
		return []int{0}
	}
	r := []int{}
	for x > 0 {
		r = append(r, x%base)
		x /= base
	}
	return r
}
