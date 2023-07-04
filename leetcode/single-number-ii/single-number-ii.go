package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var ones, twos int

	for _, i := range nums {
		ones ^= i & ^twos
		twos ^= i & ^ones
	}

	return ones
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		{
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := singleNumber(tc.nums)
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
