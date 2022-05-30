package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/divide-two-integers/

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	labs := func(x int) int64 {
		if x < 0 {
			return int64(-x)
		}
		return int64(x)
	}

	var sign = 1
	var res = 0
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	var x, y = labs(dividend), labs(divisor)

	for x >= y {
		tmp := y
		cnt := 1
		for tmp<<1 <= x {
			tmp <<= 1
			cnt <<= 1
		}
		x -= tmp
		res += cnt
	}

	return res * sign
}

func main() {
	// Example 1
	var dividend1 int = 10
	var divisor1 int = 3
	fmt.Println("Expected: 3	Output: ", divide(dividend1, divisor1))

	// Example 2
	var dividend2 int = 7
	var divisor2 int = -3
	fmt.Println("Expected: -2	Output: ", divide(dividend2, divisor2))

	// Example 3
	var dividend3 = -2147483648
	var divisor3 = -1
	fmt.Println("Expected: 2147483647	Output: ", divide(dividend3, divisor3))

}
