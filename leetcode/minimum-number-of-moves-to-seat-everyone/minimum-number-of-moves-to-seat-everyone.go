package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMovesToSeat(seats []int, students []int) int {
	ans := 0

	sort.Ints(seats)
	sort.Ints(students)

	for i := range students {
		ans += abs(seats[i] - students[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		seats    []int
		students []int
		want     int
	}{
		{
			seats:    []int{12, 14, 19, 19, 12},
			students: []int{19, 2, 17, 20, 7},
			want:     19,
		},
		{
			seats:    []int{3, 1, 5},
			students: []int{2, 7, 4},
			want:     4,
		},
		{
			seats:    []int{4, 1, 5, 9},
			students: []int{1, 3, 2, 6},
			want:     7,
		},
		{
			seats:    []int{2, 2, 6, 6},
			students: []int{1, 3, 2, 6},
			want:     4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minMovesToSeat(tc.seats, tc.students)
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
