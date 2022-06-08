package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/n-queens-ii/

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func totalNQueens(n int) int {
	var sol = make([][]int, 0)
	var diff = make(map[int]bool)
	var sum = make(map[int]bool)
	var queens = make([]int, 0)

	var util func(qs []int, d, s map[int]bool)
	util = func(qs []int, d, s map[int]bool) {
		p := len(qs)
		if p == n {
			tmp := make([]int, n)
			copy(tmp, qs)
			sol = append(sol, tmp)
			return
		}
		for q := 0; q < n; q++ {
			if !contains(qs, q) && !d[p-q] && !s[p+q] {
				d[p-q] = true
				s[p+q] = true
				util(append(qs, q), d, s)
				d[p-q] = false
				s[p+q] = false
			}
		}
	}

	util(queens, diff, sum)

	return len(sol)
}

func main() {
	// Example 1
	var n1 int = 4
	fmt.Println("Expected: 2	Output: ", totalNQueens(n1))

	// Example 2
	var n2 int = 1
	fmt.Println("Expected: 1	Output: ", totalNQueens(n2))

}
