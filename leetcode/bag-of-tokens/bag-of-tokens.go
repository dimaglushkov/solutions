package main

import (
	"sort"
)

// source: https://leetcode.com/problems/bag-of-tokens/
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	l, r := 0, len(tokens)-1
	ans := 0

	for l <= r && tokens[l] <= power {
		power -= tokens[l]
		ans++
		l++
	}

	for l < r && ans > 0 {
		ans--
		power += tokens[r]
		for l < r && tokens[l] <= power {
			power -= tokens[l]
			l++
			ans++
		}

		r--
	}

	return ans
}
