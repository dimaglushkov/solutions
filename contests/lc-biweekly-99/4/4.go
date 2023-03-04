package main

import "strconv"

func main() {
	println(rootCount([][]int{{0, 1}, {2, 0}, {0, 3}, {4, 2}, {3, 5}, {6, 0}, {1, 7}, {2, 8}, {2, 9}, {4, 10}, {9, 11}, {3, 12}, {13, 8}, {14, 9}, {15, 9}, {10, 16}}, [][]int{{8, 2}, {12, 3}, {0, 1}, {16, 10}}, 2))
	println(rootCount([][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 0}, {3, 4}, {2, 1}, {3, 2}}, 1))
	println(rootCount([][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}, [][]int{{1, 3}, {0, 1}, {1, 0}, {2, 4}}, 3))
}

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func toString(i, j int) string {
	return itoa(i) + "," + itoa(j)
}

// TLE
func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1
	graph := make(map[int]map[int]bool)
	for _, e := range edges {
		if graph[e[0]] == nil {
			graph[e[0]] = make(map[int]bool)
		}
		graph[e[0]][e[1]] = true

		if graph[e[1]] == nil {
			graph[e[1]] = make(map[int]bool)
		}
		graph[e[1]][e[0]] = true
	}

	rules := make(map[string]bool)
	for _, g := range guesses {
		rules[toString(g[0], g[1])] = true
	}

	var visited []bool
	var cnt int
	ans := 0
	var dfs func(node int)

	dfs = func(node int) {
		visited[node] = true
		for next := range graph[node] {
			if !visited[next] {
				if rules[toString(node, next)] {
					cnt++
				}
				if cnt >= k {
					return
				}
				dfs(next)
			}
		}
	}
	for i := 0; i < n; i++ {
		visited = make([]bool, n)
		cnt = 0
		dfs(i)
		if cnt >= k {
			ans++
		}
	}

	return ans
}
