package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/minimum-time-to-complete-trips/

func minimumTime(time []int, totalTrips int) int64 {
	tt := int64(totalTrips)
	check := func(t int64) bool {
		var cnt int64
		for _, i := range time {
			cnt += t / int64(i)
			if cnt > tt {
				break
			}
		}
		return cnt < tt
	}

	i, j := int64(0), int64(1<<63-1)
	for i < j {
		h := int64(uint(i+j) >> 1)
		if check(h) { // if f(h) less
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func main() {
	testCases := []struct {
		time       []int
		totalTrips int
		want       int64
	}{
		{
			time:       []int{1, 2, 3},
			totalTrips: 5,
			want:       3,
		},
		{
			time:       []int{2},
			totalTrips: 1,
			want:       2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumTime(tc.time, tc.totalTrips)
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
