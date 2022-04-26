package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	var res = make([][]int, 0)
	var util func([]int, int, int)
	sort.Ints(candidates)
	util = func(cur []int, sum, i int) {
		if sum == target {
			r := make([]int, len(cur))
			copy(r, cur)
			res = append(res, r)
			return
		}
		if sum+candidates[i] > target {
			return
		}

		for j := i; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			util(cur, sum+candidates[j], j)
			cur = cur[:len(cur)-1]
		}

	}

	util([]int{}, 0, 0)
	return res
}

func main() {
	// Example 1
	var candidates1 = []int{2, 3, 6, 7}
	var target1 int = 7
	fmt.Println("Expected: [[2,2,3],[7]]	Output: ", combinationSum(candidates1, target1))

	// Example 2
	var candidates2 = []int{2, 3, 5}
	var target2 int = 8
	fmt.Println("Expected: [[2,2,2,2],[2,3,3],[3,5]]	Output: ", combinationSum(candidates2, target2))

	// Example 3
	var candidates3 = []int{2}
	var target3 int = 1
	fmt.Println("Expected: []	Output: ", combinationSum(candidates3, target3))

}
