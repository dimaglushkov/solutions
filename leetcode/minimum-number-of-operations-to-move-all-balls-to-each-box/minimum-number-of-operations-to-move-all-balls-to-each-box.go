package main

import "fmt"

// source: https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minOperations1(boxes string) []int {
	res := make([]int, len(boxes))
	ones := make(map[int]bool)
	for i := range boxes {
		if boxes[i] == '1' {
			ones[i] = true
		}
	}

	for i := range boxes {
		for oneId := range ones {
			res[i] += abs(i - oneId)
		}
	}
	return res
}

func minOperations(boxes string) []int {
	pref := make([]int, len(boxes)+1)
	cnt := 0
	for i, c := range boxes {
		pref[i+1] = pref[i]
		if c == '1' {
			cnt++
		}
		pref[i+1] += cnt
	}

	res := make([]int, len(boxes))
	suf, cnt := 0, 0
	for i := len(boxes) - 1; i >= 0; i-- {
		res[i] = pref[i] + suf
		if boxes[i] == '1' {
			cnt++
		}
		suf += cnt
	}
	return res
}

func main() {
	// Example 1
	var boxes1 string = "110"
	fmt.Println("Expected: [1,1,3]	Output: ", minOperations(boxes1))

	// Example 2
	var boxes2 string = "001011"
	fmt.Println("Expected: [11,8,5,4,3,4]	Output: ", minOperations(boxes2))

}
