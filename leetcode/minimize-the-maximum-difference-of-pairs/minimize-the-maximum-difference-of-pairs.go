package main

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	check := func(val int) bool {
		cnt := 0
		for i := 0; i < n-1 && cnt < p; {
			if nums[i+1]-nums[i] <= val {
				i += 2
				cnt++
			} else {
				i++
			}
		}

		return cnt >= p
	}

	return sort.Search(nums[n-1]-nums[0], check)
}

func main() {
	testCases := []struct {
		nums []int
		p    int
		want int
	}{
		{
			nums: []int{1, 2},
			p:    1,
			want: 1,
		},
		{
			nums: []int{3, 3, 5, 1, 0, 5, 6, 6},
			p:    4,
			want: 1,
		},
		{
			nums: []int{1, 10, 10, 25, 50, 51},
			p:    2,
			want: 1,
		},
		{
			nums: []int{1, 10, 10, 25, 50, 51},
			p:    3,
			want: 15,
		}, {
			nums: []int{10, 1, 2, 7, 1, 3},
			p:    2,
			want: 1,
		},
		{
			nums: []int{4, 2, 1, 2},
			p:    1,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimizeMax(tc.nums, tc.p)
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
