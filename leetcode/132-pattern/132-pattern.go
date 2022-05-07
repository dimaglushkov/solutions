package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/132-pattern/

type stack struct {
	values [][2]int
}

func newStack() stack {
	return stack{}
}

func (s *stack) empty() bool {
	return len(s.values) == 0
}

func (s *stack) push(val, curMin int) {
	s.values = append(s.values, [2]int{val, curMin})
}

func (s *stack) pop() [2]int {
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val
}

func (s *stack) top() [2]int {
	return s.values[len(s.values)-1]
}

func find132pattern(nums []int) bool {
	var s = newStack()
	var curMin = nums[0]

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := range nums {
		for !s.empty() && s.top()[0] <= nums[i] {
			s.pop()
		}
		if !s.empty() && s.top()[1] < nums[i] {
			return true
		}
		s.push(nums[i], curMin)
		curMin = min(nums[i], curMin)
	}

	return false
}

func main() {
	// Example 4
	var nums4 = []int{-2, 1, 1}
	fmt.Println("Expected: false	Output: ", find132pattern(nums4))

	// Example 1
	var nums1 = []int{1, 2, 3, 4}
	fmt.Println("Expected: false	Output: ", find132pattern(nums1))

	// Example 2
	var nums2 = []int{3, 1, 4, 2}
	fmt.Println("Expected: true	Output: ", find132pattern(nums2))

	// Example 3
	var nums3 = []int{-1, 3, 2, 0}
	fmt.Println("Expected: true	Output: ", find132pattern(nums3))

}
