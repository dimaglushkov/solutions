package main

import (
	"fmt"
)

func finalString(s string) string {
	iCnt := 0
	for _, c := range []byte(s) {
		if c == 'i' {
			iCnt++
		}
	}

	if iCnt == 0 {
		return s
	}

	reverse := true
	start, end := 0, len(s)-iCnt-1
	ans := make([]byte, len(s)-iCnt)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'i' {
			reverse = !reverse
			continue
		}

		if reverse {
			ans[end] = s[i]
			end--
		} else {
			ans[start] = s[i]
			start++
		}
	}

	return string(ans)
}

func main() {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "string",
			want: "rtsng",
		},
		{
			s:    "poiinter",
			want: "ponter",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := finalString(tc.s)
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
