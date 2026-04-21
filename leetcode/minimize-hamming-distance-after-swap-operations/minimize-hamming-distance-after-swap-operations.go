package main

import (
	"fmt"
)

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	parent := make([]int, len(source))
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

	for i := range allowedSwaps {
		union(allowedSwaps[i][0], allowedSwaps[i][1])
	}

	s := make(map[int]map[int]int)
	for i := range source {
		s[find(i)] = make(map[int]int)
	}

	for i := range source {
		s[find(i)][source[i]]++
		s[find(i)][target[i]]--
	}

	ans := 0

	for _, v := range s {
		for _, t := range v {
			ans += abs(t)
		}
	}

	return ans / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	testCases := []struct {
		source       []int
		target       []int
		allowedSwaps [][]int
		want         int
	}{
		{
			source:       []int{1, 2, 3, 4},
			target:       []int{2, 1, 4, 5},
			allowedSwaps: [][]int{{0, 1}, {2, 3}},
			want:         1,
		},
		{
			source:       []int{1, 2, 3, 4},
			target:       []int{1, 3, 2, 4},
			allowedSwaps: [][]int{},
			want:         2,
		},
		{
			source:       []int{5, 1, 2, 4, 3},
			target:       []int{1, 5, 4, 2, 3},
			allowedSwaps: [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}},
			want:         0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumHammingDistance(tc.source, tc.target, tc.allowedSwaps)
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
