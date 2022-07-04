package palindrome

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(value string) bool {

	var temp string = ""

	for i := len(value) - 1; i >= 0; i-- {

		temp = temp + string(value[i])

	}

	fmt.Println(temp)

	return reflect.DeepEqual(value, temp)
}

func TestKodok(t *testing.T) {
	isPalindrome("kodok")
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, true, isPalindrome("a"))
	assert.True(t, true, isPalindrome("aba"))
	assert.True(t, true, isPalindrome("kodok"))
	assert.True(t, true, isPalindrome(""))

	assert.False(t, false, isPalindrome("ab"))
	assert.False(t, false, isPalindrome("abab"))
	assert.False(t, false, isPalindrome("kodcok"))
	assert.False(t, false, isPalindrome("tomi"))
}
