package fizzbuzz

import (
	"fmt"
	"testing"
)

func fizzBuzz(total int) {
	var i int
	for i = 1; i <= total; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}
}

func TestFizzBuzz(t *testing.T) {
	fizzBuzz(100)
}
