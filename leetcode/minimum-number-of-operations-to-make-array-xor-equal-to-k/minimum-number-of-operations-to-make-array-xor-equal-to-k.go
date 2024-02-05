package main

import (
	"fmt"
)

func minOperations(nums []int, k int) int {
	c := 0
	ans := 0
	for _, num := range nums {
		c ^= num
	}

	d := c ^ k

	for i := 0; i < 32; i++ { // Assuming 32-bit integers
		if d&(1<<i) != 0 {
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{2, 1, 3, 4},
			k:    1,
			want: 2,
		},
		{
			nums: []int{2, 0, 2, 0},
			k:    0,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minOperations(tc.nums, tc.k)
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
