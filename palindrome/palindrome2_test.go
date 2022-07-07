package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome2(value string) bool {

	for i := 0; i < len(value); i++ {
		var firstIndex int
		var lastIndex int

		firstIndex = i
		lastIndex = len(value) - i - 1

		if value[firstIndex] != value[lastIndex] {
			return false
		}
	}

	return true
}

func TestKodok2(t *testing.T) {
	isPalindrome2("aaaaaaaaaaaa")
}

func TestIsPalindrome2(t *testing.T) {
	assert.True(t, true, isPalindrome2("a"))
	assert.True(t, true, isPalindrome2("aba"))
	assert.True(t, true, isPalindrome2("kodok"))
	assert.True(t, true, isPalindrome2(""))

	assert.False(t, false, isPalindrome2("ab"))
	assert.False(t, false, isPalindrome2("abab"))
	assert.False(t, false, isPalindrome2("kodcok"))
	assert.False(t, false, isPalindrome2("tomi"))
}
