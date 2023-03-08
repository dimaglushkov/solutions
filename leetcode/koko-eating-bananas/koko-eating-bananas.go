package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/koko-eating-bananas/

func minEatingSpeed(piles []int, h int) int {
	check := func(k int) bool {
		hours := 0
		for _, p := range piles {
			hours += p / k
			if p%k != 0 {
				hours++
			}
		}
		return hours > h
	}

	i, j := 1, 1<<31-1
	for i < j {
		m := int(uint(i+j) >> 1)
		if check(m) { // if f(m) less
			i = m + 1
		} else {
			j = m
		}
	}
	return i
}

func main() {
	testCases := []struct {
		piles []int
		h     int
		want  int
	}{
		{
			piles: []int{3, 6, 7, 11},
			h:     8,
			want:  4,
		},
		{
			piles: []int{30, 11, 23, 4, 20},
			h:     5,
			want:  30,
		},
		{
			piles: []int{30, 11, 23, 4, 20},
			h:     6,
			want:  23,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minEatingSpeed(tc.piles, tc.h)
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
