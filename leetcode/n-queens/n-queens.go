package main

import (
	"bytes"
	"fmt"
)

// source: https://leetcode.com/problems/n-queens/

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func solveNQueens(n int) [][]string {
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

	var res = make([][]string, len(sol))
	var emptyStrBytes = bytes.Repeat([]byte("."), n)
	for si, sv := range sol {
		res[si] = make([]string, n)
		for i := 0; i < n; i++ {
			emptyStrBytes[sv[i]] = byte('Q')
			res[si][i] = string(emptyStrBytes)
			emptyStrBytes[sv[i]] = byte('.')
		}
	}
	return res
}

func main() {
	solveNQueens(7)

	// Example 1
	var n1 int = 4
	fmt.Println("Expected: [[\".Q..\",\"...Q\",\"Q...\",\"..Q.\"],[\"..Q.\",\"Q...\",\"...Q\",\".Q..\"]]	Output: ", solveNQueens(n1))

	// Example 2
	var n2 int = 1
	fmt.Println("Expected: [[\"Q\"]]	Output: ", solveNQueens(n2))

}
