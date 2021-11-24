package advent1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findsum(t *testing.T) {

	i := []int{1, 3, 5, 7, 8}
	r := []int{8, 7, 5, 3, 1}

	ans, err := findsum(i, r, 12)
	assert.Nil(t, err)
	assert.Equal(t, 35, ans)

}

func Test_findthree(t *testing.T) {

	i := []int{1, 3, 5, 7, 9}

	ans, err := findthree(i, 21)
	assert.Nil(t, err)
	assert.Equal(t, 9*7*5, ans)

}
