package main

/*
source: https://leetcode.com/problems/two-sum

problem: Given an array of integers nums and an integer target, return indices of the two numbers such that they add upto target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

// go through all possible pairs and check if their sum is equal to target
func twoSum1(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i != j && v1+v2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// put each value and index in map. check if target - val exists, if so,
// then return current index + index of the found val
func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, v := range nums {
		if _, found := numMap[target-v]; found {
			return []int{i, numMap[target-v]}
		}
		numMap[v] = i
	}

	return nil
}
