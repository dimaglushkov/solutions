package main

import "fmt"

// source: https://leetcode.com/problems/is-graph-bipartite/

const (
	black = iota + 1
	white
)

// while going through the graph, painting each next node with one two colors
// every adjacent nodes must have different colors for graph to be bipartite
func isBipartite(graph [][]int) bool {
	if len(graph) < 2 {
		return true
	}
	colors := make([]int, len(graph))
	var dfs func(node, color int) bool
	dfs = func(node, color int) bool {
		if colors[node] != 0 && colors[node] != color {
			return false
		}
		if colors[node] == color {
			return true
		}
		colors[node] = color
		switch color {
		case black:
			color = white
		case white:
			color = black
		}

		for _, sibling := range graph[node] {
			if !dfs(sibling, color) {
				return false
			}
		}

		return true
	}

	for i := range graph {
		if colors[i] == 0 {
			if !dfs(i, black) {
				return false
			}
		}
	}

	return true
}

func main() {
	// Example 3
	var graph3 = [][]int{nil}
	fmt.Println("Expected: true	Output: ", isBipartite(graph3))

	// Example 2
	var graph2 = [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	fmt.Println("Expected: true	Output: ", isBipartite(graph2))

	// Example 1
	var graph1 = [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	fmt.Println("Expected: false	Output: ", isBipartite(graph1))

}
