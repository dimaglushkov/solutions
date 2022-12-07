package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/combination-sum-ii/
// 1,1,2,5,6,7,10
//func deepEqual(x, y []int) bool {
//	if len(x) != len(y) {
//		return false
//	}
//	for i := range x {
//		if x[i] != y[i] {
//			return false
//		}
//	}
//	return true
//}

func combinationSum2(c []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(c)
	n := len(c)

	var util func(arr []int, t, i int)
	util = func(arr []int, t, i int) {
		if t == 0 {
			x := make([]int, len(arr))
			copy(x, arr)
			res = append(res, x)
			return
		}

		for j := i; j < n; j++ {
			if j > i && c[j] == c[j-1] {
				continue
			}

			if c[j] > t {
				break
			}
			arr = append(arr, c[j])
			util(arr, t-c[j], j+1)
			arr = arr[:len(arr)-1]
		}
	}

	util([]int{}, target, 0)
	return res
}

func main() {
	testCases := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{3, 1, 3, 5, 1, 1},
			target:     8,
			want: [][]int{
				{1, 1, 1, 5},
				{1, 1, 3, 3},
				{3, 5},
			},
		},
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := combinationSum2(tc.candidates, tc.target)
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
