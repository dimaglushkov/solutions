package main

import "fmt"

// source: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	var l, r = 0, len(numbers) - 1
	var sum int
	for l < r {
		sum = numbers[l] + numbers[r]
		if sum == target {
			break
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}

	return []int{l + 1, r + 1}
}

func main() {
	// Example 1
	var numbers1 = []int{2, 7, 11, 15}
	var target1 int = 9
	fmt.Println("Expected: [1,2]	Output: ", twoSum(numbers1, target1))

	// Example 2
	var numbers2 = []int{2, 3, 4}
	var target2 int = 6
	fmt.Println("Expected: [1,3]	Output: ", twoSum(numbers2, target2))

	// Example 3
	var numbers3 = []int{-1, 0}
	var target3 int = -1
	fmt.Println("Expected: [1,2]	Output: ", twoSum(numbers3, target3))

}
