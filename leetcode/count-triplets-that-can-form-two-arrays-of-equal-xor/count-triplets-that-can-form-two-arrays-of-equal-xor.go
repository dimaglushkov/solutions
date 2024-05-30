package main

import (
	"fmt"
)

func countTriplets(arr []int) int {
	ans := 0
	n := len(arr)
	for i, val := range arr {
		for k := i + 1; k < n; k++ {
			val ^= arr[k]
			if val == 0 {
				ans += k - i
			}
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{2, 3, 1, 6, 7},
			want: 4,
		},
		{
			arr:  []int{1, 1, 1, 1, 1},
			want: 10,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countTriplets(tc.arr)
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
