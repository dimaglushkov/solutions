package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/count-prefixes-of-a-given-string/

func countPrefixes(words []string, s string) int {
	var res int
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			res++
		}
	}

	return res
}

func main() {
	// Example 1
	var words1 = []string{"a", "b", "c", "ab", "bc", "abc"}
	var s1 string = "abc"
	fmt.Println("Expected: 3	Output: ", countPrefixes(words1, s1))

	// Example 2
	var words2 = []string{"a", "a"}
	var s2 string = "aa"
	fmt.Println("Expected: 2	Output: ", countPrefixes(words2, s2))

}
