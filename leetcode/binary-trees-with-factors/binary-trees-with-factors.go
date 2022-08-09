package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/binary-trees-with-factors/

func numFactoredBinaryTrees(arr []int) int {
	const mod = 1_000_000_007
	n := len(arr)
	sort.Ints(arr)

	dp := make([]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[arr[i]] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				r := arr[i] / arr[j]
				if _, ok := m[r]; ok {
					dp[i] = (dp[i] + dp[j]*dp[m[r]]) % mod
				}
			}
		}
	}

	var res int64
	for _, x := range dp {
		res += x
	}
	return int(res % mod)
}

func main() {
	// Example 1
	var arr1 = []int{2, 4}
	fmt.Println("Expected: 3	Output: ", numFactoredBinaryTrees(arr1))

	// Example 2
	var arr2 = []int{2, 4, 5, 10}
	fmt.Println("Expected: 7	Output: ", numFactoredBinaryTrees(arr2))

}
