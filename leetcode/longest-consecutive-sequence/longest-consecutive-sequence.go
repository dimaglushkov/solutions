package main

import "fmt"

// source: https://leetcode.com/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, ok := set[n]; !ok {
			set[n] = true
		}
	}

	var res int
	for num := range set {
		if set[num-1] {
			continue
		}
		length := 1
		for set[num+1] {
			num++
			length++
		}
		if length > res {
			res = length
		}
	}
	return res
}

func main() {
	// Example 1
	var nums1 = []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Expected: 4	Output: ", longestConsecutive(nums1))

	// Example 2
	var nums2 = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println("Expected: 9	Output: ", longestConsecutive(nums2))

}
