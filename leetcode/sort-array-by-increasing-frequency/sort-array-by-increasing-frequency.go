package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	cnt := make(map[int]int)

	for i := range nums {
		cnt[nums[i]]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if cnt[nums[i]] == cnt[nums[j]] {
			return nums[i] > nums[j]
		}
		return cnt[nums[i]] < cnt[nums[j]]
	})

	return nums
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 1, 2, 2, 2, 3},
			want: []int{3, 1, 1, 2, 2, 2},
		},
		{
			nums: []int{2, 3, 1, 3, 2},
			want: []int{1, 3, 3, 2, 2},
		},
		{
			nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			want: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := frequencySort(tc.nums)
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
