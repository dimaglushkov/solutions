package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	ans := make([]int, 2)
	fbp := 0

	for _, v := range nums {
		ans[0] ^= v
	}

	for i := 0; i < 32; i++ {
		if ans[0]&(1<<i) > 0 {
			fbp = i
			break
		}
	}
	ans[0] = 0

	for _, v := range nums {
		if v&(1<<fbp) > 0 {
			ans[0] ^= v
		} else {
			ans[1] ^= v
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 1, 3, 2, 5},
			want: []int{3, 5},
		},
		{
			nums: []int{-1, 0},
			want: []int{-1, 0},
		},
		{
			nums: []int{0, 1},
			want: []int{1, 0},
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
