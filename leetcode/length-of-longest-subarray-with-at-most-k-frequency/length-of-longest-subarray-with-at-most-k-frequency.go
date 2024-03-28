package main

import (
	"fmt"
)

func maxSubarrayLength(nums []int, k int) int {
	ids := make(map[int][]int)
	ans := 0
	start := 0

	for i, x := range nums {
		ids[x] = append(ids[x], i)
		if len(ids[x]) > k {
			start = max(start, ids[x][0]+1)
			ids[x] = ids[x][1:]
		}
		ans = max(ans, i-start+1)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			nums: []int{1, 2, 2, 1, 3},
			k:    1,
			want: 3,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3, 1, 2},
			k:    2,
			want: 6,
		},
		{
			nums: []int{1, 2, 1, 2, 1, 2, 1, 2},
			k:    1,
			want: 2,
		},
		{
			nums: []int{5, 5, 5, 5, 5, 5, 5},
			k:    4,
			want: 4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxSubarrayLength(tc.nums, tc.k)
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
