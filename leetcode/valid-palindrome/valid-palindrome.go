package main

import (
	"fmt"
	"strings"
	"unicode"
)

// source: https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	rs := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			rs = append(rs, r)
		}
	}

	n := len(rs)
	for i := 0; i < n/2; i++ {
		if rs[i] != rs[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("true <> ", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("false <> ", isPalindrome("race a car"))
	fmt.Println("true <> ", isPalindrome(" "))
}
