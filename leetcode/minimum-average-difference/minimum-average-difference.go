package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/minimum-average-difference/

func minimumAverageDifference(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var minDiff = math.MaxInt
	var res int
	var N = len(nums)
	var sums = make([]int, N+1)
	sums[0] = nums[0]
	for i := 1; i < N; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	sums[N] = sums[N-1]

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < N-1; i++ {
		avg1 := sums[i] / (i + 1)
		avg2 := (sums[N-1] - sums[i]) / (N - i - 1)
		tmp := abs(avg1 - avg2)
		if tmp < minDiff {
			minDiff = tmp
			res = i
		}
	}
	if sums[N-1]/N < minDiff {
		res = N - 1
	}
	return res
}
func main() {
	// Example 1
	var nums1 = []int{2, 5, 3, 9, 5, 3}
	fmt.Println("Expected: 3	Output: ", minimumAverageDifference(nums1))

	// Example 2
	var nums2 = []int{0}
	fmt.Println("Expected: 0	Output: ", minimumAverageDifference(nums2))

}
