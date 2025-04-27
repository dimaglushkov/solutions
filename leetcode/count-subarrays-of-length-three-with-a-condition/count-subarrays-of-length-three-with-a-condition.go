package main

// source: https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
func countSubarrays(nums []int) int {
	ans := 0

	for i := 0; i+2 < len(nums); i++ {
		if float64(nums[i]+nums[i+2]) == float64(nums[i+1])/2 {
			ans++
		}
	}

	return ans
}
