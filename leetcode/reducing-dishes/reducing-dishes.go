package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/reducing-dishes/

// -9 -8 -1 0 5
func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	sum, ans := 0, 0
	for i := len(satisfaction) - 1; i >= 0; i-- {
		sum += satisfaction[i]
		if sum < 0 {
			break
		}
		ans += sum
	}
	return ans
}

func main() {
	testCases := []struct {
		satisfaction []int
		want         int
	}{
		{
			satisfaction: []int{-1, -8, 0, 5, -9},
			want:         14,
		},
		{
			satisfaction: []int{4, 3, 2},
			want:         20,
		},
		{
			satisfaction: []int{-1, -4, -5},
			want:         0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxSatisfaction(tc.satisfaction)
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
