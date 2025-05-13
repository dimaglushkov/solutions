package main

// source: https://leetcode.com/problems/total-characters-in-string-after-transformations-i/
const mod = 1e9 + 7

func lengthAfterTransformations(s string, t int) int {
	dp := make([]int, 26)
	for _, c := range s {
		dp[c-'a']++
	}

	for i := 0; i < t; i++ {
		next := make([]int, 26)
		next[0] = dp[25]
		next[1] = (dp[25] + dp[0]) % mod
		for j := 2; j < 26; j++ {
			next[j] = dp[j-1]
		}

		dp = next
	}

	ans := 0
	for i := 0; i < 26; i++ {
		ans = (ans + dp[i]) % mod
	}

	return ans
}
