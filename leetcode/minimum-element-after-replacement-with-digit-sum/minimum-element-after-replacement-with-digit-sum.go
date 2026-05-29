package main

import "math"

// source: https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/
func minElement(nums []int) int {
	ans := math.MaxInt

	for _, i := range nums {
		t := 0
		for i > 0 {
			t += i % 10
			i /= 10
		}
		ans = min(ans, t)
	}

	return ans
}
