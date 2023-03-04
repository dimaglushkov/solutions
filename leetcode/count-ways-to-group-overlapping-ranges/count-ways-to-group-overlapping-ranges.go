package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/count-ways-to-group-overlapping-ranges/

const MOD = 1000000000 + 7

func countWays(ranges [][]int) int {
	n := len(ranges)
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

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[i][1]
		}
		return ranges[i][0] < ranges[j][0]
	})

	active := 0
	for i := 1; i < n; i++ {
		if active != -1 {
			if ranges[active][1] >= ranges[i][0] {
				union(active, i)
				if ranges[i][1] > ranges[active][1] {
					active = i
				}
			} else {
				active = -1
			}
		}
		if ranges[i][0] <= ranges[i-1][1] {
			union(i-1, i)
		}
		if active == -1 {
			active = i
		}
	}

	groups := make(map[int]bool)
	for i := range parent {
		groups[find(i)] = true
	}
	size := len(groups)
	var ans uint64 = 1
	for i := 0; i < size; i++ {
		ans = ans * 2 % MOD
	}
	return int(ans)
}

func main() {
	testCases := []struct {
		ranges [][]int
		want   int
	}{
		{
			ranges: [][]int{{6, 10}, {5, 15}},
			want:   2,
		},
		{
			ranges: [][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}},
			want:   4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countWays(tc.ranges)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
