package main

import (
	"fmt"
)

func minDays(bloomDay []int, m int, k int) int {
	l, r := 1, int(1e9)
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		consecutiveLength, bouquets := 0, 0
		for _, bloom := range bloomDay {
			if bloom <= mid {
				consecutiveLength++
				if consecutiveLength >= k {
					consecutiveLength = 0
					bouquets++
				}
			} else {
				consecutiveLength = 0
			}
		}
		if bouquets >= m {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		bloomDay []int
		m        int
		k        int
		want     int
	}{
		{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        1,
			want:     3,
		},
		{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        2,
			want:     -1,
		},
		{
			bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
			m:        2,
			k:        3,
			want:     12,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDays(tc.bloomDay, tc.m, tc.k)
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
