package main

import (
	"fmt"
	"sort"
)

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] < tasks[j][1]-tasks[j][0]
	})

	ans := 0

	for _, t := range tasks {
		if ans+t[0] > t[1] {
			ans = ans + t[0]
		} else {
			ans = t[1]
		}
	}

	return ans
}

func main() {
    testCases := []struct {
		tasks [][]int
		want int
    }{
		{
			tasks: [][]int {{1,2},{2,4},{4,8}},
			want:  8,
		},
		{
			tasks: [][]int {{1,3},{2,4},{10,11},{10,12},{8,9}},
			want:  32,
		},
		{
			tasks: [][]int {{1,7},{2,8},{3,9},{4,10},{5,11},{6,12}},
			want:  27,
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := minimumEffort(tc.tasks)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
