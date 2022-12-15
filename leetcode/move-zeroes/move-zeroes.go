package main

// source: https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int)  {
	ni := 0
	for i := range nums {
		if nums[i] == 0 {
		} else{
			nums[ni] = nums[i]
			ni++
		}
	}
	for ni < len(nums) {
		nums[ni] = 0
		ni++
	}
}

