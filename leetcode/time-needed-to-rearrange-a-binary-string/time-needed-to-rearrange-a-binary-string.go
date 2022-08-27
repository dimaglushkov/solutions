package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/time-needed-to-rearrange-a-binary-string/

func secondsToRemoveOccurrences(s string) int {
	var cnt int

	for {
		if !strings.Contains(s, "01") {
			break
		}
		cnt++
		s = strings.ReplaceAll(s, "01", "10")
	}

	return cnt
}

func main() {
	// Example 1
	var s1 string = "0110101"
	fmt.Println("Expected: 4	Output: ", secondsToRemoveOccurrences(s1))

	// Example 2
	var s2 string = "11100"
	fmt.Println("Expected: 0	Output: ", secondsToRemoveOccurrences(s2))

}
