package main

import "fmt"

// source: https://leetcode.com/problems/search-in-rotated-sorted-array/

// At first, using binary search to find pivot's position and also checking
// for target - it may occur randomly while searching for pivot
// After pivot id is found, choosing in which part to look for the target
// and using binary search to find the target
func search(nums []int, target int) int {
	var ind = -1
	var offset int
	var l = len(nums)

	if l == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	if l == 2 {
		for i, v := range nums {
			if v == target {
				return i
			}
		}
		return -1
	}

	if nums[0] > nums[l-1] {
		var isTarget bool
		if ind, isTarget = pivotSearch(nums, 0, l, target); isTarget {
			return ind
		}

		if target < nums[0] {
			offset = ind
			l -= offset
		} else {
			l = ind + 1
		}
	}

	return binarySearch(nums, offset, l, target)
}

func pivotSearch(nums []int, offset int, len int, target int) (index int, isTarget bool) {
	midIndex := offset + (len)/2

	if len == 1 || midIndex == 0 {
		if nums[midIndex] == target {
			isTarget = true
		}
		return midIndex, isTarget
	}

	curVal := nums[midIndex]
	prevVal := nums[midIndex-1]

	if curVal == target {
		return midIndex, true
	}
	if prevVal == target {
		return midIndex - 1, true
	}

	if curVal < prevVal {
		return midIndex, false
	}

	if curVal > nums[0] {
		return pivotSearch(nums, midIndex+1, len-(midIndex+1)+offset, target)
	} else {
		return pivotSearch(nums, offset, len-(midIndex+1)+offset, target)
	}
}

func binarySearch(nums []int, offset int, len int, target int) int {
	if len == 1 {
		if nums[offset] == target {
			return offset
		}
		return -1
	}

	midIndex := offset + (len-1)/2
	curVal := nums[midIndex]

	if curVal == target {
		return midIndex
	}

	if curVal < target {
		return binarySearch(nums, midIndex+1, len-(midIndex+1)+offset, target)
	} else {
		return binarySearch(nums, offset, len-(midIndex+1)+offset, target)
	}
}

func main() {
	// Example 12
	var nums12 = []int{4, 6, 1, 2}
	var target12 int = 11
	fmt.Println("Expected: 0	Output: ", search(nums12, target12))

	// Example 11
	var nums11 = []int{4, 5, 6, 7, 0, 1, 2}
	var target11 int = 5
	fmt.Println("Expected: 1	Output: ", search(nums11, target11))

	// Example 10
	var nums10 = []int{1, 3, 5}
	var target10 int = 3
	fmt.Println("Expected: 1	Output: ", search(nums10, target10))

	// Example 1
	var nums1 = []int{4, 5, 6, 7, 0, 1, 2}
	var target1 int = 0
	fmt.Println("Expected: 4	Output: ", search(nums1, target1))

	// Example 2
	var nums2 = []int{4, 5, 6, 7, 0, 1, 2}
	var target2 int = 3
	fmt.Println("Expected: -1	Output: ", search(nums2, target2))

	// Example 3
	var nums3 = []int{1}
	var target3 int = 0
	fmt.Println("Expected: -1	Output: ", search(nums3, target3))

	// Example 4
	var nums4 = []int{0}
	var target4 int = 0
	fmt.Println("Expected: 0	Output: ", search(nums4, target4))

	// Example 5
	var nums5 = []int{2, 1}
	var target5 int = 2
	fmt.Println("Expected: 0	Output: ", search(nums5, target5))

	// Example 6
	var nums6 = []int{2, 1}
	var target6 int = 3
	fmt.Println("Expected: -1	Output: ", search(nums6, target6))

	// Example 7
	var nums7 = []int{3, 5, 1}
	var target7 int = 0
	fmt.Println("Expected: -1	Output: ", search(nums7, target7))

	// Example 9
	var nums9 = []int{4, 5, 6, 7, 0, 1, 2}
	var target9 int = 2
	fmt.Println("Expected: 6	Output: ", search(nums9, target9))

	// Example 8
	var nums8 = []int{5, 1, 3}
	var target8 int = 3
	fmt.Println("Expected: 2	Output: ", search(nums8, target8))
}
