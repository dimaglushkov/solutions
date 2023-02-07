package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/fruit-into-baskets/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func totalFruit(fruits []int) int {
	m := make(map[int]int)
	l, r, ans := 0, 0, 0

	for r < len(fruits) {
		m[fruits[r]]++

		for len(m) > 2 {
			m[fruits[l]]--
			if m[fruits[l]] == 0 {
				delete(m, fruits[l])
			}
			l++
		}

		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func main() {
	testCases := []struct {
		fruits []int
		want   int
	}{
		{
			fruits: []int{1, 0, 3, 4, 3},
			want:   3,
		},
		{
			fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want:   5,
		},
		{
			fruits: []int{1, 2, 1},
			want:   3,
		},
		{
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},
		{
			fruits: []int{1, 2, 3, 2, 2},
			want:   4,
		},
		{
			fruits: []int{4, 1, 1, 1, 3, 1, 7, 5},
			want:   5,
		},
		{
			fruits: []int{1, 1},
			want:   2,
		},
		{
			fruits: []int{1, 1, 1, 1, 2, 3, 2, 2},
			want:   5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := totalFruit(tc.fruits)
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
