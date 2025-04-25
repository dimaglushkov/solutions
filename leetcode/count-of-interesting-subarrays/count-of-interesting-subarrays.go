package main

import (
	"fmt"
)

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	var ans int64
	eq := 0
	mem := make(map[int]int)
	mem[0] = 1

	for _, i := range nums {
		if i%modulo == k {
			eq++
		}
		need := (eq%modulo - k + modulo) % modulo
		ans += int64(mem[need])
		mem[eq%modulo]++
	}

	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		modulo int
		k      int
		want   int64
	}{
		{
			nums:   []int{3, 2, 4},
			modulo: 2,
			k:      1,
			want:   3,
		},
		{
			nums:   []int{3, 1, 9, 6},
			modulo: 3,
			k:      0,
			want:   2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countInterestingSubarrays(tc.nums, tc.modulo, tc.k)
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
