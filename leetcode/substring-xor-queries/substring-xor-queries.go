package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/substring-xor-queries/

func substringXorQueries(s string, queries [][]int) [][]int {
	n := len(queries)
	ans := make([][]int, 0, n)

	for _, q := range queries {
		target := strconv.FormatInt(int64(q[0]^q[1]), 2)
		if ind := strings.Index(s, target); ind != -1 {
			ans = append(ans, []int{ind, ind + len(target) - 1})
		} else {
			ans = append(ans, []int{-1, -1})
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s       string
		queries [][]int
		want    [][]int
	}{
		{
			s:       "101101",
			queries: [][]int{{0, 5}, {1, 2}},
			want:    [][]int{{0, 2}, {2, 3}},
		},
		{
			s:       "0101",
			queries: [][]int{{12, 8}},
			want:    [][]int{{-1, -1}},
		},
		{
			s:       "1",
			queries: [][]int{{4, 5}},
			want:    [][]int{{0, 0}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := substringXorQueries(tc.s, tc.queries)
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
