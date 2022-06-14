package main

import "fmt"

// source: https://leetcode.com/problems/delete-operation-for-two-strings/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(w1 string, w2 string) int {
	var dp = make([][]int, len(w1)+1)
	for i := range dp {
		dp[i] = make([]int, len(w2)+1)
	}

	for i := 0; i <= len(w1); i++ {
		for j := 0; j <= len(w2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = i + j
			} else if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(w1)][len(w2)]
}

func main() {
	// Example 1
	var word11 string = "sea"
	var word21 string = "eat"
	fmt.Println("Expected: 2	Output: ", minDistance(word11, word21))

	// Example 2
	var word12 string = "leetcode"
	var word22 string = "etco"
	fmt.Println("Expected: 4	Output: ", minDistance(word12, word22))

}
