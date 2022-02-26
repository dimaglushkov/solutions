package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/implement-strstr/

/*func strStr(haystack string, needle string) int {
	nLen := len(needle)
	res := -1
	curNeedleI := 0

	if nLen == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		for curNeedleI < nLen && i < len(haystack) && haystack[i] == needle[curNeedleI] {
			curNeedleI++
			i++
		}
		if curNeedleI == nLen {
			res = i - curNeedleI
			break
		} else {
			i -= curNeedleI
		}
		curNeedleI = 0
	}

	return res
}*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Index(haystack, needle)
}

func main() {
	// Example 1
	var haystack1 string = "hello"
	var needle1 string = "ll"
	fmt.Println("Expected: 2 Output: ", strStr(haystack1, needle1))

	// Example 2
	var haystack2 string = "aaaaa"
	var needle2 string = "bba"
	fmt.Println("Expected: -1 Output: ", strStr(haystack2, needle2))

	// Example 3
	var haystack3 string = ""
	var needle3 string = ""
	fmt.Println("Expected: 0 Output: ", strStr(haystack3, needle3))

	// Example 4
	var haystack4 string = "hello"
	var needle4 string = "low"
	fmt.Println("Expected: -1 Output: ", strStr(haystack4, needle4))

	// Example 5
	var haystack5 string = "a"
	var needle5 string = "a"
	fmt.Println("Expected: 0 Output: ", strStr(haystack5, needle5))

	// Example 6
	var haystack6 string = "mississippi"
	var needle6 string = "issip"
	fmt.Println("Expected: 4 Output: ", strStr(haystack6, needle6))

}
