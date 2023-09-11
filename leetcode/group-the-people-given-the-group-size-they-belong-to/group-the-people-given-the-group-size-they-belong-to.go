package main

import (
	"fmt"
)

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	ans := make([][]int, 0)

	for i, x := range groupSizes {
		if _, ok := groups[x]; !ok {
			groups[x] = make([]int, 0, x)
		}

		groups[x] = append(groups[x], i)
		if len(groups[x]) == x {
			ans = append(ans, groups[x])
			delete(groups, x)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		groupSizes []int
		want       [][]int
	}{
		{
			groupSizes: []int{3, 3, 3, 3, 3, 1, 3},
			want:       [][]int{{5}, {0, 1, 2}, {3, 4, 6}},
		},
		{
			groupSizes: []int{2, 1, 3, 3, 3, 2},
			want:       [][]int{{1}, {0, 5}, {2, 3, 4}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := groupThePeople(tc.groupSizes)
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
