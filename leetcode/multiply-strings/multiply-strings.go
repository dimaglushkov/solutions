package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	prod := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			x := int(num1[i]-'0') * int(num2[j]-'0')
			prod[i+j+1] += x
			prod[i+j] += prod[i+j+1] / 10
			prod[i+j+1] = prod[i+j+1] % 10
		}
	}
	ans := ""
	idx := 0
	for idx < len(prod) && prod[idx] == 0 {
		idx++
	}

	for i := idx; i < len(prod); i++ {
		ans += strconv.Itoa(prod[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			num1: "17",
			num2: "23",
			want: "391",
		},
		{
			num1: "2",
			num2: "3",
			want: "6",
		},
		{
			num1: "123",
			num2: "456",
			want: "56088",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := multiply(tc.num1, tc.num2)
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
