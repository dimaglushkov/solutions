package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/3sum-with-multiplicity/
func threeSumMulti(arr []int, target int) int {
	const maxVal = 100
	var valCnt [maxVal + 1]int
	var k, res int

	for _, v := range arr {
		valCnt[v]++
	}
	for i := 0; i <= maxVal; i++ {
		if valCnt[i] == 0 {
			continue
		}
		for j := i; j <= maxVal; j++ {
			k = target - i - j
			if valCnt[j] == 0 || k > 100 || k < 0 || valCnt[k] == 0 {
				continue
			}
			if i == j && j == k {
				res += valCnt[i] * (valCnt[i] - 1) * (valCnt[i] - 2) / 6
			} else if i == j && j != k {
				res += valCnt[i] * (valCnt[i] - 1) / 2 * valCnt[k]
			} else if j < k {
				res += valCnt[i] * valCnt[j] * valCnt[k]
			}
		}

	}

	return res % (1e9 + 7)
}

func main() {
	// Example 1
	var arr1 = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	var target1 int = 8
	fmt.Println("Expected: 20	Output: ", threeSumMulti(arr1, target1))

	// Example 2
	var arr2 = []int{1, 1, 2, 2, 2, 2}
	var target2 int = 5
	fmt.Println("Expected: 12	Output: ", threeSumMulti(arr2, target2))

	// Example 3
	var arr3 = []int{1, 2, 1}
	var target3 int = 5
	fmt.Println("Expected: 0	Output: ", threeSumMulti(arr3, target3))

	// Example 4
	var arr4 = []int{1, 3, 2, 1}
	var target4 int = 6
	fmt.Println("Expected: 2	Output: ", threeSumMulti(arr4, target4))

}
