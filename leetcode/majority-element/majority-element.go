package main

import "fmt"

// source: https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	n := len(nums)
	t := n/2 + 1
	m := make(map[int]int, 0)
	for _, i := range nums {
		m[i]++
		if m[i] == t {
			return i
		}
	}
	return -1
}

func main() {
	// Example 1
	var nums1 = []int{3, 2, 3}
	fmt.Println("Expected: 3	Output: ", majorityElement(nums1))

	// Example 2
	var nums2 = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Expected: 2	Output: ", majorityElement(nums2))

}
