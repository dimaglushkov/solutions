package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, n*(n+1)/2)

	for i := 0; i < n; i++ {
		t := nums[i]
		sums = append(sums, t)
		for j := i + 1; j < n; j++ {
			t += nums[j]
			sums = append(sums, t)
		}
	}

	sort.Ints(sums)

	ans := 0
	for i := left - 1; i < right; i++ {
		ans = (ans + sums[i]) % 1000000007
	}

	return ans
}

func main() {
	testCases := []struct {
		nums  []int
		n     int
		left  int
		right int
		want  int
	}{
		{
			nums:  []int{1, 2, 3, 4},
			n:     4,
			left:  1,
			right: 5,
			want:  13,
		},
		{
			nums:  []int{1, 2, 3, 4},
			n:     4,
			left:  3,
			right: 4,
			want:  6,
		},
		{
			nums:  []int{1, 2, 3, 4},
			n:     4,
			left:  1,
			right: 10,
			want:  50,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := rangeSum(tc.nums, tc.n, tc.left, tc.right)
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
