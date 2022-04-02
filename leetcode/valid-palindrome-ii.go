package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/valid-palindrome-ii/

func isPalindrome(s string, skip bool) bool {
	if len(s) < 2 {
		return true
	}
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			if skip {
				return isPalindrome(s[l+1:r+1], false) || isPalindrome(s[l:r], false)
			}
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	return isPalindrome(s, true)
}

func main() {
	var s0 = "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	fmt.Println("Expected: true	Output: ", validPalindrome(s0))

	// Example 1
	var s1 string = "aba"
	fmt.Println("Expected: true	Output: ", validPalindrome(s1))

	// Example 2
	var s2 string = "abca"
	fmt.Println("Expected: true	Output: ", validPalindrome(s2))

	// Example 3
	var s3 string = "abc"
	fmt.Println("Expected: false	Output: ", validPalindrome(s3))

	// Example 4
	var s4 string = ""
	fmt.Println("Expected: true	Output: ", validPalindrome(s4))

	// Example 5
	var s5 string = "a"
	fmt.Println("Expected: true	Output: ", validPalindrome(s5))

	// Example 6
	var s6 string = "aaaa"
	fmt.Println("Expected: true	Output: ", validPalindrome(s6))

	// Example 7
	var s7 string = "aaaaa"
	fmt.Println("Expected: true	Output: ", validPalindrome(s7))

	// Example 8
	var s8 string = "abbcda"
	fmt.Println("Expected: false	Output: ", validPalindrome(s8))

	// Example 9
	var s9 string = "ab"
	fmt.Println("Expected: true	Output: ", validPalindrome(s9))
}
