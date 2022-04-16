package main

import "fmt"

// source: https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	var tmp, res int
	for i := 0; i < len(accounts); i++ {
		tmp = 0
		for j := 0; j < len(accounts[i]); j++ {
			tmp += accounts[i][j]
		}
		if tmp > res {
			res = tmp

		}
	}
	return res
}

func main() {
	// Example 1 
	var accounts1 = [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println("Expected: 6	Output: ", maximumWealth(accounts1))

	// Example 2 
	var accounts2 = [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println("Expected: 10	Output: ", maximumWealth(accounts2))

	// Example 3 
	var accounts3 = [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println("Expected: 17	Output: ", maximumWealth(accounts3))

}
