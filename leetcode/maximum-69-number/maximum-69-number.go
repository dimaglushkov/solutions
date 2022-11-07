package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/maximum-69-number/

func maximum69Number(num int) int {
	res, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return res
}

func main() {
	// Example 1
	var num1 int = 9669
	fmt.Println("Expected: 9969	Output: ", maximum69Number(num1))

	// Example 2
	var num2 int = 9996
	fmt.Println("Expected: 9999	Output: ", maximum69Number(num2))

	// Example 3
	var num3 int = 9999
	fmt.Println("Expected: 9999	Output: ", maximum69Number(num3))

}
