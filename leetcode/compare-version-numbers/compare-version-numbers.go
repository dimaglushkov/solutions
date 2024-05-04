package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := make([]int, 0), make([]int, 0)

	for _, v1r := range strings.Split(version1, ".") {
		x, _ := strconv.Atoi(v1r)
		v1 = append(v1, x)
	}

	for _, v2r := range strings.Split(version2, ".") {
		x, _ := strconv.Atoi(v2r)
		v2 = append(v2, x)
	}

	for len(v1) < len(v2) {
		v1 = append(v1, 0)
	}
	for len(v2) < len(v1) {
		v2 = append(v2, 0)
	}

	for i := range v1 {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}

	return 0
}

func main() {
	testCases := []struct {
		version1 string
		version2 string
		want     int
	}{
		{
			version1: "1.01",
			version2: "1.001",
			want:     0,
		},
		{
			version1: "1.0",
			version2: "1.0.0",
			want:     0,
		},
		{
			version1: "0.1",
			version2: "1.1",
			want:     -1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := compareVersion(tc.version1, tc.version2)
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
