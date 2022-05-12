package main

import "fmt"

// source: https://leetcode.com/problems/count-sorted-vowel-strings/

func countVowelStrings(n int) int {
	var res int
	var util func(i, l int)
	util = func(i, l int) {
		if l == n {
			res++
			return
		}
		for j := i; j < 5; j++ {
			util(j, l+1)
		}
	}

	util(0, 0)
	return res
}

func main() {
	// Example 1
	var n1 int = 1
	fmt.Println("Expected: 5	Output: ", countVowelStrings(n1))

	// Example 2
	var n2 int = 2
	fmt.Println("Expected: 15	Output: ", countVowelStrings(n2))

	// Example 3
	var n3 int = 33
	fmt.Println("Expected: 66045	Output: ", countVowelStrings(n3))

}
