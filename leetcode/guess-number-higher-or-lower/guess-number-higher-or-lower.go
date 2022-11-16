package main

// source: https://leetcode.com/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	i, j := 1, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if guess(h) < 0 {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
