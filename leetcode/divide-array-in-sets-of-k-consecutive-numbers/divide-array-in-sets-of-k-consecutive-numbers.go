package main

import (
	"fmt"
	"sort"
)

func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)

	sort.Ints(nums)
	cnt := make(map[int]int)
	for i := range nums {
		cnt[nums[i]]++
	}

	if n%k != 0 {
		return false
	}

	for i := range nums {
		if cnt[nums[i]] > 0 {
			cnt[nums[i]]--

			for j := 1; j < k; j++ {
				if cnt[nums[i]+j] == 0 {
					return false
				}

				cnt[nums[i]+j]--
			}
		}
	}

	return true
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 3, 4, 4, 5, 6},
			k:    4,
			want: true,
		},
		{
			nums: []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			k:    3,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isPossibleDivide(tc.nums, tc.k)
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
