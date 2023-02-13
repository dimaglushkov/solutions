package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/third-maximum-number/

func thirdMax(nums []int) int {
	const minVal = -2147483649
	m := [3]int{minVal, minVal, minVal}
	for _, i := range nums {
		if i == m[0] || i == m[1] || i == m[2] {
			continue
		}
		if i > m[0] {
			m[0], m[1], m[2] = i, m[0], m[1]
		} else if i > m[1] {
			m[1], m[2] = i, m[1]
		} else if i > m[2] {
			m[2] = i
		}
	}

	if m[2] == minVal {
		return m[0]
	}
	return m[2]
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := thirdMax(tc.nums)
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
