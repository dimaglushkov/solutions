package main

import "fmt"

// source: https://leetcode.com/problems/break-a-palindrome/

func breakPalindrome(palindrome string) string {
	n := len(palindrome)

	if n < 2 {
		return ""
	}

	i := 0
	for i < n && palindrome[i] == 'a' {
		i++
	}

	if i == n || (n%2 == 1 && i == n/2) {
		return palindrome[:n-1] + "b"
	}

	return palindrome[:i] + "a" + palindrome[i+1:]
}

func main() {
	var y = "aba"
	fmt.Println("Expected: \"abb\"	Output: ", breakPalindrome(y))

	var x = "aaaaa"
	fmt.Println("Expected: \"aaaab\"	Output: ", breakPalindrome(x))

	// Example 1
	var palindrome1 string = "abccba"
	fmt.Println("Expected: \"aaccba\"	Output: ", breakPalindrome(palindrome1))

	// Example 2
	var palindrome2 string = "a"
	fmt.Println("Expected: \"\"	Output: ", breakPalindrome(palindrome2))

}
