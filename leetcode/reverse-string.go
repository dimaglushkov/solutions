package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/reverse-string/

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func main() {
	// Example 1
	var s1 = []byte{'h', 'e', 'l', 'a', 'o'}
	reverseString(s1)
	fmt.Println("Expected: [\"o\",\"l\",\"l\",\"e\",\"h\"]	Output: ", string(s1))

	// Example 2
	var s2 = []byte{'H', 'a', 'n', 'j', 'a', 'h'}
	reverseString(s2)
	fmt.Println("Expected: [\"h\",\"a\",\"n\",\"n\",\"a\",\"H\"]	Output: ", string(s2))

}
