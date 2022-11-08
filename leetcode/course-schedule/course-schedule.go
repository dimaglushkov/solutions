package main

import "fmt"

// source: https://leetcode.com/problems/course-schedule/

func canFinish(n int, prerequisites [][]int) bool {
	graph := make([][]int, n)
	for _, p := range prerequisites {
		x, y := p[0], p[1]
		if graph[x] == nil {
			graph[x] = make([]int, 0)
		}
		graph[x] = append(graph[x], y)
	}

	colors := make([]uint8, n)

	var dfs func(x int) bool
	dfs = func(x int) bool {
		if colors[x] == 1 {
			return false
		}
		if colors[x] == 2 {
			return true
		}
		colors[x] = 1
		for _, i := range graph[x] {
			if !dfs(i) {
				return false
			}
		}
		colors[x] = 2
		return true
	}

	for i := range graph {
		if !dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	var numCourses int = 5
	var prerequisites = [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	fmt.Println("Expected: true	Output: ", canFinish(numCourses, prerequisites))

	var numCourses0 int = 2
	var prerequisites0 = [][]int{}
	fmt.Println("Expected: true	Output: ", canFinish(numCourses0, prerequisites0))

	// Example 1
	var numCourses1 int = 2
	//var prerequisites1 = [][]int{{1, 0}}
	var prerequisites1 = [][]int{{0, 1}}
	fmt.Println("Expected: true	Output: ", canFinish(numCourses1, prerequisites1))

	// Example 2
	var numCourses2 int = 2
	var prerequisites2 = [][]int{{1, 0}, {0, 1}}
	fmt.Println("Expected: false	Output: ", canFinish(numCourses2, prerequisites2))

}
