package main

import (
	"fmt"
	"sort"
)

func beautifulSubsets(nums []int, k int) int {
	ans := 0
	sort.Ints(nums)
	var vis [1001]int

	var dfs func(s int)
	dfs = func(s int) {
		for i := s; i < len(nums); i++ {
			if nums[i]-k < 0 || vis[nums[i]-k] == 0 {
				ans++

				vis[nums[i]]++
				dfs(i + 1)
				vis[nums[i]]--
			}
		}
	}

	dfs(0)
	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{2, 4, 6},
			k:    2,
			want: 4,
		},
		{
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			nums: []int{2, 4, 6, 8},
			k:    2,
			want: 7,
		},
		{
			nums: []int{4, 2, 5, 9, 10, 3},
			k:    1,
			want: 23,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := beautifulSubsets(tc.nums, tc.k)
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
