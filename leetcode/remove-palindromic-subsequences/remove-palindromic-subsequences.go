package main

import "fmt"

// source: https://leetcode.com/problems/remove-palindromic-subsequences/

func reverseString(str string) string {
	s := []byte(str)
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return string(s)
}
func removePalindromeSub_(s string) int {
	if len(s) == 0 {
		return 0
	}
	if reverseString(s) == s {
		return 1
	}
	return 2
}

func isPalindrome(x string) bool {
	for i, _ := range x[:len(x)/2] {
		if x[i] != x[len(x)-1-i] {
			return false
		}
	}
	return true
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func main() {
	// Example 1
	var s1 string = "ababa"
	fmt.Println("Expected: 1	Output: ", removePalindromeSub(s1))

	// Example 2
	var s2 string = "abb"
	fmt.Println("Expected: 2	Output: ", removePalindromeSub(s2))

	// Example 3
	var s3 string = "baabb"
	fmt.Println("Expected: 2	Output: ", removePalindromeSub(s3))

}
