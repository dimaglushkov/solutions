package main

import "fmt"

// source: https://leetcode.com/problems/string-compression-ii/

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if j > 0 {
				dp[i][j] = dp[i-1][j-1]
			}
			var removed, cnt int
			for p := i; p > 0; p-- {
				if s[p-1] == s[i-1] {
					cnt++
				} else {
					removed++
				}
				if removed > j {
					break
				}
				dp[i][j] = min(dp[i][j], dp[p-1][j-removed]+calcLen(cnt))
			}
		}
	}
	return dp[n][k]
}

func calcLen(len int) int {
	switch {
	case len == 0:
		return 0
	case len == 1:
		return 1
	case len < 10:
		return 2
	case len < 100:
		return 3
	default:
		return 4
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	var s1 string = "aaabcccd"
	var k1 int = 2
	fmt.Println("Expected: 4	Output: ", getLengthOfOptimalCompression(s1, k1))

	// Example 2
	var s2 string = "aabbaa"
	var k2 int = 2
	fmt.Println("Expected: 2	Output: ", getLengthOfOptimalCompression(s2, k2))

	// Example 3
	var s3 string = "aaaaaaaaaaa"
	var k3 int = 0
	fmt.Println("Expected: 3	Output: ", getLengthOfOptimalCompression(s3, k3))

}
