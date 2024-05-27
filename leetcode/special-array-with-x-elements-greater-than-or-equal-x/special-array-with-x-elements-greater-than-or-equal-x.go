package main

import (
	"fmt"
)

func specialArray(nums []int) int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v] += 1
	}

	ans := -1
	for i := 1; i <= len(nums); i++ {
		count := 0
		for k, v := range cnt {
			if k >= i {
				count += v
			}
		}
		if count == i {
			ans = i
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
			nums: []int{3, 5},
			want: 2,
		},
		{
			nums: []int{0, 0},
			want: -1,
		},
		{
			nums: []int{0, 4, 3, 0, 4},
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := specialArray(tc.nums)
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
