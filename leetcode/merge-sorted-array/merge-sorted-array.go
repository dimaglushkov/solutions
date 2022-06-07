package main

import "fmt"

// source: https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	var i, j int

	var insert = func() {
		nums1 = append(nums1[:i], append([]int{nums2[j]}, nums1[i:]...)...)
		i++
		j++
		m++
	}

	nums1 = nums1[:m]
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			i++
			continue
		}
		insert()
	}

	if j < n {
		nums1 = append(nums1, nums2[j:]...)
	}

	return
}

func main() {
	// Example 0
	var nums10 = []int{4, 4, 9, 0, 0, 0}
	var m0 int = 3
	var nums20 = []int{2, 5, 6}
	var n0 int = 3
	merge(nums10, m0, nums20, n0)
	fmt.Println("Expected: [2,4,4,5,6,9]	Output: ", nums10)

	// Example 1
	var nums11 = []int{1, 2, 3, 0, 0, 0}
	var m1 int = 3
	var nums21 = []int{2, 5, 6}
	var n1 int = 3
	merge(nums11, m1, nums21, n1)
	fmt.Println("Expected: [1,2,2,3,5,6]	Output: ", nums11)

	// Example 2
	var nums12 = []int{1}
	var m2 int = 1
	var nums22 = []int{}
	var n2 int = 0
	merge(nums12, m2, nums22, n2)
	fmt.Println("Expected: [1]	Output: ", nums12)

	// Example 3
	var nums13 = []int{0}
	var m3 int = 0
	var nums23 = []int{1}
	var n3 int = 1
	merge(nums13, m3, nums23, n3)
	fmt.Println("Expected: [1]	Output: ", nums13)

}
