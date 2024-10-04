package main

import (
	"fmt"
)

func minSubarray(nums []int, p int) int {
	totalSum := 0
	for i := range nums {
		totalSum = (nums[i] + totalSum) % p
	}

	target := totalSum % p
	if target == 0 {
		return 0
	}

	prefixMap := map[int]int{0: -1}
	prefixSum := 0
	ans := len(nums)

	for i, num := range nums {
		prefixSum = (prefixSum + num) % p

		if val, ok := prefixMap[(prefixSum-target+p)%p]; ok && i-val < ans {
			ans = i - val
		}

		prefixMap[prefixSum] = i
	}

	if ans == len(nums) {
		return -1
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		p    int
		want int
	}{
		{
			nums: []int{3, 1, 4, 2},
			p:    6,
			want: 1,
		},
		{
			nums: []int{6, 3, 5, 2},
			p:    9,
			want: 2,
		},
		{
			nums: []int{1, 2, 3},
			p:    3,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minSubarray(tc.nums, tc.p)
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
