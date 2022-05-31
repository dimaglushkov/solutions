package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/

// using rolling hash technique
func hasAllCodes(s string, k int) bool {
	var cnt = 0
	var size = 1 << k
	var ones = 1<<k - 1
	var set = make(map[int]bool)
	var hash int
	for i, b := range s {
		hash = ((hash << 1) & ones) | int(b-'0')
		if i >= k-1 && !set[hash] {
			set[hash] = true
			cnt++
		}
	}

	return cnt == size
}

func main() {
	// Example 4
	var s4 string = "110101"
	var k4 int = 3
	fmt.Println("Expected: false	Output: ", hasAllCodes(s4, k4))

	// Example 1
	var s1 string = "00110110"
	var k1 int = 2
	fmt.Println("Expected: true	Output: ", hasAllCodes(s1, k1))

	// Example 2
	var s2 string = "0110"
	var k2 int = 1
	fmt.Println("Expected: true	Output: ", hasAllCodes(s2, k2))

	// Example 3
	var s3 string = "0110"
	var k3 int = 2
	fmt.Println("Expected: false	Output: ", hasAllCodes(s3, k3))
}
