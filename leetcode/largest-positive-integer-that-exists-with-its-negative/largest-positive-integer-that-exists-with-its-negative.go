package main

import (
	"fmt"
)

func findMaxK(nums []int) int {
	ans := -1
	m := make(map[int]bool, len(nums))

	for _, x := range nums {
		if x > 0 {
			m[x] = m[-x]
		} else {
			m[x] = true
			if _, ok := m[-x]; ok {
				m[-x] = true
			}
		}
	}

	for i := range nums {
		if nums[i] > ans && m[nums[i]] {
			ans = nums[i]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{-1, 2, -3, 3},
			want: 3,
		},
		{
			nums: []int{-1, 10, 6, 7, -7, 1},
			want: 7,
		},
		{
			nums: []int{-10, 8, 6, 7, -2, -3},
			want: -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findMaxK(tc.nums)
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
