package main

import (
	"fmt"
	"sort"
	"strings"
)

// source: https://leetcode.com/problems/maximum-product-of-word-lengths/

// 1 2 3 4 5 6 7
// 1 * 7 = 7
// 2 * 6 = 12
// 3 * 4 = 12

// 1 2 3 4
// 1 * 4 = 4
// 2 * 3 = 6

func maxProduct_(words []string) int {
	var res int
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	for i := range words {
		for j := range words {
			if !strings.ContainsAny(words[i], words[j]) && len(words[i])*len(words[j]) > res {
				res = len(words[i]) * len(words[j])
			}
		}
	}

	return res
}

func maxProduct(words []string) int {
	best, buf := 0, make([]int, len(words))
	for i, x := range words {
		mask, n := 0, len(x)
		for _, y := range []byte(x) {
			mask |= 1 << (y - 'a')
		}
		buf[i] = mask
		for j, m := range buf[:i] {
			if m&mask != 0 {
				continue
			}
			if now := len(words[j]) * n; best < now {
				best = now
			}
		}
	}
	return best
}

func main() {
	// Example 1
	var words1 = []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println("Expected: 16	Output: ", maxProduct(words1))

	// Example 2
	var words2 = []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	fmt.Println("Expected: 4	Output: ", maxProduct(words2))

	// Example 3
	var words3 = []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println("Expected: 0	Output: ", maxProduct(words3))

}
