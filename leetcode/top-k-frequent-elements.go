package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	counters := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counters[nums[i]]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range counters {
		if buckets[freq] == nil {
			buckets[freq] = []int{}
		}
		buckets[freq] = append(buckets[freq], num)
	}

	res := make([]int, k)
	resCounter := 0
	for i := len(buckets) - 1; i >= 0; i-- {
		if buckets[i] == nil {
			continue
		}
		for _, val := range buckets[i] {
			res[resCounter] = val
			resCounter++
		}
		if resCounter == k {
			break
		}
	}
	return res
}

func main() {
	// Example 3
	var nums3 = []int{1, 1, 1, 2, 2, 3, 1, 2, 3, 4, 1, 2}
	var k3 int = 3
	fmt.Println("Expected: [1,2,3]	Output: ", topKFrequent(nums3, k3))

	// Example 1
	var nums1 = []int{1, 1, 1, 2, 2, 3}
	var k1 int = 2
	fmt.Println("Expected: [1,2]	Output: ", topKFrequent(nums1, k1))

	// Example 2
	var nums2 = []int{1}
	var k2 int = 1
	fmt.Println("Expected: [1]	Output: ", topKFrequent(nums2, k2))

}
