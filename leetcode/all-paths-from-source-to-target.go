package main

import "fmt"

// source: https://leetcode.com/problems/all-paths-from-source-to-target/
func allPathsSourceTarget(graph [][]int) [][]int {
	pathsCache := make([][][]int, len(graph))
	return findPaths(graph, 0, len(graph)-1, pathsCache)
}

func findPaths(graph [][]int, start, end int, pathsCache [][][]int) (res [][]int) {
	if start == end {
		return [][]int{{start}}
	}

	res = make([][]int, 0)
	for _, node := range graph[start] {
		var nodeToEndPaths [][]int
		if nodeToEndPaths = pathsCache[node]; nodeToEndPaths == nil {
			nodeToEndPaths = findPaths(graph, node, end, pathsCache)
			pathsCache[node] = nodeToEndPaths
		}
		for _, path := range nodeToEndPaths {
			res = append(res, append([]int{start}, path...))
		}
	}
	return res
}

func main() {
	// Example 1
	var graph1 = [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println("Expected: [[0,1,3],[0,2,3]]	Output: ", allPathsSourceTarget(graph1))

	// Example 2
	var graph2 = [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println("Expected: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]	Output: ", allPathsSourceTarget(graph2))

}
