package main

import (
	"fmt"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSum(nums []int) int {
	ans := -1
	maxDigit := make(map[int][]int)

	sort.Ints(nums)
	for i, x := range nums {
		md := 0
		for x > 0 {
			md = max(md, x%10)
			x /= 10
		}
		maxDigit[md] = append(maxDigit[md], nums[i])
	}

	for i := 0; i < 10; i++ {
		nmd := len(maxDigit[i])
		if nmd < 2 {
			continue
		}

		ans = max(ans, maxDigit[i][nmd-1]+maxDigit[i][nmd-2])
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{51, 71, 17, 24, 42},
			want: 88,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxSum(tc.nums)
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
