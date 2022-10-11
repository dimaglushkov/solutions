package main

import "fmt"

// source: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	var res, min int
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		} else {
			if x := prices[i] - prices[min]; x > res {
				res = x
			}
		}
	}

	return res
}

func main() {
	// Example 1
	var prices1 = []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Expected: 5	Output: ", maxProfit(prices1))

	// Example 2
	var prices2 = []int{7, 6, 4, 3, 1}
	fmt.Println("Expected: 0	Output: ", maxProfit(prices2))

}
