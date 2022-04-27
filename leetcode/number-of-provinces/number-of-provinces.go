package main

import "fmt"

// source: https://leetcode.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
	var res int = len(isConnected)
	parent := make([]int, len(isConnected))
	var find func(int) int
	var union func(int, int)

	union = func(x, y int) {
		fx, fy := find(x), find(y)
		parent[fx] = fy
		if fx != fy {
			res--
		}
	}
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	N := len(isConnected)
	for i := 0; i < N; i++ {
		parent[i] = i
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	return res
}

func main() {
	// Example 1
	var isConnected1 = [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println("Expected: 2	Output: ", findCircleNum(isConnected1))

	// Example 2
	var isConnected2 = [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println("Expected: 3	Output: ", findCircleNum(isConnected2))

}
