package main

import (
	"fmt"
)

func countSubarrays(nums []int, k int) int64 {
	mv := 0
	var mi []int
	var ans int64

	for i, x := range nums {
		if x > mv {
			mv, ans, mi = x, 0, []int{}
		}

		if x == mv {
			mi = append(mi, i)
		}

		if len(mi) >= k {
			ans += int64(mi[len(mi)-k]) + 1
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want int64
	}{
		{
			nums: []int{1, 3, 2, 3, 3},
			k:    2,
			want: 6,
		},
		{
			nums: []int{1, 3, 2, 3, 3, 1},
			k:    2,
			want: 10,
		},
		{
			nums: []int{1, 4, 2, 1},
			k:    3,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countSubarrays(tc.nums, tc.k)
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
