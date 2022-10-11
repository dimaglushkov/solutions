package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/increasing-triplet-subsequence/

// straight-forward greedy solution, TLE
/*func increasingTriplet(nums []int) bool {
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if nums[i] < nums[j] {
				for k := j + 1; k < n; k++ {
					if nums[j] < nums[k] {
						return true
					}
				}
			}
		}
	}
	return false
}*/

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	s, ss := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		switch x := nums[i]; {
		case x < s:
			s = x
		case x > s && x < ss:
			ss = x
		case x > ss:
			return true
		}
	}
	return false
}

func main() {
	numsb := []int{1, 5, 0, 4, 1, 3}
	fmt.Println("Expected: false\tOutput: ", increasingTriplet(numsb))

	numsa := []int{1, 2, 2, 1}
	fmt.Println("Expected: false\tOutput: ", increasingTriplet(numsa))

	nums := []int{0, 4, 2, 1, 0, -1, -3}
	fmt.Println("Expected: false\tOutput: ", increasingTriplet(nums))

	// Example 1
	var nums1 = []int{1, 2, 3, 4, 5}
	fmt.Println("Expected: true	Output: ", increasingTriplet(nums1))

	// Example 2
	var nums2 = []int{5, 4, 3, 2, 1}
	fmt.Println("Expected: false	Output: ", increasingTriplet(nums2))

	// Example 3
	var nums3 = []int{2, 1, 5, 0, 4, 6}
	fmt.Println("Expected: true	Output: ", increasingTriplet(nums3))

}
