package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/evaluate-division/

// this is the worst problem I've solved so far thus I bet the
// solution is far from being optimal

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var graph = make(map[string]map[string]float64)
	var addEdge func(x, y string, val float64)
	addEdge = func(x, y string, val float64) {
		if _, ok := graph[x]; !ok {
			graph[x] = make(map[string]float64)
			graph[x][x] = 1
		}
		graph[x][y] = val
		if _, ok := graph[y]; !ok {
			graph[y] = make(map[string]float64)
			graph[y][y] = 1
		}
		graph[y][x] = 1 / val
		for i := range graph {
			if graph[i][y] == 0 && graph[i][x] != 0 && i != y {
				addEdge(i, y, val*graph[i][x])
			}
			if graph[i][x] == 0 && graph[i][y] != 0 && i != x {
				addEdge(i, x, graph[i][y]/val)
			}
		}
	}

	// build the graph
	for i, eq := range equations {
		addEdge(eq[0], eq[1], values[i])
	}

	// getting queries values from the graph
	res := make([]float64, len(queries))
	for i, q := range queries {
		var x, y = q[0], q[1]

		if graph[x][y] == 0 {
			res[i] = -1
			continue
		}
		res[i] = graph[x][y]
	}
	return res
}

func main() {
	// Example 6
	var equations6 = [][]string{{"a", "e"}, {"b", "e"}}
	var values6 = []float64{4.0, 3.0}
	var queries6 = [][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}}
	fmt.Println("Expected: [1.33333,1.00000,-1.00000]	Output: ", calcEquation(equations6, values6, queries6))

	// Example 5
	var equations5 = [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}
	var values5 = []float64{3, 4, 5, 6}
	var queries5 = [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}
	fmt.Println("Expected: [360.00000,0.00833,20.00000,1.00000,-1.00000,-1.00000]	Output: ", calcEquation(equations5, values5, queries5))

	// Example 4
	var equations4 = [][]string{{"a", "aa"}}
	var values4 = []float64{9}
	var queries4 = [][]string{{"aa", "a"}, {"aa", "aa"}}
	fmt.Println("Expected: [0.1111,1]	Output: ", calcEquation(equations4, values4, queries4))

	// Example 2
	var equations2 = [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	var values2 = []float64{1.5, 2.5, 5.0}
	var queries2 = [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	fmt.Println("Expected: [3.75000,0.40000,5.00000,0.20000]	Output: ", calcEquation(equations2, values2, queries2))

	// Example 1
	var equations1 = [][]string{{"a", "b"}, {"b", "c"}}
	var values1 = []float64{2.0, 3.0}
	var queries1 = [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println("Expected: [6.00000,0.50000,-1.00000,1.00000,-1.00000]	Output: ", calcEquation(equations1, values1, queries1))

	// Example 3
	var equations3 = [][]string{{"a", "b"}}
	var values3 = []float64{0.5}
	var queries3 = [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}
	fmt.Println("Expected: [0.50000,2.00000,-1.00000,-1.00000]	Output: ", calcEquation(equations3, values3, queries3))

}
