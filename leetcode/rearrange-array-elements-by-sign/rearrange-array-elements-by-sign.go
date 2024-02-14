package main

import (
	"fmt"
)

func rearrangeArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	ip, in := 0, 0
	for i := 0; i < n; i += 2 {
		for in < n && nums[in] > 0 {
			in++
		}
		for ip < n && nums[ip] < 0 {
			ip++
		}

		ans[i], ans[i+1] = nums[ip], nums[in]

		ip++
		in++
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{3, 1, -2, -5, 2, -4},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			nums: []int{-1, 1},
			want: []int{1, -1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rearrangeArray(tc.nums)
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
