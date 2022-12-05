package main

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func magnificentSets(n int, edges [][]int) int {
	type edge struct {
		id, level int
	}
	queue := make([]edge, 0)
	push := func(x edge) {
		queue = append(queue, x)
	}
	pop := func() edge {
		x := queue[0]
		queue = queue[1:]
		return x
	}

	parent := make([]int, n)
	var find func(int) int
	var union func(int, int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for i := range parent {
		parent[i] = i
	}

	for _, e := range edges {
		union(e[0]-1, e[1]-1)
	}

	graphs := make(map[int]map[int][]int, 0)
	for _, e := range edges {
		p := find(e[0] - 1)
		if graphs[p] == nil {
			graphs[p] = make(map[int][]int)
		}
		graphs[p][e[0]-1] = append(graphs[p][e[0]-1], e[1]-1)
		graphs[p][e[1]-1] = append(graphs[p][e[1]-1], e[0]-1)
	}

	res := 0

	for _, graph := range graphs {
		graphRes := -1
		starts, minCnt := make([]int, 0), math.MaxInt
		for i := range graph {
			c := len(graph[i])
			if c == minCnt {
				starts = append(starts, i)
			}
			if c < minCnt {
				minCnt = c
				starts = []int{i}
			}

		}

		levels := make([]int, n)
	loop:
		for i := range starts {
			for j := range levels {
				levels[j] = -1
			}

			m := 0
			push(edge{starts[i], 0})
			for len(queue) > 0 {
				x := pop()
				v, l := x.id, x.level
				levels[v] = l

				for _, k := range graph[v] {
					if levels[k] == -1 {
						push(edge{k, l + 1})
					} else if levels[k] != l+1 && levels[k] != l-1 {
						continue loop
					}
				}

				if l+1 > m {
					m = l + 1
				}
			}

			graphRes = max(graphRes, m)
		}
		if graphRes == -1 {
			return -1
		}
		res += graphRes
	}

	for i := 0; i < n; i++ {
		if graphs[find(i)] == nil {
			res++
		}
	}

	return res
}

func main() {
	println(magnificentSets(92, [][]int{{67, 29}, {13, 29}, {77, 29}, {36, 29}, {82, 29}, {54, 29}, {57, 29}, {53, 29}, {68, 29}, {26, 29}, {21, 29}, {46, 29}, {41, 29}, {45, 29}, {56, 29}, {88, 29}, {2, 29}, {7, 29}, {5, 29}, {16, 29}, {37, 29}, {50, 29}, {79, 29}, {91, 29}, {48, 29}, {87, 29}, {25, 29}, {80, 29}, {71, 29}, {9, 29}, {78, 29}, {33, 29}, {4, 29}, {44, 29}, {72, 29}, {65, 29}, {61, 29}}))
	println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	println(magnificentSets(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}
