package main

import (
	"fmt"
	"sort"
	"strings"
)

// source: https://leetcode.com/problems/orderly-queue/

func orderlyQueue(s string, k int) string {
	if k == 1 {
		minStr := s
		for i := range s {
			if xs := s[i:] + s[:i]; strings.Compare(xs, minStr) == -1 {
				minStr = xs
			}
		}
		return minStr

	}

	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	// Example 1
	var s1 string = "cba"
	var k1 int = 1
	fmt.Println("Expected: acb	Output: ", orderlyQueue(s1, k1))

	// Example 2
	var s2 string = "baaca"
	var k2 int = 3
	fmt.Println("Expected: aaabc	Output: ", orderlyQueue(s2, k2))

}
