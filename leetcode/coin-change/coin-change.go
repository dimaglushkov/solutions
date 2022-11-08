package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/coin-change/
func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = 10001
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] >= 10001 {
		return -1
	}
	return dp[amount]
}

func main() {
	// Example 1
	var coins1 = []int{1, 2, 5}
	var amount1 int = 11
	fmt.Println("Expected: 3	Output: ", coinChange(coins1, amount1))

	// Example 2
	var coins2 = []int{2}
	var amount2 int = 3
	fmt.Println("Expected: -1	Output: ", coinChange(coins2, amount2))

	// Example 3
	var coins3 = []int{1}
	var amount3 int = 0
	fmt.Println("Expected: 0	Output: ", coinChange(coins3, amount3))

}
