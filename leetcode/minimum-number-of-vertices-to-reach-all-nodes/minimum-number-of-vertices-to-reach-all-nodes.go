package main

import "fmt"

// source: https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	out := make([]bool, n)
	for i := 0; i < len(edges); i++ {
		out[edges[i][1]] = true
	}
	var res = []int{}
	for i := 0; i < n; i++ {
		if !out[i] {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	// Example 1
	var n1 int = 6
	var edges1 = [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	fmt.Println("Expected: [0,3]	Output: ", findSmallestSetOfVertices(n1, edges1))

	// Example 2
	var n2 int = 5
	var edges2 = [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}
	fmt.Println("Expected: [0,2,3]	Output: ", findSmallestSetOfVertices(n2, edges2))

	// Example 3
	var n3 int = 4
	var edges3 = [][]int{{1, 2}, {3, 2}, {1, 3}, {1, 0}, {0, 2}, {0, 3}}
	fmt.Println("Expected: [1]	Output: ", findSmallestSetOfVertices(n3, edges3))

}
