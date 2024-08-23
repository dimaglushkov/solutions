package main

import (
	"fmt"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func fractionAddition(exp string) string {
	var nums []int

	// parsing numbers
	for l, r := 0, 1; r < len(exp); r++ {
		if exp[r] == '+' || exp[r] == '-' || exp[r] == '/' {
			x, _ := strconv.Atoi(exp[l:r])
			nums = append(nums, x)

			l = r
			if exp[r] == '/' {
				l++
			}

		} else if r == len(exp)-1 {
			x, _ := strconv.Atoi(exp[l:])
			nums = append(nums, x)
		}
	}

	// calculating
	var x, y = nums[0], nums[1]
	for i := 2; i < len(nums); i += 2 {
		dx, dy := nums[i], nums[i+1]

		if y != dy {
			dd := lcm(y, dy)
			d1, d2 := dd/y, dd/dy
			x, y, dx, dy = d1*x, d1*y, d2*dx, d2*dy
		}

		x += dx
	}

	// minimize
	for gcd(x, y) != 1 && gcd(x, y) != -1 {
		d := gcd(x, y)
		if d < 0 {
			d *= -1
		}
		x, y = x/d, y/d
	}

	return strconv.Itoa(x) + "/" + strconv.Itoa(y)
}

func main() {
	testCases := []struct {
		expression string
		want       string
	}{
		{
			expression: "1/2-4/1-4/3+1/2-5/1",
			want:       "-28/3",
		},
		{
			expression: "1/3-1/2",
			want:       "-1/6",
		},
		{
			expression: "-1/2+1/2",
			want:       "0/1",
		},
		{
			expression: "-1/2+1/2+1/3",
			want:       "1/3",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := fractionAddition(tc.expression)
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
