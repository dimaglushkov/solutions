package main

import (
	"fmt"
	"math"
)

func maxN(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] > x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func minimizedMaximum(n int, quantities []int) int {
	low, high := 1, maxN(quantities...)
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		storesNeeded := 0
		for _, quantity := range quantities {
			storesNeeded += int(math.Ceil(float64(quantity) / float64(mid)))
		}

		if storesNeeded <= n {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		n          int
		quantities []int
		want       int
	}{
		{
			n:          6,
			quantities: []int{11, 6},
			want:       3,
		},
		{
			n:          7,
			quantities: []int{15, 10, 10},
			want:       5,
		},
		{
			n:          1,
			quantities: []int{100000},
			want:       100000,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimizedMaximum(tc.n, tc.quantities)
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
