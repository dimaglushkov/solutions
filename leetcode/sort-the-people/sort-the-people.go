package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	type pair struct {
		x string
		y int
	}

	pairs := make([]pair, len(names))
	for i := range names {
		pairs[i] = pair{names[i], heights[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].y > pairs[j].y
	})

	for i := range pairs {
		names[i] = pairs[i].x
	}

	return names
}

func main() {
	testCases := []struct {
		names   []string
		heights []int
		want    []string
	}{
		{
			names:   []string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"},
			heights: []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224},
			want:    []string{"EPCFFt", "RPJOFYZUBFSIYp", "VOYGWWNCf", "Vk", "Sgizfdfrims", "IEO", "QTASHKQ", "WSpmqvb"},
		},
		{
			names:   []string{"Mary", "John", "Emma"},
			heights: []int{180, 165, 170},
			want:    []string{"Mary", "Emma", "John"},
		},
		{
			names:   []string{"Alice", "Bob", "Bob"},
			heights: []int{155, 185, 150},
			want:    []string{"Bob", "Alice", "Bob"},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := sortPeople(tc.names, tc.heights)
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
