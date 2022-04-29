package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/number-of-ways-to-split-a-string/

func numWays(s string) int {
	const mod = 1000000007
	var onesCnt, onesInStr int
	for _, r := range s {
		if r == '1' {
			onesCnt++
		}
	}
	if onesCnt%3 != 0 {
		return 0
	}
	if onesCnt == 0 {
		return ((len(s) - 1) * (len(s) - 2) / 2) % mod
	}

	onesInStr = onesCnt / 3
	var l, r = 0, len(s) - 1
	for s[l] == '0' {
		l++
	}
	for s[r] == '0' {
		r--
	}
	s = s[l : r+1]

	var res = 1
	var i = 0

	for i < len(s) {
		freeZerosCnt := 1
		curOnesCnt := 0
		for curOnesCnt < onesInStr {
			for s[i] != '1' {
				i++
			}
			curOnesCnt++
			i++
		}
		for i < len(s) && s[i] != '1' {
			freeZerosCnt++
			i++
		}
		res *= freeZerosCnt
	}

	return res % mod
}

func main() {
	// Example 4
	var s4 string = "100100010100110"
	fmt.Println("Expected: 12	Output: ", numWays(s4))

	// Example 1
	var s1 string = "10101"
	fmt.Println("Expected: 4	Output: ", numWays(s1))

	// Example 2
	var s2 string = "1001"
	fmt.Println("Expected: 0	Output: ", numWays(s2))

	// Example 3
	var s3 string = "0000"
	fmt.Println("Expected: 3	Output: ", numWays(s3))

}
