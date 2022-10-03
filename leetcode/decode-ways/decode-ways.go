package main

import "fmt"

// source: https://leetcode.com/problems/decode-ways/

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		switch s[i] {
		case '0':
			if s[i+1] == '0' {
				return 0
			}
			dp[i] = 0

		case '1':
			dp[i] = dp[i+1] + dp[i+2]

		case '2':
			if s[i+1] < '7' {
				dp[i] = dp[i+1] + dp[i+2]
			} else {
				dp[i] = dp[i+1]
			}

		default:
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}

func main() {
	var s6 string = "1123"
	fmt.Println("Expected: 5	Output: ", numDecodings(s6))

	var s4 string = "2101"
	fmt.Println("Expected: 1	Output: ", numDecodings(s4))

	var s5 string = "10"
	fmt.Println("Expected: 1	Output: ", numDecodings(s5))

	var s0 string = "0000"
	fmt.Println("Expected: 0	Output: ", numDecodings(s0))

	// Example 1
	var s1 string = "12"
	fmt.Println("Expected: 2	Output: ", numDecodings(s1))

	// Example 2
	var s2 string = "226"
	fmt.Println("Expected: 3	Output: ", numDecodings(s2))

	// Example 3
	var s3 string = "06"
	fmt.Println("Expected: 0	Output: ", numDecodings(s3))

}
