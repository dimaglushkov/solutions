package main

import "fmt"

// source: https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	var res, tmp int
	var l, r = 0, len(height) - 1

	for l < r {
		if height[l] < height[r] {
			tmp = height[l] * (r - l)
			l++
		} else {
			tmp = height[r] * (r - l)
			r--
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func main() {
	// Example 1
	var height1 = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Expected: 49	Output: ", maxArea(height1))

	// Example 2
	var height2 = []int{1, 1}
	fmt.Println("Expected: 1	Output: ", maxArea(height2))

	// Example 3
	var height3 = []int{10, 1000, 10}
	fmt.Println("Expected: 20	Output: ", maxArea(height3))

	// Example 4
	var height4 = []int{0, 0, 0, 1, 0, 0}
	fmt.Println("Expected: 0	Output: ", maxArea(height4))

}
