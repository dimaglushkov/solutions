package main

import "fmt"

// source: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

func maxDepth(s string) int {
	var res, tmp int
	for _, r := range s {
		if r == '(' {
			tmp++
			if tmp > res {
				res = tmp
			}
		} else if r == ')' {
			tmp--
		}

	}
	return res
}

func main() {
	// Example 1 
	var s1 string = "(1+(2*3)+((<u>8</u>)/4))+1"
	fmt.Println("Expected: 3	Output: ", maxDepth(s1))

	// Example 2 
	var s2 string = "(1)+((2))+(((<u>3</u>)))"
	fmt.Println("Expected: 3	Output: ", maxDepth(s2))

}
