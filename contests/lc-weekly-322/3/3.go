package main

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minScore(n int, roads [][]int) int {
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

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for _, r := range roads {
		union(r[0]-1, r[1]-1)
	}

	t := find(0)
	res := math.MaxInt
	for _, r := range roads {
		if find(r[0]-1) != t || find(r[1]-1) != t {
			continue
		}
		res = min(res, r[2])
	}

	return res
}

func main() {
	println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))
	println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
}
