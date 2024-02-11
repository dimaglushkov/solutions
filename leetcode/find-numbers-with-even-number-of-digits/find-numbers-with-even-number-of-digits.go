package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	ans := 0

	for _, x := range nums {
		if len(strconv.Itoa(x))%2 == 0 {
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findNumbers(tc.nums)
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
