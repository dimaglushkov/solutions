package main

import "fmt"

// source: https://leetcode.com/problems/combination-sum-iii/

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}

	var dfs func(sum, pos int, values []int)
	dfs = func(sum, pos int, values []int) {
		if sum == n && len(values) == k {
			res = append(res, append([]int{}, values...))
			return
		}
		if sum > n || len(values) > k {
			return
		}
		for i := pos; i < 10; i++ {
			dfs(sum+i, i+1, append(values, i))
		}
	}
	dfs(0, 1, []int{})

	return res
}

func main() {
	// Example 1
	var k1 int = 3
	var n1 int = 7
	fmt.Println("Expected: [[1,2,4]]	Output: ", combinationSum3(k1, n1))

	// Example 2
	var k2 int = 3
	var n2 int = 9
	fmt.Println("Expected: [[1,2,6],[1,3,5],[2,3,4]]	Output: ", combinationSum3(k2, n2))

	// Example 3
	var k3 int = 4
	var n3 int = 1
	fmt.Println("Expected: []	Output: ", combinationSum3(k3, n3))

}
