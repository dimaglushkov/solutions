package main

import "fmt"

// source: https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = 1
	}
	pref, post := 1, 1

	for i := range res {
		res[i] *= pref
		pref *= nums[i]

		res[n-1-i] *= post
		post *= nums[n-1-i]
	}
	return res
}

// this also works, but the above solutions is memory optimized
func _productExceptSelf(nums []int) []int {
	n := len(nums)
	pref := make([]int, n)
	post := make([]int, n)

	pref[0] = nums[0]
	post[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] * nums[i]
		post[n-1-i] = post[n-i] * nums[n-1-i]
	}

	res := make([]int, n)
	for i := range res {
		res[i] = 1
		if i > 0 {
			res[i] *= pref[i-1]
		}
		if i < n-1 {
			res[i] *= post[i+1]
		}
	}

	return res
}

func main() {
	var nums = []int{0, 0}
	fmt.Println("Expected: [0, 0]	Output: ", productExceptSelf(nums))

	// Example 1
	var nums1 = []int{1, 2, 3, 4}
	fmt.Println("Expected: [24,12,8,6]	Output: ", productExceptSelf(nums1))

	// Example 2
	var nums2 = []int{-1, 1, 0, -3, 3}
	fmt.Println("Expected: [0,0,9,0,0]	Output: ", productExceptSelf(nums2))

}
