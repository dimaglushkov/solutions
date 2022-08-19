package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/reduce-array-size-to-the-half/

func minSetSize(arr []int) int {
	cnt := make([]uint32, 100001)
	freq := make([]uint32, 100001)
	for _, num := range arr {
		cnt[num]++
		freq[cnt[num]]++
		freq[cnt[num]-1]--
	}
	var nremoved int
	var actions int
	for count := 100000; count >= 1; count-- {
		for freq[count] > 0 {
			nremoved += count
			actions += 1
			if nremoved >= len(arr)/2 {
				return actions
			}
			freq[count]--
		}
	}
	return -1 // cannot happen
}

func main() {
	// Example 1
	var arr1 = []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
	fmt.Println("Expected: 2	Output: ", minSetSize(arr1))

	// Example 2
	var arr2 = []int{7, 7, 7, 7, 7, 7}
	fmt.Println("Expected: 1	Output: ", minSetSize(arr2))

}
