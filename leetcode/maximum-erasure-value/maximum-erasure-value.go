package main

import "fmt"

// source: https://leetcode.com/problems/maximum-erasure-value/

// first calculating sum of all contiguous subsequences
// then using sliding window technique to find the biggest one with unique numbers
func maximumUniqueSubarray(nums []int) int {
	var N = len(nums)
	var s = make([]int, N+1)
	var sum = 0
	for i, v := range nums {
		s[i] = sum
		sum += v
	}
	s[N] = sum
	nums = append(nums, 0)

	var res int
	var pos = make(map[int]int)
	var l, r = 0, -1

	for r < N {
		r++
		if p, ok := pos[nums[r]]; (ok && p >= l) || r == N {
			if s[r]-s[l] > res {
				res = s[r] - s[l]
			}
			if p == l {
				l++
			} else {
				l = p + 1
			}
		}
		pos[nums[r]] = r
	}

	return res
}

func main() {
	// Example 3
	var nums3 = []int{4, 2, 4, 5, 6, 4, 2, 10}
	fmt.Println("Expected: 27	Output: ", maximumUniqueSubarray(nums3))

	// Example 0
	var nums0 = []int{187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970, 306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434}
	fmt.Println("Expected: 16911	Output: ", maximumUniqueSubarray(nums0))

	// Example 1
	var nums1 = []int{4, 2, 4, 5, 6}
	fmt.Println("Expected: 17	Output: ", maximumUniqueSubarray(nums1))

	// Example 2
	var nums2 = []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
	fmt.Println("Expected: 8	Output: ", maximumUniqueSubarray(nums2))

}
