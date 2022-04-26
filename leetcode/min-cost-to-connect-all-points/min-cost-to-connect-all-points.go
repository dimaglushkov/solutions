package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/min-cost-to-connect-all-points/
const max = math.MaxInt

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Optimized Prim's algorithm
func minCostConnectPoints(points [][]int) int {
	l := len(points)
	visited := make([]bool, l)

	dist := make([]int, l)
	for i := 0; i < l; i++ {
		dist[i] = max
	}

	var v, res, minDist int
	dist[v] = 0

	for minDist != max {
		visited[v] = true
		res += dist[v]

		for i := 0; i < l; i++ {
			if i == v || visited[i] {
				continue
			}
			if d := abs(points[v][0]-points[i][0]) + abs(points[v][1]-points[i][1]); d < dist[i] {
				dist[i] = d
			}
		}

		minDist = max
		for i := 0; i < l; i++ {
			if !visited[i] && (dist[i] < minDist) {
				minDist, v = dist[i], i
			}
		}
	}

	return res
}

// First, for given set of points build a graph, then
// use Prim's algorithm to form min spanning tree
// return sum of edges in this tree
func minCostConnectPoints_(points [][]int) (res int) {
	l := len(points)
	if l == 1 {
		return 0
	}

	var verCnt, ver, min int
	var visited = make([]bool, l)

	edges := make([][]int, l)
	for i := 0; i < l; i++ {
		edges[i] = make([]int, l)
		for j := 0; j < l; j++ {
			if i == j {
				edges[i][j] = max
			}
			edges[i][j] = abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
		}
	}

	visited[0] = true
	for verCnt < l-1 {
		min = max
		for i := 0; i < l; i++ {
			if visited[i] {
				for j := 0; j < l; j++ {
					if !visited[j] && edges[i][j] < min {
						min = edges[i][j]
						ver = j
					}
				}
			}
		}
		res += min
		visited[ver] = true
		verCnt++
	}

	return res
}

func main() {
	// Example 5
	var points5 = [][]int{{2, -3}, {-17, -8}, {13, 8}, {-17, -15}}
	fmt.Println("Expected: 53	Output: ", minCostConnectPoints(points5))

	// Example 6
	var points6 = [][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}
	fmt.Println("Expected: 4	Output: ", minCostConnectPoints(points6))

	// Example 1
	var points1 = [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println("Expected: 20	Output: ", minCostConnectPoints(points1))

	// Example 2
	var points2 = [][]int{{3, 12}, {-2, 5}, {-4, 1}}
	fmt.Println("Expected: 18	Output: ", minCostConnectPoints(points2))

	// Example 3
	var points3 = [][]int{{3, 13}}
	fmt.Println("Expected: 0	Output: ", minCostConnectPoints(points3))

	// Example 4
	var points4 = [][]int{{0, 0}, {1, 1}}
	fmt.Println("Expected: 2	Output: ", minCostConnectPoints(points4))

}
