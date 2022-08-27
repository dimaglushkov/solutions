package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/split-array-into-consecutive-subsequences/

func isPossible(nums []int) bool {
	prevNum := math.MinInt32
	state := [3]int{}

	for len(nums) > 0 {
		num := nums[0]
		if prevNum+1 != num {
			if state[1] != 0 || state[2] != 0 {
				return false
			}
			state[0] = 0
		}

		count := 0
		for len(nums) > 0 && nums[0] == num {
			count++
			nums = nums[1:]
		}

		count -= state[1] + state[2]
		if count < 0 {
			return false
		}

		state[0], state[1], state[2] = state[1]+min(count, state[0]), state[2], max(0, count-state[0])

		prevNum = num
	}

	return state[1] == 0 && state[2] == 0
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	// Example 1
	var nums1 = []int{1, 2, 3, 3, 4, 5}
	fmt.Println("Expected: true	Output: ", isPossible(nums1))

	// Example 2
	var nums2 = []int{1, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println("Expected: true	Output: ", isPossible(nums2))

	// Example 3
	var nums3 = []int{1, 2, 3, 4, 4, 5}
	fmt.Println("Expected: false	Output: ", isPossible(nums3))

}
