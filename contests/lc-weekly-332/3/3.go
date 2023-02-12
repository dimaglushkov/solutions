package main

import (
	"strconv"
	"strings"
)

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
