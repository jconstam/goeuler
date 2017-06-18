package goeuler

import (
	"fmt"
	"math"
)

type problem00004 struct {
	largestPalindrome float64
}

func (prob problem00004) getInfo() (string, string) {
	return "Largest palindrome product", "A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.\nFind the largest palindrome made from the product of two 3-digit numbers."
}

func (prob *problem00004) solve() {
	for digit1 := 999; digit1 >= 100; digit1-- {
		for digit2 := 999; digit2 >= 100; digit2-- {
			if isPalindrome(digit1 * digit2) {
				prob.largestPalindrome = math.Max(float64(digit1*digit2), prob.largestPalindrome)
			}
		}
	}
}

func isPalindrome(value int) bool {
	palindrome := true
	valAsString := fmt.Sprintf("%v", value)

	for index := 0; index < len(valAsString); index++ {
		reverseIndex := len(valAsString) - index - 1

		if reverseIndex < index {
			break
		} else if string(valAsString[index]) != string(valAsString[reverseIndex]) {
			palindrome = false
			break
		}
	}

	return palindrome
}

func (prob problem00004) getExpectedResult() float64 {
	return float64(906609)
}

func (prob problem00004) getActualResult() float64 {
	return prob.largestPalindrome
}
