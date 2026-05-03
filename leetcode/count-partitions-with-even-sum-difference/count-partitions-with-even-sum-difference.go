package main

// source: https://leetcode.com/problems/count-partitions-with-even-sum-difference/
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countPartitions(nums []int) int {
	ans := 0

	sum := 0
	for _, x := range nums {
		sum += x
	}

	tsum := 0
	for i := 0; i < len(nums)-1; i++ {
		tsum += nums[i]
		if abs(sum-2*tsum)%2 == 0 {
			ans++
		}
	}

	return ans
}
