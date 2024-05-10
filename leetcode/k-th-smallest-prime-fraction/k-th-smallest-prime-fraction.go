package main

import (
	"fmt"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	l, r := 0.0, 1.0

	for l < r {
		m := (l + r) / 2
		cnt := 0
		maxFraction := []int{0, 1}

		for i, j := 0, 1; i < n-1; i++ {
			for j < n && float64(arr[i])/float64(arr[j]) > m {
				j++
			}
			cnt += n - j
			if j < n && float64(arr[i])/float64(arr[j]) > float64(maxFraction[0])/float64(maxFraction[1]) {
				maxFraction = []int{arr[i], arr[j]}
			}
		}

		if cnt < k {
			l = m
		} else if cnt > k {
			r = m
		} else {
			return maxFraction
		}
	}

	return []int{}
}

func main() {
	testCases := []struct {
		arr  []int
		k    int
		want []int
	}{
		{
			arr:  []int{1, 2, 3, 5},
			k:    3,
			want: []int{2, 5},
		},
		{
			arr:  []int{1, 7},
			k:    1,
			want: []int{1, 7},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := kthSmallestPrimeFraction(tc.arr, tc.k)
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
