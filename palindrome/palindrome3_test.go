package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindromeRecursive(value string, i int) bool {
	if i < len(value)/2 {

		if value[i] != value[len(value)-i-1] {
			return false
		} else {
			return isPalindromeRecursive(value, i+1)
		}
	} else {
		return true
	}
}

func isPalindrome3(value string) bool {
	return isPalindromeRecursive(value, 0)
}

func TestIsPalindrome3(t *testing.T) {
	assert.True(t, true, isPalindrome3("a"))
	assert.True(t, true, isPalindrome3("aba"))
	assert.True(t, true, isPalindrome3("kodok"))
	assert.True(t, true, isPalindrome3(""))

	assert.False(t, false, isPalindrome3("ab"))
	assert.False(t, false, isPalindrome3("abab"))
	assert.False(t, false, isPalindrome3("kodcok"))
	assert.False(t, false, isPalindrome3("tomi"))
}
