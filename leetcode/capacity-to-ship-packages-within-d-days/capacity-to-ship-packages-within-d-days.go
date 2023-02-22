package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
func shipWithinDays(weights []int, days int) int {
	fitsCapacity := func(cap int) bool {
		s, d := 0, 0
		for i := 0; i < len(weights) && d < days; i++ {
			w := weights[i]
			if s+w <= cap {
				s += w
			} else {
				d++
				s = w
			}
		}
		return d < days
	}

	i, j := 0, 0
	for _, w := range weights {
		if w > i {
			i = w
		}
		j += w
	}

	for i < j {
		h := int(uint(i+j) >> 1)
		if !fitsCapacity(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func main() {
	testCases := []struct {
		weights []int
		days    int
		want    int
	}{
		{
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:    5,
			want:    15,
		},
		{
			weights: []int{3, 2, 2, 4, 1, 4},
			days:    3,
			want:    6,
		},
		{
			weights: []int{1, 2, 3, 1, 1},
			days:    4,
			want:    3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := shipWithinDays(tc.weights, tc.days)
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
