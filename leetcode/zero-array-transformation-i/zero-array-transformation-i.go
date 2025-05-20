package main

import (
	"fmt"
	"sort"
)

type qs struct {
	m     map[int]int
	total int
}

func (q *qs) add(x int) {
	q.m[x]++
	q.total++
}

func (q *qs) remove(x int) {
	q.total -= q.m[x]
}

func (q *qs) count() int {
	return q.total
}

// source: https://leetcode.com/problems/zero-array-transformation-i/
func isZeroArray(nums []int, queries [][]int) bool {
	sort.Slice(queries, func(i, j int) bool {
		if queries[i][0] == queries[j][0] {
			return queries[i][1] < queries[j][1]
		}
		return queries[i][0] < queries[j][0]
	})

	var qi int
	q := qs{m: make(map[int]int), total: 0}

	for i := range nums {
		for qi < len(queries) && queries[qi][0] <= i {
			q.add(queries[qi][1])
			qi++
		}

		if nums[i]-q.count() > 0 {
			return false
		}

		q.remove(i)
	}

	return true
}

func main() {
	fmt.Println(true, isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))
	fmt.Println(false, isZeroArray([]int{1, 6, 6, 7}, [][]int{{1, 2}, {2, 2}, {3, 3}, {2, 2}, {1, 2}, {3, 3}, {3, 3}, {0, 2}, {0, 1}, {0, 0}, {0, 3}, {1, 3}}))
	fmt.Println(true, isZeroArray([]int{7, 0, 9}, [][]int{{0, 0}, {2, 2}, {2, 2}, {1, 1}, {1, 2}, {0, 2}, {0, 2}, {0, 0}, {2, 2}, {0, 2}, {0, 2}, {0, 2}, {0, 2}, {2, 2}, {0, 0}}))
}
