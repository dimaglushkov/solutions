package main

import (
	"fmt"
	"math/rand"
)

// source: https://leetcode.com/problems/sort-an-array/

func sortArray(data []int) []int {
	var doSort func(data []int, l, r int)
	doSort = func(data []int, l, r int) {
		if l < r {
			p := randomPartition(data, l, r)
			doSort(data, l, p-1)
			doSort(data, p+1, r)
		}
	}
	doSort(data, 0, len(data)-1)
	return data
}

// randomPartition implements randomized partition algorithm by randomly selecting pivot
func randomPartition(data []int, l, r int) int {
	x := rand.Intn(r-l) + l
	data[x], data[r] = data[r], data[x]

	i, pivot := l-1, data[r]
	for j := l; j < r; j++ {
		if data[j] <= pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[r] = data[r], data[i+1]
	return i + 1
}

func main() {
	// Example 1
	var nums1 []int = []int{5, 2, 3, 1}
	fmt.Println("Expected: [1,2,3,5]	Output: ", sortArray(nums1))

	// Example 2
	var nums2 []int = []int{5, 1, 1, 2, 0, 0}
	fmt.Println("Expected: [0,0,1,1,2,5]	Output: ", sortArray(nums2))

}
