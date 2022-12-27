package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/longest-subsequence-with-limited-sum/

func answerQueries(nums []int, queries []int) []int {
	n, m := len(nums), len(queries)
	res := make([]int, m)
	ps := make([]int, n)
	sum := 0
	sort.Ints(nums)
	for i := range nums {
		sum += nums[i]
		ps[i] = sum
	}
	for i := range res {
		res[i] = sort.SearchInts(ps, queries[i])

		if res[i] < n && ps[res[i]] == queries[i] {
			res[i]++
		}
	}

	return res
}

func main() {
	testCases := []struct {
		nums    []int
		queries []int
		want    []int
	}{
		{
			nums:    []int{4, 5, 2, 1},
			queries: []int{3, 10, 21},
			want:    []int{2, 3, 4},
		},
		{
			nums:    []int{2, 3, 4, 5},
			queries: []int{1},
			want:    []int{0},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := answerQueries(tc.nums, tc.queries)
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
