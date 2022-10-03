package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/minimum-time-to-make-rope-colorful/

//func minCost(c string, nt []int) int {
//	var l, r, res int
//	var n = len(c)
//	for l < n-1 {
//		r++
//		if c[l] != c[r] {
//			l = r
//			continue
//		}
//		for r < n && c[l] == c[r] {
//			r++
//		}
//
//		mi := l
//		for i := l; i < r; i++ {
//			if nt[i] > nt[mi] {
//				mi = i
//			}
//		}
//		for i := l; i < r; i++ {
//			if i != mi {
//				res += nt[i]
//			}
//		}
//		l = r
//	}
//	return res
//}

func minCost(colors string, neededTime []int) int {
	res, max := neededTime[0], neededTime[0]
	for i := 1; i < len(neededTime); i++ {
		res += neededTime[i]
		if colors[i] == colors[i-1] {
			if neededTime[i] > max {
				max = neededTime[i]
			}
		} else {
			res -= max
			max = neededTime[i]
		}
	}

	return res - max
}

func main() {
	// Example 1
	var colors1 string = "abaac"
	var neededTime1 = []int{1, 2, 3, 4, 5}
	fmt.Println("Expected: 3	Output: ", minCost(colors1, neededTime1))

	var c0 string = "aaa"
	var nt0 = []int{1, 1, 5}
	fmt.Println("Expected: 2	Output: ", minCost(c0, nt0))

	// Example 2
	var colors2 string = "abc"
	var neededTime2 = []int{1, 2, 3}
	fmt.Println("Expected: 0	Output: ", minCost(colors2, neededTime2))

	// Example 3
	var colors3 string = "aabaa"
	var neededTime3 = []int{1, 2, 3, 4, 1}
	fmt.Println("Expected: 2	Output: ", minCost(colors3, neededTime3))

}
