package advent2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parsePasswordLine(t *testing.T) {

	s := "1-3 a: abcde"
	r, p, err := parsePasswordLine(s)
	assert.Nil(t, err)
	assert.Equal(t, p, "abcde")
	assert.Equal(t, r, PasswordRule{min: 1, max: 3, letter: "a"})

}

func Test_parsePasswordLine_error(t *testing.T) {

	s := "fgdgjfdhgkdf"
	_, _, err := parsePasswordLine(s)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Bad password string: fgdgjfdhgkdf")
	//assert.Equal(t, p, "abcde")
	//assert.Equal(t, r, PasswordRule{min: 1, max: 3, Letter: "a"})

}

func Test_checkPasswordOldRule(t *testing.T) {
	r := PasswordRule{min: 1, max: 3, letter: "a"}
	assert.Equal(t, true, r.checkPasswordOldRule("xxxaaXXX"))
	assert.Equal(t, true, r.checkPasswordOldRule("xxxaaaXXX"))
	assert.Equal(t, true, r.checkPasswordOldRule("xxxaXXX"))

	assert.Equal(t, false, r.checkPasswordOldRule("xxxXXX"))
	assert.Equal(t, false, r.checkPasswordOldRule("xxxaaaaXXX"))
}

func Test_checkPasswordNewRule(t *testing.T) {
	r := PasswordRule{min: 1, max: 3, letter: "a"}

	assert.Equal(t, true, r.checkPasswordNewRule("abcdefg"))
	assert.Equal(t, true, r.checkPasswordNewRule("bcadefg"))
	assert.Equal(t, false, r.checkPasswordNewRule("abacdefg"))
	assert.Equal(t, false, r.checkPasswordNewRule("bcdefgh"))

}
