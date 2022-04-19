package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/word-break/

//TODO: solve with Trie

func wordBreak(s string, wordDict []string) bool {
	return wordBreakUtil(s, wordDict, 0, map[int]bool{})
}

func wordBreakUtil(s string, wordDict []string, i int, dp map[int]bool) bool {
	if val, ok := dp[i]; ok {
		return val
	}
	if i == len(s) {
		dp[i] = true
		return dp[i]
	}
	for _, word := range wordDict {
		if len(s[i:]) < len(word) || s[i:i+len(word)] != word {
			continue
		}
		if wordBreakUtil(s, wordDict, i+len(word), dp) {
			dp[i] = true
			return true
		}
	}
	dp[i] = false
	return dp[i]
}

func main() {
	// Example 4
	var s4 string = "cars"
	var wordDict4 = []string{"car", "ca", "rs"}
	fmt.Println("Expected: true	Output: ", wordBreak(s4, wordDict4))

	// Example 1
	var s1 string = "leetcode"
	var wordDict1 = []string{"leet", "code"}
	fmt.Println("Expected: true	Output: ", wordBreak(s1, wordDict1))

	// Example 2
	var s2 string = "applepenapple"
	var wordDict2 = []string{"apple", "pen"}
	fmt.Println("Expected: true	Output: ", wordBreak(s2, wordDict2))

	// Example 3
	var s3 string = "catsandog"
	var wordDict3 = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println("Expected: false	Output: ", wordBreak(s3, wordDict3))

}
