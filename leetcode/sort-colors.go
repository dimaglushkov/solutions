package main

import "fmt"

// source: https://leetcode.com/problems/sort-colors/

// Since nums[i] is either 0, 1, or 2, we can basically count them
// and then insert proper amount of each digit
// But I guess this could be solved faster with quickSort
func sortColors(nums []int) {
	var i int
	counts := make([]int, 3)
	for _, v := range nums {
		counts[v]++
	}
	for n, c := range counts {
		for j := 0; j < c; j++ {
			nums[i] = n
			i++
		}
	}
}

func main() {
	// Example 1
	var nums1 = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	fmt.Println("Expected: [0,0,1,1,2,2]	Output: ", nums1)

	// Example 2
	var nums2 = []int{2, 0, 1}
	sortColors(nums2)
	fmt.Println("Expected: [0,1,2]	Output: ", nums2)

}
