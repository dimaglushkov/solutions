package main

import (
	"fmt"
)

func maxFrequencyElements(nums []int) int {
	cnt := [101]int{}
	ans := 0

	for _, i := range nums {
		cnt[i]++
	}

	m := 0
	for _, c := range cnt {
		if c == m {
			ans += c
		}
		if c > m {
			ans = c
			m = c
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
			nums: []int{1, 2, 2, 3, 1, 4},
			want: 4,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxFrequencyElements(tc.nums)
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
