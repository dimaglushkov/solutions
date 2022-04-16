package main

import (
	"fmt"
	"math"
	"strconv"
)

// source: https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	var start, end, digitNum int

	for start < len(s) && s[start] == ' ' {
		start++
	}
	if start == len(s) {
		return 0
	}

	if s[start] == '-' {
		end = start + 1
	} else if s[start] == '+' {
		start++
		end = start
	} else {
		end = start
	}

	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		digitNum++
		end++
	}

	if digitNum == 0 {
		return 0
	}

	val, err := strconv.ParseInt(s[start:end], 10, 32)
	if err != nil {
		if s[start] == '-' {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	return int(val)
}

func main() {
	// Example 1
	var s1 string = "42"
	fmt.Println("Expected: 42	Output: ", myAtoi(s1))

	// Example 2
	var s2 string = "   -42"
	fmt.Println("Expected: -42	Output: ", myAtoi(s2))

	// Example 3
	var s3 string = "4193 with words"
	fmt.Println("Expected: 4193	Output: ", myAtoi(s3))

	// Example 4
	var s4 string = "words and 987"
	fmt.Println("Expected: 0	Output: ", myAtoi(s4))

	// Example 5
	var s5 string = "-+12"
	fmt.Println("Expected: 0	Output: ", myAtoi(s5))

	// Example 6
	var s6 string = " "
	fmt.Println("Expected: 0	Output: ", myAtoi(s6))

	// Example 7
	var s7 string = "-121237182371237981927839781293781122112"
	fmt.Println("Expected: -2147483648	Output: ", myAtoi(s7))

	// Example 8
	var s8 string = "+1"
	fmt.Println("Expected: 1	Output: ", myAtoi(s8))

	// Example 9
	var s9 string = "-1"
	fmt.Println("Expected: -1	Output: ", myAtoi(s9))

	// Example 10
	var s10 string = "      -11-1"
	fmt.Println("Expected: 0	Output: ", myAtoi(s10))

}
