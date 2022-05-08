package main

import "fmt"

// source: https://leetcode.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	var letters = make(map[rune]bool, len(s))
	var res int
	for _, r := range s {
		if _, ok := letters[r]; ok {
			res += 2
			delete(letters, r)
		} else {
			letters[r] = true
		}
	}
	for _ = range letters {
		return res + 1
	}
	return res
}

func main() {
	// Example 1
	var s1 string = "abccccdd"
	fmt.Println("Expected: 7	Output: ", longestPalindrome(s1))

	// Example 2
	var s2 string = "a"
	fmt.Println("Expected: 1	Output: ", longestPalindrome(s2))

	// Example 3
	var s3 string = "bb"
	fmt.Println("Expected: 2	Output: ", longestPalindrome(s3))

}
