package main

import "fmt"

// source: https://leetcode.com/problems/find-center-of-star-graph/

func findCenter(edges [][]int) int {
	for i := 0; i < 2; i++ {
		if edges[0][i] == edges[1][0] || edges[0][i] == edges[1][1] {
			return edges[0][i]
		}
	}
	return 0
}

func main() {
	// Example 1
	var edges1 = [][]int{{1, 2}, {2, 3}, {4, 2}}
	fmt.Println("Expected: 2	Output: ", findCenter(edges1))

	// Example 2
	var edges2 = [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}
	fmt.Println("Expected: 1	Output: ", findCenter(edges2))

}
