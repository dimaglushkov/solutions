package main

import (
	"fmt"
)

func distance(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)

	lsum := make(map[int]int64)
	lcnt := make(map[int]int)
	rsum := make(map[int]int64)
	rcnt := make(map[int]int)

	for i, x := range nums {
		ans[i] += int64(lcnt[x]*i) - lsum[x]
		lsum[x] += int64(i)
		lcnt[x]++

		j, y := n-1-i, nums[n-1-i]
		ans[j] += rsum[y] - int64(rcnt[y]*j)
		rsum[y] += int64(j)
		rcnt[y]++
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []int64
	}{
		{
			nums: []int{0, 5, 3, 1, 2, 8, 6, 6, 6},
			want: []int64{0, 0, 0, 0, 0, 0, 3, 2, 3},
		},
		{
			nums: []int{1, 3, 1, 1, 2},
			want: []int64{5, 0, 3, 4, 0},
		},
		{
			nums: []int{0, 5, 3},
			want: []int64{0, 0, 0},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := distance(tc.nums)
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
