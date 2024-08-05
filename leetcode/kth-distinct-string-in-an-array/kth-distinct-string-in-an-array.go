package main

import (
	"fmt"
)

func kthDistinct(arr []string, k int) string {
	cnt := make(map[string]int)

	for _, s := range arr {
		cnt[s]++
	}

	for _, s := range arr {
		if cnt[s] == 1 {
			k--
			if k == 0 {
				return s
			}
		}
	}

	return ""

}

func main() {
	testCases := []struct {
		arr  []string
		k    int
		want string
	}{
		{
			arr:  []string{"d", "b", "c", "b", "c", "a"},
			k:    2,
			want: "a",
		},
		{
			arr:  []string{"aaa", "aa", "a"},
			k:    1,
			want: "aaa",
		},
		{
			arr:  []string{"a", "b", "a"},
			k:    3,
			want: "",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := kthDistinct(tc.arr, tc.k)
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
