package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/repeated-substring-pattern/

// First find all possible substring lengths by getting all dividers of the len(s) but len(s)
// Then for all of them check if repeating substr of these lengths
// going to create a string equal to the source one
func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}
	possibleSubstrLengths := getDivisors(len(s))
	for _, psl := range possibleSubstrLengths {
		ss := s[:psl]
		ok := true
		for i := 0; i < len(s)/psl; i++ {
			if s[i*psl:(i+1)*psl] != ss {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func getDivisors(n int) []int {
	var res []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if n/i != i && i != 1 {
				res = append(res, i, n/i)
			} else {
				res = append(res, i)
			}
		}
	}
	return res
}

func main() {
	//// Example 1
	//var s1 string = "abab"
	//fmt.Println("Expected: true Output: ", repeatedSubstringPattern(s1))

	// Example 2
	var s2 string = "aba"
	fmt.Println("Expected: false Output: ", repeatedSubstringPattern(s2))

	// Example 3
	var s3 string = "abcabcabcabc"
	fmt.Println("Expected: true Output: ", repeatedSubstringPattern(s3))

	// Example 4
	var s4 string = "a"
	fmt.Println("Expected: false Output: ", repeatedSubstringPattern(s4))

	// Example 5
	var s5 string = "aaaaaaaaaaaaa"
	fmt.Println("Expected: true Output: ", repeatedSubstringPattern(s5))

	// Example 6
	var s6 string = "bb"
	fmt.Println("Expected: true Output: ", repeatedSubstringPattern(s6))

}
