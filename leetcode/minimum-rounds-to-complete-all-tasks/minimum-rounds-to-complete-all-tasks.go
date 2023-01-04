package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/

func minimumRounds(tasks []int) int {
	cnt := make(map[int]int)
	for _, x := range tasks {
		cnt[x]++
	}
	res := 0
	for _, c := range cnt {
		for c > 4 {
			c -= 3
			res++
		}
		switch c {
		case 4:
			res += 2
		case 3:
			res++
		case 2:
			res++
		case 1:
			return -1
		}
	}
	return res
}

func main() {
	testCases := []struct {
		tasks []int
		want  int
	}{
		{
			tasks: []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4},
			want:  4,
		},
		{
			tasks: []int{2, 3, 3},
			want:  -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumRounds(tc.tasks)
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
