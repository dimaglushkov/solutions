package main

import "fmt"

// source: https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r, m := matrix[0][0], matrix[n-1][n-1], 0
	count := 0
	for l < r {
		m = l + (r-l)/2
		count = countLessEqual(matrix, m)
		if count < k {
			l = m + 1
			continue
		}
		r = m
	}
	return l
}

func countLessEqual(matrix [][]int, mid int) int {
	n := len(matrix)
	cnt := 0
	for r, c := n-1, 0; r >= 0 && c < n; {
		if matrix[r][c] > mid {
			r--
			continue
		}
		cnt += r + 1
		c++
	}
	return cnt
}

func main() {
	// Example 1
	var matrix1 = [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	var k1 int = 8
	fmt.Println("Expected: 13	Output: ", kthSmallest(matrix1, k1))

	// Example 2
	var matrix2 = [][]int{{-5}}
	var k2 int = 1
	fmt.Println("Expected: -5	Output: ", kthSmallest(matrix2, k2))

}
