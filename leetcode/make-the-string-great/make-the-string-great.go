package main

import "fmt"

// source: https://leetcode.com/problems/make-the-string-great/

func makeGood(s string) string {
	const diff = 'a' - 'A'
	improve := func(x string) string {
		n := len(x)
		for i := 0; i < n-1; i++ {
			if x[i]-x[i+1] == diff || x[i+1]-x[i] == diff {
				return x[:i] + x[i+2:]
			}
		}
		return x
	}

	x := improve(s)
	for s != x {
		s = x
		x = improve(s)
	}

	return s

}

func main() {
	var s0 string = "Pp"
	fmt.Println("Expected: 	Output: ", makeGood(s0))

	// Example 1
	var s1 string = "leEeetcode"
	fmt.Println("Expected: leetcode	Output: ", makeGood(s1))

	// Example 2
	var s2 string = "abBAcC"
	fmt.Println("Expected: 	Output: ", makeGood(s2))

	// Example 3
	var s3 string = "s"
	fmt.Println("Expected: s	Output: ", makeGood(s3))

}
