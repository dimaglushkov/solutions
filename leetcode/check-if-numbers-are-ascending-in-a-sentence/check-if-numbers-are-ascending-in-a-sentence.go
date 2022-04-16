package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/

func areNumbersAscending(s string) bool {
	var prevVal int64 = 0
	for _, token := range strings.Split(s, " ") {
		val, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			if val <= prevVal {
				return false
			}
			prevVal = val
		}
	}
	return true
}

func main() {

	// Example 1
	var s1 string = "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
	fmt.Println("Expected: true	Output: ", areNumbersAscending(s1))

	// Example 2
	var s2 string = "hello world 5 x 5"
	fmt.Println("Expected: false	Output: ", areNumbersAscending(s2))

	// Example 3
	var s3 string = "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"
	fmt.Println("Expected: false	Output: ", areNumbersAscending(s3))

}
