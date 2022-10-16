package main

// source: https://leetcode.com/problems/first-bad-version/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version > 3 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !isBadVersion(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func main() {
	print(firstBadVersion(5))
}
