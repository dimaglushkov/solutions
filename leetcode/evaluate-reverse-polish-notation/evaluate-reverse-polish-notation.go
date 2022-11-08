package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/evaluate-reverse-polish-notation/

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	push := func(x string) {
		stack = append(stack, x)
	}
	pop := func() string {
		n := len(stack)
		x := stack[n-1]
		stack = stack[:n-1]
		return x
	}
	for _, t := range tokens {
		push(t)
	}

	var evaluate func() int
	evaluate = func() int {
		x := pop()
		if v, err := strconv.ParseInt(x, 10, 64); err == nil {
			return int(v)
		}

		r := evaluate()
		l := evaluate()

		switch x {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			return l / r
		default:
			return -1
		}
	}

	return evaluate()
}

func main() {
	// Example 1
	var tokens1 = []string{"2", "1", "+", "3", "*"}
	fmt.Println("Expected: 9	Output: ", evalRPN(tokens1))

	// Example 2
	var tokens2 = []string{"4", "13", "5", "/", "+"}
	fmt.Println("Expected: 6	Output: ", evalRPN(tokens2))

	// Example 3
	var tokens3 = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println("Expected: 22	Output: ", evalRPN(tokens3))

}
