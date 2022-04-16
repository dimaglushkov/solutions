package main

import "fmt"

// source: https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	var sp, tp int
	var sl, tl = len(s), len(t)
	if sl > tl {
		return false
	}
	for ; sp < sl && tp < tl; tp++ {
		if s[sp] == t[tp] {
			sp++
		}
	}
	return sp == sl
}

func main() {
	// Example 1
	var s1 string = "abc"
	var t1 string = "ahbgdc"
	fmt.Println("Expected: true	Output: ", isSubsequence(s1, t1))

	// Example 2
	var s2 string = "axc"
	var t2 string = "ahbgdc"
	fmt.Println("Expected: false	Output: ", isSubsequence(s2, t2))

	// Example 3
	var s3 string = ""
	var t3 string = "a"
	fmt.Println("Expected: true	Output: ", isSubsequence(s3, t3))

	// Example 4
	var s4 string = ""
	var t4 string = ""
	fmt.Println("Expected: true	Output: ", isSubsequence(s4, t4))

	// Example 5
	var s5 string = "aa"
	var t5 string = "a"
	fmt.Println("Expected: false	Output: ", isSubsequence(s5, t5))

	// Example 6
	var s6 string = "abc"
	var t6 string = "aaabbbccc"
	fmt.Println("Expected: true	Output: ", isSubsequence(s6, t6))

	// Example 7
	var s7 string = "x"
	var t7 string = "xxxxxxx"
	fmt.Println("Expected: true	Output: ", isSubsequence(s7, t7))

}
