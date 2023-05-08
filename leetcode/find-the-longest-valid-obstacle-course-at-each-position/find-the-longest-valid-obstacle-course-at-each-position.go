package main

import (
	"fmt"
	"sort"
)

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var stack []int
	ans := make([]int, len(obstacles))
	for i, obstacle := range obstacles {
		if len(stack) == 0 || obstacle >= stack[len(stack)-1] {
			stack = append(stack, obstacle)
			ans[i] = len(stack)
			continue
		}
		index := sort.Search(len(stack), func(j int) bool {
			return stack[j] > obstacle
		})
		ans[i] = index + 1
		stack[index] = obstacle
	}
	return ans
}

func main() {
	testCases := []struct {
		obstacles []int
		want      []int
	}{
		{
			obstacles: []int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4},
			want:      []int{1, 1, 2, 3, 2, 3, 4, 5, 3, 5},
		},
		// 1: 5
		// 1: 1
		// 2: 1 5
		// 3: 1 5 5
		// 2: 1 1
		// 3: 1 1 3
		// 4: 1 1 3 4
		// 5: 1 1 3 4
		// 3: 1 1 1
		// 5: 1 1 4 4
		{

			obstacles: []int{1, 2, 3, 2},
			want:      []int{1, 2, 3, 3},
		},
		{
			obstacles: []int{2, 2, 1},
			want:      []int{1, 2, 1},
		},
		{
			obstacles: []int{3, 1, 5, 6, 4, 2},
			want:      []int{1, 1, 2, 3, 2, 2},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := longestObstacleCourseAtEachPosition(tc.obstacles)
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
