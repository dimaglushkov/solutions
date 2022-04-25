package main

import "fmt"

// source: https://leetcode.com/problems/generate-parentheses/

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	generateParenthesisUtil(n, n, "")
	return res
}

func generateParenthesisUtil(l, r int, s string) {
	if r < l || r < 0 || l < 0 {
		return
	}
	if l == 0 && r == 0 {
		res = append(res, s)
		return
	}
	generateParenthesisUtil(l-1, r, s+"(")
	generateParenthesisUtil(l, r-1, s+")")
}

func main() {
	// Example 1
	var n1 int = 3
	fmt.Println("Expected: [\"((()))\",\"(()())\",\"(())()\",\"()(())\",\"()()()\"]	Output: ", generateParenthesis(n1))

	// Example 2
	var n2 int = 1
	fmt.Println("Expected: [\"()\"]	Output: ", generateParenthesis(n2))

}
