package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/

func numRollsToTarget(n int, k int, target int) int {
	const mod = 1000000007
	var memo = make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var util func(x, y int) int
	util = func(n, target int) int {
		if n < 0 || target < 0 {
			return 0
		}
		if n == 0 && target == 0 {
			return 1
		}
		if memo[n][target] != -1 {
			return memo[n][target]
		}
		var sum = 0
		for i := 1; i <= k; i++ {
			if target-i < 0 {
				break
			}
			sum += util(n-1, target-i)
		}
		memo[n][target] = sum % mod
		return memo[n][target]
	}

	return util(n, target)
}

func main() {
	//Example 5
	var n5 int = 10
	var k5 int = 5
	var target5 int = 180
	fmt.Println("Expected: 1	Output: ", numRollsToTarget(n5, k5, target5))

	//Example 1
	var n1 int = 1
	var k1 int = 6
	var target1 int = 3
	fmt.Println("Expected: 1	Output: ", numRollsToTarget(n1, k1, target1))

	// Example 2
	var n2 int = 2
	var k2 int = 6
	var target2 int = 7
	fmt.Println("Expected: 6	Output: ", numRollsToTarget(n2, k2, target2))

	// Example 3
	var n3 int = 30
	var k3 int = 30
	var target3 int = 500
	fmt.Println("Expected: 222616187	Output: ", numRollsToTarget(n3, k3, target3))

}
