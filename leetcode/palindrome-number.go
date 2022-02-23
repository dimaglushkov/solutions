package main

import "fmt"

/*
source: https://leetcode.com/problems/palindrome-number/

problem: Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.
    For example, 121 is a palindrome while 123 is not.
-231 <= x <= 231 - 1
*/

// go digit by digit from the both sides of the number and compare them
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	strX := fmt.Sprint(x)
	for i, _ := range strX[:len(strX)/2] {
		if strX[i] != strX[len(strX)-1-i] {
			return false
		}
	}
	return true
}

// flip the number and check if flipped one is equal to the source one
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	flippedX := 0
	for dx := x; dx > 0; dx /= 10 {
		flippedX = flippedX*10 + dx%10
	}
	return flippedX == x
}
